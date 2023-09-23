package v1

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/filesystem"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"path"
	"strconv"
	"strings"
	"time"
)

type sFile struct {
	Filesystem *filesystem.Filesystem
}

func init() {
	service.RegisterFile(New())
}

func New() service.IFile {
	return &sFile{
		Filesystem: filesystem.NewFilesystem(config.Cfg),
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

	if err := s.Filesystem.Default.Write(stream, object); err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("文件上传失败")
	}

	return &model.UploadAvatarRes{
		Avatar: s.Filesystem.Default.PublicUrl(object),
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

	if err := s.Filesystem.Default.Write(stream, object); err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("文件上传失败")
	}

	return &model.UploadImageRes{
		Src: s.Filesystem.Default.PublicUrl(object),
	}, nil
}

// 批量上传初始化
func (s *sFile) InitiateMultipart(ctx context.Context, params model.UploadInitiateMultipartReq) (*model.UploadInitiateMultipartRes, error) {

	uploadId, err := service.SplitUpload().InitiateMultipartUpload(ctx, &model.MultipartInitiateOpt{
		Name:   params.FileName,
		Size:   params.FileSize,
		UserId: service.Session().GetUid(ctx),
	})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.UploadInitiateMultipartRes{
		UploadId:    uploadId,
		UploadIdMd5: gmd5.MustEncryptString(uploadId),
		SplitSize:   2 << 20,
	}, nil
}

// MultipartUpload 批量分片上传
func (s *sFile) MultipartUpload(ctx context.Context, params model.UploadMultipartReq) (*model.UploadMultipartRes, error) {

	_, file, err := g.RequestFromCtx(ctx).Request.FormFile("file")
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("文件上传失败")
	}

	if err = service.SplitUpload().MultipartUpload(ctx, &model.MultipartUploadOpt{
		UserId:     service.Session().GetUid(ctx),
		UploadId:   params.UploadId,
		SplitIndex: params.SplitIndex,
		SplitNum:   params.SplitNum,
		File:       file,
	}); err != nil {
		logger.Error(ctx, err)
		return nil, err
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
