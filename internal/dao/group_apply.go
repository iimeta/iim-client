package dao

import (
	"context"
	"errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
	"github.com/iimeta/iim-client/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
)

var GroupApply = NewGroupApplyDao()

type GroupApplyDao struct {
	*MongoDB[entity.GroupApply]
}

func NewGroupApplyDao(database ...string) *GroupApplyDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &GroupApplyDao{
		MongoDB: NewMongoDB[entity.GroupApply](database[0], do.GROUP_APPLY_COLLECTION),
	}
}

func (d *GroupApplyDao) List(ctx context.Context, groupIds []int) ([]*entity.GroupApply, []*entity.User, error) {

	groupApplyList, err := d.Find(ctx, bson.M{"group_id": bson.M{"$in": groupIds}, "status": model.GroupApplyStatusWait}, "-updated_at")
	if err != nil {
		return nil, nil, err
	}

	userIds := make([]int, 0)
	for _, apply := range groupApplyList {
		userIds = append(userIds, apply.UserId)
	}

	userList, err := User.FindUserListByUserIds(ctx, userIds)
	if err != nil {
		return nil, nil, err
	}

	return groupApplyList, userList, nil
}

func (d *GroupApplyDao) Auth(ctx context.Context, applyId string, userId int) bool {

	groupApply, err := d.FindById(ctx, applyId)
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	groupMember, err := GroupMember.FindOne(ctx, bson.M{"group_id": groupApply.GroupId, "user_id": userId, "leader": bson.M{"$in": []int{1, 2}, "is_quit": 0}})
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	return groupMember.Id != ""
}

func (d *GroupApplyDao) Delete(ctx context.Context, applyId string, userId int) error {

	if !d.Auth(ctx, applyId, userId) {
		return errors.New("auth failed")
	}

	return d.DeleteById(ctx, applyId)
}
