package dao

import (
	"context"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
	"go.mongodb.org/mongo-driver/bson"
)

var GroupNotice = NewGroupNoticeDao()

type GroupNoticeDao struct {
	*MongoDB[entity.GroupNotice]
}

func NewGroupNoticeDao(database ...string) *GroupNoticeDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &GroupNoticeDao{
		MongoDB: NewMongoDB[entity.GroupNotice](database[0], do.GROUP_NOTICE_COLLECTION),
	}
}

func (d *GroupNoticeDao) GetListAll(ctx context.Context, groupId int) ([]*entity.GroupNotice, []*entity.User, error) {

	groupNoticeList, err := d.Find(ctx, bson.M{"group_id": groupId, "is_delete": 0}, "-is_top", "-created_at")
	if err != nil {
		return nil, nil, err
	}

	userIds := make([]int, 0)
	for _, apply := range groupNoticeList {
		userIds = append(userIds, apply.CreatorId)
	}

	userList, err := User.FindUserListByUserIds(ctx, userIds)
	if err != nil {
		return nil, nil, err
	}

	return groupNoticeList, userList, nil
}

// 获取最新公告
func (d *GroupNoticeDao) GetLatestNotice(ctx context.Context, groupId int) (*entity.GroupNotice, error) {

	groupNotice, err := d.FindOne(ctx, bson.M{"group_id": groupId, "is_delete": 0}, "-created_at")
	if err != nil {
		return nil, err
	}

	return groupNotice, nil
}
