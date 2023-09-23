package note_class

import (
	"context"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
)

type sNoteClass struct{}

func init() {
	service.RegisterNoteClass(New())
}

func New() service.INoteClass {
	return &sNoteClass{}
}

// 分类列表
func (s *sNoteClass) List(ctx context.Context) (*model.ArticleClassListRes, error) {

	noteClassList, err := dao.NoteClass.List(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	data, err := dao.NoteClass.GroupCount(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	list := make([]*model.ArticleClassItem, 0)
	list = append(list, &model.ArticleClassItem{
		ClassName: "默认分类",
	})

	for i := range noteClassList {
		if num, ok := data[noteClassList[i].Id]; ok {
			list[i].Count = num
		}
	}

	items := make([]*model.ArticleClassListResponse_Item, 0, len(list))
	for _, item := range list {
		items = append(items, &model.ArticleClassListResponse_Item{
			Id:        item.Id,
			ClassName: item.ClassName,
			IsDefault: item.IsDefault,
			Count:     item.Count,
		})
	}

	return &model.ArticleClassListRes{
		Items: items,
		Paginate: &model.Paginate{
			Page:  1,
			Size:  100000,
			Total: len(items),
		},
	}, nil
}

// 添加或修改分类
func (s *sNoteClass) Edit(ctx context.Context, params model.ArticleClassEditReq) (*model.ArticleClassEditRes, error) {

	uid := service.Session().GetUid(ctx)

	if params.ClassId == "" || params.ClassId == "0" { // todo

		id, err := dao.NoteClass.Create(ctx, uid, params.ClassName)
		if err != nil {
			logger.Error(ctx, err)
			return nil, errors.New("笔记分类创建失败")
		}

		params.ClassId = id

	} else {
		if err := dao.NoteClass.Update(ctx, uid, params.ClassId, params.ClassName); err != nil {
			logger.Error(ctx, err)
			return nil, errors.New("笔记分类编辑失败")
		}
	}

	return &model.ArticleClassEditRes{
		Id: params.ClassId,
	}, nil
}

// 删除分类
func (s *sNoteClass) Delete(ctx context.Context, params model.ArticleClassDeleteReq) error {

	err := dao.NoteClass.Delete(ctx, service.Session().GetUid(ctx), params.ClassId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 分类排序
func (s *sNoteClass) Sort(ctx context.Context, params model.ArticleClassSortReq) error {

	err := dao.NoteClass.Sort(ctx, service.Session().GetUid(ctx), params.ClassId, params.SortType)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
