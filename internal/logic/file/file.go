package v1

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/filesystem"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.mongodb.org/mongo-driver/bson"
	"math"
	"path"
	"strconv"
	"strings"
	"time"
)

type sFile struct {
	filesystem *filesystem.Filesystem
}

func init() {
	service.RegisterFile(New())
}

func New() service.IFile {
	return &sFile{
		filesystem: filesystem.NewFilesystem(config.Cfg),
	}
}

// 头像上传上传
func (s *sFile) Avatar(ctx context.Context) (*model.UploadAvatarRes, error) {

	_, file, err := g.RequestFromCtx(ctx).Request.FormFile("file")
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("文件上传失败")
	}

	stream, _ := filesystem.ReadMultipartStream(file)
	object := fmt.Sprintf("public/media/image/avatar/%s/%s", time.Now().Format("20060102"), util.GenImageName("png", 200, 200))

	if err := s.filesystem.Default.Write(stream, object); err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("文件上传失败")
	}

	return &model.UploadAvatarRes{
		Avatar: s.filesystem.Default.PublicUrl(object),
	}, nil
}

// 图片上传
func (s *sFile) Image(ctx context.Context) (*model.UploadImageRes, error) {

	_, file, err := g.RequestFromCtx(ctx).Request.FormFile("file")
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("文件上传失败")
	}

	var (
		ext       = strings.TrimPrefix(path.Ext(file.Filename), ".")
		width, _  = strconv.Atoi(g.RequestFromCtx(ctx).PostFormValue("width"))
		height, _ = strconv.Atoi(g.RequestFromCtx(ctx).PostFormValue("height"))
	)

	stream, _ := filesystem.ReadMultipartStream(file)
	if width == 0 || height == 0 {
		meta := util.ReadImageMeta(bytes.NewReader(stream))
		width = meta.Width
		height = meta.Height
	}

	object := fmt.Sprintf("public/media/image/common/%s/%s", time.Now().Format("20060102"), util.GenImageName(ext, width, height))

	if err := s.filesystem.Default.Write(stream, object); err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("文件上传失败")
	}

	return &model.UploadImageRes{
		Src: s.filesystem.Default.PublicUrl(object),
	}, nil
}

// 批量上传初始化
func (s *sFile) InitiateMultipart(ctx context.Context, params model.UploadInitiateMultipartReq) (*model.UploadInitiateMultipartRes, error) {

	uid := service.Session().GetUid(ctx)

	// 计算拆分数量 3M
	num := math.Ceil(float64(params.FileSize) / float64(3<<20))

	m := &do.SplitUpload{
		Type:         1,
		Drive:        consts.FileDriveMode(s.filesystem.Driver()),
		UserId:       uid,
		OriginalName: params.FileName,
		SplitNum:     int(num),
		FileExt:      strings.TrimPrefix(path.Ext(params.FileName), "."),
		FileSize:     params.FileSize,
		Path:         fmt.Sprintf("private/tmp/multipart/%s/%s.tmp", util.DateNumber(), gmd5.MustEncryptString(util.Random(20))),
		Attr:         "{}",
	}

	uploadId, err := s.filesystem.Default.InitiateMultipartUpload(m.Path, m.OriginalName)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	m.UploadId = uploadId

	if _, err := dao.SplitUpload.Insert(ctx, m); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.UploadInitiateMultipartRes{
		UploadId:    m.UploadId,
		UploadIdMd5: gmd5.MustEncryptString(m.UploadId),
		SplitSize:   2 << 20,
	}, nil
}

// 批量分片上传
func (s *sFile) MultipartSplitUpload(ctx context.Context, params model.UploadMultipartReq) (*model.UploadMultipartRes, error) {

	_, file, err := g.RequestFromCtx(ctx).Request.FormFile("file")
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("文件上传失败")
	}

	uid := service.Session().GetUid(ctx)

	info, err := dao.SplitUpload.FindOne(ctx, bson.M{"upload_id": params.UploadId, "type": 1})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	stream, err := filesystem.ReadMultipartStream(file)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	dirPath := fmt.Sprintf("private/tmp/%s/%s/%d-%s.tmp", util.DateNumber(), gmd5.MustEncryptString(params.UploadId), params.SplitIndex, params.UploadId)

	data := &do.SplitUpload{
		Type:         2,
		Drive:        info.Drive,
		UserId:       uid,
		UploadId:     params.UploadId,
		OriginalName: info.OriginalName,
		SplitIndex:   params.SplitIndex,
		SplitNum:     params.SplitNum,
		Path:         dirPath,
		FileExt:      info.FileExt,
		FileSize:     file.Size,
		Attr:         "{}",
	}

	switch data.Drive {
	case consts.FileDriveLocal:
		_ = s.filesystem.Default.Write(stream, data.Path)
	case consts.FileDriveCos:
		etag, err := s.filesystem.Cos.UploadPart(info.Path, data.UploadId, data.SplitIndex+1, stream)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		data.Attr = gjson.MustEncodeString(map[string]string{
			"etag": etag,
		})

	default:
		return nil, errors.New("未知文件驱动类型")
	}

	if _, err := dao.SplitUpload.Insert(ctx, data); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	// 判断是否为最后一个分片上传
	if params.SplitNum == params.SplitIndex+1 {
		if err = s.merge(ctx, info); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if params.SplitIndex != params.SplitNum-1 {
		return &model.UploadMultipartRes{
			IsMerge: false,
		}, nil
	}

	return &model.UploadMultipartRes{
		UploadId: params.UploadId,
		IsMerge:  true,
	}, nil
}

func (s *sFile) merge(ctx context.Context, info *entity.SplitUpload) error {

	items, err := dao.SplitUpload.GetSplitList(ctx, info.UploadId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	switch info.Drive {
	case consts.FileDriveLocal:
		for _, item := range items {
			stream, err := s.filesystem.Default.ReadStream(item.Path)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			if err := s.filesystem.Local.AppendWrite(stream, info.Path); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	case consts.FileDriveCos:
		opt := &cos.CompleteMultipartUploadOptions{}
		for _, item := range items {
			attr := make(map[string]string)

			if err := gjson.Unmarshal([]byte(item.Attr), &attr); err != nil {
				logger.Error(ctx, err)
				return err
			}

			opt.Parts = append(opt.Parts, cos.Object{
				PartNumber: item.SplitIndex + 1,
				ETag:       attr["etag"],
			})
		}

		if err := s.filesystem.Cos.CompleteMultipartUpload(info.Path, info.UploadId, opt); err != nil {
			logger.Error(ctx, err)
			return err
		}
	default:
		return errors.New("未知文件驱动类型")
	}

	return nil
}
