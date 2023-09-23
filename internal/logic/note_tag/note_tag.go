package note_tag

import (
	"context"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
)

type sNoteTag struct{}

func init() {
	service.RegisterNoteTag(New())
}

func New() service.INoteTag {
	return &sNoteTag{}
}

// 标签列表
func (s *sNoteTag) List(ctx context.Context) (*model.ArticleTagListRes, error) {

	list, err := dao.NoteTag.List(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.ArticleTagListResponse_Item, 0, len(list))
	for _, item := range list {
		items = append(items, &model.ArticleTagListResponse_Item{
			Id:      item.Id,
			TagName: item.TagName,
			Count:   item.Count,
		})
	}

	return &model.ArticleTagListRes{Tags: items}, nil
}

// 添加或修改标签
func (s *sNoteTag) Edit(ctx context.Context, params model.ArticleTagEditReq) (*model.ArticleTagEditRes, error) {

	uid := service.Session().GetUid(ctx)

	if params.TagId == "" || params.TagId == "0" { // todo

		id, err := dao.NoteTag.Create(ctx, uid, params.TagName)
		if err != nil {
			logger.Error(ctx, err)
			return nil, errors.New("笔记标签创建失败")
		}

		params.TagId = id

	} else {
		if err := dao.NoteTag.Update(ctx, uid, params.TagId, params.TagName); err != nil {
			logger.Error(ctx, err)
			return nil, errors.New("笔记标签编辑失败")
		}
	}

	return &model.ArticleTagEditRes{Id: params.TagId}, nil
}

// 删除标签
func (s *sNoteTag) Delete(ctx context.Context, params model.ArticleTagDeleteReq) error {

	err := dao.NoteTag.Delete(ctx, service.Session().GetUid(ctx), params.TagId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
