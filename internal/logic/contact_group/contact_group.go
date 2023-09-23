package contact_group

import (
	"context"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
)

type sContactGroup struct{}

func init() {
	service.RegisterContactGroup(New())
}

func New() service.IContactGroup {
	return &sContactGroup{}
}

func (s *sContactGroup) Delete(ctx context.Context, id int, uid int) error {

	result, err := dao.ContactGroup.DeleteOne(ctx, bson.M{"_id": id, "user_id": uid})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if result == 0 {
		return errors.New("数据不存在")
	}

	if err := dao.Contact.UpdateOne(ctx, bson.M{"user_id": uid, "group_id": id}, bson.M{
		"group_id": 0,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 用户联系人分组列表
func (s *sContactGroup) GetUserGroup(ctx context.Context, uid int) ([]*model.ContactGroup, error) {

	contactGroupList, err := dao.ContactGroup.FindContactGroupList(ctx, uid)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.ContactGroup, 0)
	for _, contactGroup := range contactGroupList {
		items = append(items, &model.ContactGroup{
			Id:        contactGroup.Id,
			UserId:    contactGroup.UserId,
			Name:      contactGroup.Name,
			Num:       contactGroup.Num,
			Sort:      contactGroup.Sort,
			CreatedAt: contactGroup.CreatedAt,
			UpdatedAt: contactGroup.UpdatedAt,
		})
	}

	return items, nil
}

// List 联系人分组列表
func (s *sContactGroup) List(ctx context.Context) (*model.ContactGroupListRes, error) {

	uid := service.Session().GetUid(ctx)

	items := make([]*model.ContactGroupListResponse_Item, 0)

	count, err := dao.Contact.CountDocuments(ctx, bson.M{"user_id": uid, "status": 1})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items = append(items, &model.ContactGroupListResponse_Item{
		Name:  "全部",
		Count: int(count),
	})

	group, err := s.GetUserGroup(ctx, uid)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, v := range group {
		items = append(items, &model.ContactGroupListResponse_Item{
			Id:    v.Id,
			Name:  v.Name,
			Count: v.Num,
			Sort:  v.Sort,
		})
	}

	return &model.ContactGroupListRes{Items: items}, nil
}

func (s *sContactGroup) Save(ctx context.Context, params model.ContactGroupSaveReq) error {

	uid := service.Session().GetUid(ctx)

	updateItems := make([]*do.ContactGroup, 0)
	deleteItems := make([]string, 0)
	insertItems := make([]interface{}, 0)

	ids := make(map[string]struct{})
	for i, item := range params.Items {
		if item.Id != "" && item.Id != "0" {
			ids[item.Id] = struct{}{}
			updateItems = append(updateItems, &do.ContactGroup{
				Id:   item.Id,
				Sort: i + 1,
				Name: item.Name,
			})
		} else {
			insertItems = append(insertItems, &do.ContactGroup{
				Sort:   i + 1,
				Name:   item.Name,
				UserId: uid,
			})
		}
	}

	contactGroupList, err := dao.ContactGroup.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, m := range contactGroupList {
		if _, ok := ids[m.Id]; !ok {
			deleteItems = append(deleteItems, m.Id)
		}
	}

	if len(insertItems) > 0 {
		if _, err := dao.ContactGroup.Inserts(ctx, insertItems); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if len(deleteItems) > 0 {
		if _, err := dao.ContactGroup.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": deleteItems}, "user_id": uid}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err := dao.Contact.UpdateMany(ctx, bson.M{"user_id": uid, "group_id": bson.M{"$in": deleteItems}}, bson.M{
			"group_id": 0,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	for _, item := range updateItems {
		if err := dao.ContactGroup.UpdateOne(ctx, bson.M{"_id": item.Id, "user_id": uid}, bson.M{
			"name": item.Name,
			"sort": item.Sort,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}
