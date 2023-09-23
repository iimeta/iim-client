package dao

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	relation "github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/db"
	"github.com/iimeta/iim-client/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
)

var GroupMember = NewGroupMemberDao()

type GroupMemberDao struct {
	*MongoDB[entity.GroupMember]
}

func NewGroupMemberDao(database ...string) *GroupMemberDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &GroupMemberDao{
		MongoDB: NewMongoDB[entity.GroupMember](database[0], do.GROUP_MEMBER_COLLECTION),
	}
}

// 判断是否是群主
func (d *GroupMemberDao) IsMaster(ctx context.Context, gid, uid int) bool {

	count, err := d.CountDocuments(ctx, bson.M{"group_id": gid, "user_id": uid, "leader": 2, "is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes}})
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	return count > 0
}

// 判断是否是群主或管理员
func (d *GroupMemberDao) IsLeader(ctx context.Context, gid, uid int) bool {

	count, err := d.CountDocuments(ctx, bson.M{"group_id": gid, "user_id": uid, "leader": bson.M{"$in": []int{1, 2}}, "is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes}})
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	return count > 0
}

// 检测是属于群成员
func (d *GroupMemberDao) IsMember(ctx context.Context, gid, uid int, cache bool) bool {

	if cache && relation.IsGroupRelation(ctx, uid, gid) {
		return true
	}

	count, err := d.CountDocuments(ctx, bson.M{"group_id": gid, "user_id": uid, "is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes}})
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	if count > 0 {
		relation.SetGroupRelation(ctx, uid, gid)
	}

	return count > 0
}

func (d *GroupMemberDao) FindByUserId(ctx context.Context, gid, uid int) (*entity.GroupMember, error) {

	groupMember, err := d.FindOne(ctx, bson.M{"group_id": gid, "user_id": uid, "is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes}})
	if err != nil {
		return nil, err
	}

	return groupMember, nil
}

// 获取所有群成员用户ID
func (d *GroupMemberDao) GetMemberIds(ctx context.Context, groupId int) []int {

	groupMemberList, err := d.Find(ctx, bson.M{"group_id": groupId, "is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes}})
	if err != nil {
		logger.Error(ctx, err)
		return nil
	}

	ids := make([]int, 0)
	for _, member := range groupMemberList {
		ids = append(ids, member.UserId)
	}

	return ids
}

// 获取所有群成员ID
func (d *GroupMemberDao) GetUserGroupIds(ctx context.Context, uid int) []int {

	groupMemberList, err := d.Find(ctx, bson.M{"user_id": uid, "is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes}})
	if err != nil {
		logger.Error(ctx, err)
		return nil
	}

	ids := make([]int, 0)
	for _, member := range groupMemberList {
		ids = append(ids, member.GroupId)
	}

	return ids
}

// 统计群成员总数
func (d *GroupMemberDao) CountMemberTotal(ctx context.Context, gid int) int64 {

	count, err := d.CountDocuments(ctx, bson.M{"group_id": gid, "is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes}})
	if err != nil {
		logger.Error(ctx, err)
		return 0
	}

	return count
}

// 获取指定群成员的备注信息
func (d *GroupMemberDao) GetMemberRemark(ctx context.Context, groupId int, userId int) string {

	groupMember, err := d.FindOne(ctx, bson.M{"group_id": groupId, "user_id": userId})
	if err != nil {
		logger.Error(ctx, err)
		return ""
	}

	return groupMember.UserCard
}

// 获取群聊成员列表
func (d *GroupMemberDao) GetMembers(ctx context.Context, groupId int) ([]*entity.GroupMember, []*entity.User, error) {

	groupMemberList, err := d.Find(ctx, bson.M{"group_id": groupId, "is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes}}, "-leader")
	if err != nil {
		return nil, nil, err
	}

	userIds := make([]int, 0)
	for _, member := range groupMemberList {
		userIds = append(userIds, member.UserId)
	}

	userList, err := User.FindUserListByUserIds(ctx, userIds)
	if err != nil {
		return nil, nil, err
	}

	return groupMemberList, userList, nil
}

func (d *GroupMemberDao) CountGroupMemberNum(ctx context.Context, ids []int) ([]*entity.CountGroupMember, error) {

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"group_id": bson.M{
					"$in": ids,
				},
				"is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes},
			},
		},
		{
			"$group": bson.M{
				"_id":      "$group_id",
				"count":    bson.M{"$sum": 1},
				"group_id": bson.M{"$first": "$group_id"},
			},
		},
	}

	results := make([]map[string]interface{}, 0)
	if err := d.Aggregate(ctx, pipeline, &results); err != nil {
		return nil, err
	}

	items := make([]*entity.CountGroupMember, 0)
	for _, result := range results {
		countGroupMember := &entity.CountGroupMember{
			GroupId: gconv.Int(result["group_id"]),
			Count:   gconv.Int(result["count"]),
		}
		items = append(items, countGroupMember)
	}

	return items, nil
}

func (d *GroupMemberDao) CheckUserGroup(ctx context.Context, ids []int, userId int) ([]int, error) {

	groupMemberList, err := d.Find(ctx, bson.M{"group_id": bson.M{"$in": ids}, "user_id": userId, "is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes}})
	if err != nil {
		return nil, err
	}

	groupIds := make([]int, 0)
	for _, member := range groupMemberList {
		groupIds = append(groupIds, member.GroupId)
	}

	return groupIds, nil

}

// 交接群主权限
func (d *GroupMemberDao) Handover(ctx context.Context, groupId int, userId int, memberId int) error {

	if err := d.UpdateOne(ctx, bson.M{"group_id": groupId, "user_id": userId, "leader": 2}, &do.GroupMember{
		Leader: 0,
	}); err != nil {
		return err
	}

	if err := d.UpdateOne(ctx, bson.M{"group_id": groupId, "user_id": memberId}, &do.GroupMember{
		Leader: 2,
	}); err != nil {
		return err
	}

	return nil
}

func (d *GroupMemberDao) SetLeaderStatus(ctx context.Context, groupId int, userId int, leader int) error {

	if err := d.UpdateOne(ctx, bson.M{"group_id": groupId, "user_id": userId}, &do.GroupMember{
		Leader: leader,
	}); err != nil {
		return err
	}

	return nil
}

func (d *GroupMemberDao) SetMuteStatus(ctx context.Context, groupId int, userId int, status int) error {

	if err := d.UpdateOne(ctx, bson.M{"group_id": groupId, "user_id": userId}, &do.GroupMember{
		IsMute: status,
	}); err != nil {
		return err
	}

	return nil
}

func (d *GroupMemberDao) FindGroupIds(ctx context.Context, userId int) ([]int, error) {

	groupMemberList, err := d.Find(ctx, bson.M{"user_id": userId, "is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes}})
	if err != nil {
		return nil, err
	}

	groupIds := make([]int, 0)
	for _, member := range groupMemberList {
		groupIds = append(groupIds, member.GroupId)
	}

	return groupIds, nil
}

func (d *GroupMemberDao) FindGroupMemberListByUserId(ctx context.Context, userId int) ([]*entity.GroupMember, error) {

	groupMemberList, err := d.Find(ctx, bson.M{"user_id": userId, "is_quit": bson.M{"$ne": model.GroupMemberQuitStatusYes}})
	if err != nil {
		return nil, err
	}

	return groupMemberList, nil
}
