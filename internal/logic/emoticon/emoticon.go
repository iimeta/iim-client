package emoticon

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/filesystem"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type sEmoticon struct {
	Filesystem *filesystem.Filesystem
	RedisLock  *cache.RedisLock
}

func init() {
	service.RegisterEmoticon(New())
}

func New() service.IEmoticon {
	return &sEmoticon{
		Filesystem: filesystem.NewFilesystem(config.Cfg),
		RedisLock:  cache.NewRedisLock(redis.Client),
	}
}

// 收藏列表
func (s *sEmoticon) CollectList(ctx context.Context) (*model.ListRes, error) {

	var (
		uid  = service.Session().GetUid(ctx)
		resp = &model.ListRes{
			SysEmoticon:     make([]*model.ListResponse_SysEmoticon, 0),
			CollectEmoticon: make([]*model.ListItem, 0),
		}
	)

	if ids := dao.Emoticon.GetUserInstallIds(ctx, uid); len(ids) > 0 {

		emoticonList, err := dao.Emoticon.Find(ctx, bson.M{"_id": bson.M{"$in": ids}})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		for _, item := range emoticonList {
			data := &model.ListResponse_SysEmoticon{
				EmoticonId: item.Id,
				Url:        item.Icon,
				Name:       item.Name,
				List:       make([]*model.ListItem, 0),
			}

			if list, err := dao.Emoticon.GetDetailsAll(ctx, item.Id, 0); err == nil {
				for _, v := range list {
					data.List = append(data.List, &model.ListItem{
						MediaId: v.Id,
						Src:     v.Url,
					})
				}
			}

			resp.SysEmoticon = append(resp.SysEmoticon, data)
		}
	}

	if items, err := dao.Emoticon.GetDetailsAll(ctx, "", uid); err == nil {
		for _, item := range items {
			resp.CollectEmoticon = append(resp.CollectEmoticon, &model.ListItem{
				MediaId: item.Id,
				Src:     item.Url,
			})
		}
	}

	return resp, nil
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

	if !util.Include(util.FileSuffix(file.Filename), []string{"png", "jpg", "jpeg", "gif"}) {
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
	ext := util.FileSuffix(file.Filename)
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

// 系统表情包列表
func (s *sEmoticon) SystemList(ctx context.Context) ([]*model.SysListResponse_Item, error) {

	items, err := dao.Emoticon.GetSystemEmoticonList(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	ids := dao.Emoticon.GetUserInstallIds(ctx, service.Session().GetUid(ctx))

	data := make([]*model.SysListResponse_Item, 0)
	for _, item := range items {
		data = append(data, &model.SysListResponse_Item{
			Id:     item.Id,
			Name:   item.Name,
			Icon:   item.Icon,
			Status: util.BoolToInt(util.Include(item.Id, ids)), // 查询用户是否使用
		})
	}

	return data, nil
}

// 添加或移除系统表情包
func (s *sEmoticon) SetSystemEmoticon(ctx context.Context, params model.SetSystemReq) (*model.SetSystemRes, error) {

	var (
		err error
		uid = service.Session().GetUid(ctx)
		key = fmt.Sprintf("sys-emoticon:%d", uid)
	)

	if !s.RedisLock.Lock(ctx, key, 5) {
		return nil, errors.New("请求频繁")
	}
	defer s.RedisLock.UnLock(ctx, key)

	if params.Type == 2 {
		if err = dao.Emoticon.RemoveUserSysEmoticon(ctx, uid, params.EmoticonId); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		return nil, nil
	}

	// 查询表情包是否存在
	emoticon, err := dao.Emoticon.FindById(ctx, params.EmoticonId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if err = dao.Emoticon.AddUserSysEmoticon(ctx, uid, params.EmoticonId); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.ListItem, 0)
	if list, err := dao.Emoticon.GetDetailsAll(ctx, params.EmoticonId, 0); err == nil {
		for _, item := range list {
			items = append(items, &model.ListItem{
				MediaId: item.Id,
				Src:     item.Url,
			})
		}
	}

	return &model.SetSystemRes{
		EmoticonId: emoticon.Id,
		Url:        emoticon.Icon,
		Name:       emoticon.Name,
		List:       items,
	}, nil
}
