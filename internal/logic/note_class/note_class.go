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
func (s *sNoteClass) List(ctx context.Context) (*model.ClassListRes, error) {

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

	list := make([]*model.ClassItem, 0)
	list = append(list, &model.ClassItem{
		ClassName: "默认分类",
	})

	for _, noteClass := range noteClassList {

		classItem := &model.ClassItem{
			Id:        noteClass.Id,
			ClassName: noteClass.ClassName,
			IsDefault: noteClass.IsDefault,
		}

		if num, ok := data[noteClass.Id]; ok {
			classItem.Count = num
		}

		list = append(list, classItem)
	}

	items := make([]*model.ClassListResponse_Item, 0, len(list))
	for _, item := range list {
		items = append(items, &model.ClassListResponse_Item{
			Id:        item.Id,
			ClassName: item.ClassName,
			IsDefault: item.IsDefault,
			Count:     item.Count,
		})
	}

	return &model.ClassListRes{
		Items: items,
		Paginate: &model.Paginate{
			Page:  1,
			Size:  100000,
			Total: len(items),
		},
	}, nil
}

// 添加或修改分类
func (s *sNoteClass) Edit(ctx context.Context, params model.ClassEditReq) (*model.ClassEditRes, error) {

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

	return &model.ClassEditRes{
		Id: params.ClassId,
	}, nil
}

// 删除分类
func (s *sNoteClass) Delete(ctx context.Context, params model.ClassDeleteReq) error {

	err := dao.NoteClass.Delete(ctx, service.Session().GetUid(ctx), params.ClassId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 分类排序
func (s *sNoteClass) Sort(ctx context.Context, params model.ClassSortReq) error {

	err := dao.NoteClass.Sort(ctx, service.Session().GetUid(ctx), params.ClassId, params.SortType)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
