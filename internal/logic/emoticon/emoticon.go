package emoticon

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/filesystem"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"slices"
	"time"
)

type sEmoticon struct {
	Filesystem *filesystem.Filesystem
}

func init() {
	service.RegisterEmoticon(New())
}

func New() service.IEmoticon {
	return &sEmoticon{
		Filesystem: filesystem.NewFilesystem(config.Cfg),
	}
}

// 收藏列表
func (s *sEmoticon) CollectList(ctx context.Context) ([]*model.CollectEmoticon, error) {

	collectEmoticons := make([]*model.CollectEmoticon, 0)

	if items, err := dao.Emoticon.GetDetailsAll(ctx, "", service.Session().GetUid(ctx)); err == nil {
		for _, item := range items {
			collectEmoticons = append(collectEmoticons, &model.CollectEmoticon{
				MediaId: item.Id,
				Src:     item.Url,
			})
		}
	}

	return collectEmoticons, nil
}

// 删除收藏表情包
func (s *sEmoticon) DeleteCollect(ctx context.Context, params model.DeleteReq) error {

	if _, err := dao.Emoticon.DeleteCollect(ctx, service.Session().GetUid(ctx), util.ParseIds(params.Ids)); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 上传自定义表情包
func (s *sEmoticon) Upload(ctx context.Context) (*model.UploadRes, error) {

	_, file, err := g.RequestFromCtx(ctx).Request.FormFile("emoticon")
	if err != nil {
		return nil, errors.New("emoticon 字段必传")
	}

	if !slices.Contains([]string{"png", "jpg", "jpeg", "gif"}, gfile.ExtName(file.Filename)) {
		return nil, errors.New("上传文件格式不正确,仅支持 png、jpg、jpeg 和 gif")
	}

	// 判断上传文件大小(5M)
	if file.Size > 5<<20 {
		return nil, errors.New("上传文件大小不能超过5M")
	}

	stream, err := filesystem.ReadMultipartStream(file)
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("上传失败")
	}

	meta := util.ReadImageMeta(bytes.NewReader(stream))
	ext := gfile.ExtName(file.Filename)
	src := fmt.Sprintf("public/media/image/emoticon/%s/%s", time.Now().Format("20060102"), util.GenImageName(ext, meta.Width, meta.Height))
	if err = s.Filesystem.Default.Write(stream, src); err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("上传失败")
	}

	m := &do.EmoticonItem{
		UserId:     service.Session().GetUid(ctx),
		Describe:   "自定义表情包",
		Url:        s.Filesystem.Default.PublicUrl(src),
		FileSuffix: ext,
		FileSize:   int(file.Size),
	}

	id, err := dao.Insert(ctx, dao.Emoticon.Database, m)
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("上传失败")
	}

	return &model.UploadRes{
		MediaId: id,
		Src:     m.Url,
	}, nil
}
