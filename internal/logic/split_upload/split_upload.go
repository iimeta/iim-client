package split_upload

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/filesystem"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"math"
	"path"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type sSplitUpload struct {
	fileSystem *filesystem.Filesystem
}

func init() {
	service.RegisterSplitUpload(New())
}

func New() service.ISplitUpload {
	return &sSplitUpload{
		fileSystem: filesystem.NewFilesystem(config.Cfg),
	}
}

func (s *sSplitUpload) InitiateMultipartUpload(ctx context.Context, params *model.MultipartInitiateOpt) (string, error) {

	// 计算拆分数量 3M
	num := math.Ceil(float64(params.Size) / float64(3<<20))

	m := &do.SplitUpload{
		Type:         1,
		Drive:        consts.FileDriveMode(s.fileSystem.Driver()),
		UserId:       params.UserId,
		OriginalName: params.Name,
		SplitNum:     int(num),
		FileExt:      strings.TrimPrefix(path.Ext(params.Name), "."),
		FileSize:     params.Size,
		Path:         fmt.Sprintf("private/tmp/multipart/%s/%s.tmp", util.DateNumber(), gmd5.MustEncryptString(util.Random(20))),
		Attr:         "{}",
	}

	uploadId, err := s.fileSystem.Default.InitiateMultipartUpload(m.Path, m.OriginalName)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	m.UploadId = uploadId

	if _, err := dao.SplitUpload.Insert(ctx, m); err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	return m.UploadId, nil
}

func (s *sSplitUpload) MultipartUpload(ctx context.Context, opt *model.MultipartUploadOpt) error {

	info, err := dao.SplitUpload.FindOne(ctx, bson.M{"upload_id": opt.UploadId, "type": 1})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	stream, err := filesystem.ReadMultipartStream(opt.File)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	dirPath := fmt.Sprintf("private/tmp/%s/%s/%d-%s.tmp", util.DateNumber(), gmd5.MustEncryptString(opt.UploadId), opt.SplitIndex, opt.UploadId)

	data := &do.SplitUpload{
		Type:         2,
		Drive:        info.Drive,
		UserId:       opt.UserId,
		UploadId:     opt.UploadId,
		OriginalName: info.OriginalName,
		SplitIndex:   opt.SplitIndex,
		SplitNum:     opt.SplitNum,
		Path:         dirPath,
		FileExt:      info.FileExt,
		FileSize:     opt.File.Size,
		Attr:         "{}",
	}

	switch data.Drive {
	case consts.FileDriveLocal:
		_ = s.fileSystem.Default.Write(stream, data.Path)
	case consts.FileDriveCos:
		etag, err := s.fileSystem.Cos.UploadPart(info.Path, data.UploadId, data.SplitIndex+1, stream)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		data.Attr = gjson.MustEncodeString(map[string]string{
			"etag": etag,
		})

	default:
		return errors.New("未知文件驱动类型")
	}

	if _, err := dao.SplitUpload.Insert(ctx, data); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 判断是否为最后一个分片上传
	if opt.SplitNum == opt.SplitIndex+1 {
		if err = s.merge(ctx, info); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

func (s *sSplitUpload) merge(ctx context.Context, info *entity.SplitUpload) error {

	items, err := dao.SplitUpload.GetSplitList(ctx, info.UploadId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	switch info.Drive {
	case consts.FileDriveLocal:
		for _, item := range items {
			stream, err := s.fileSystem.Default.ReadStream(item.Path)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			if err := s.fileSystem.Local.AppendWrite(stream, info.Path); err != nil {
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

		if err := s.fileSystem.Cos.CompleteMultipartUpload(info.Path, info.UploadId, opt); err != nil {
			logger.Error(ctx, err)
			return err
		}
	default:
		return errors.New("未知文件驱动类型")
	}

	return nil
}
