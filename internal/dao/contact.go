package dao

import (
	"context"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/db"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

var Contact = NewContactDao()

type ContactDao struct {
	*MongoDB[entity.Contact]
	cache    *cache.ContactRemark
	relation *cache.Relation
}

func NewContactDao(database ...string) *ContactDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &ContactDao{
		MongoDB:  NewMongoDB[entity.Contact](database[0], do.CONTACT_COLLECTION),
		cache:    cache.NewContactRemark(redis.Client),
		relation: cache.NewRelation(redis.Client),
	}
}

// 获取好友备注
func (d *ContactDao) Remarks(ctx context.Context, uid int, fids []int) (map[int]string, error) {

	if !d.cache.Exist(ctx, uid) {
		_ = d.LoadContactCache(ctx, uid)
	}

	return d.cache.MGet(ctx, uid, fids)
}

// 判断是否为好友关系
func (d *ContactDao) IsFriend(ctx context.Context, uid int, friendId int, cache bool) bool {

	if cache && d.relation.IsContactRelation(ctx, uid, friendId) == nil {
		return true
	}

	filter := bson.M{
		"status": consts.ContactStatusNormal,
		"$or": bson.A{
			bson.M{
				"user_id":   uid,
				"friend_id": friendId,
			},
			bson.M{
				"user_id":   friendId,
				"friend_id": uid,
			},
		},
	}

	count, err := d.CountDocuments(ctx, filter)
	if err != nil {
		return false
	}

	if count == 2 {
		d.relation.SetContactRelation(ctx, uid, friendId)
	} else {
		d.relation.DelContactRelation(ctx, uid, friendId)
	}

	return count == 2
}

func (d *ContactDao) GetFriendRemark(ctx context.Context, uid int, friendId int) string {

	if d.cache.Exist(ctx, uid) {
		return d.cache.Get(ctx, uid, friendId)
	}

	filter := bson.M{
		"user_id":   uid,
		"friend_id": friendId,
	}

	contact, err := d.FindOne(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return ""
	}

	return contact.Remark
}

func (d *ContactDao) SetFriendRemark(ctx context.Context, uid int, friendId int, remark string) error {
	return d.cache.Set(ctx, uid, friendId, remark)
}

func (d *ContactDao) LoadContactCache(ctx context.Context, uid int) error {

	filter := bson.M{
		"user_id": uid,
		"status":  consts.ContactStatusNormal,
	}

	contactList, err := d.Find(ctx, filter)
	if err != nil {
		return err
	}

	items := make(map[string]any)
	for _, value := range contactList {
		if len(value.Remark) > 0 {
			items[strconv.Itoa(value.FriendId)] = value.Remark
		}
	}

	return d.cache.MSet(ctx, uid, items)
}

// 修改好友备注
func (d *ContactDao) UpdateRemark(ctx context.Context, uid int, friendId int, remark string) error {

	filter := bson.M{
		"user_id":   uid,
		"friend_id": friendId,
	}

	if err := d.UpdateOne(ctx, filter, bson.M{"remark": remark}); err != nil {
		return err
	}

	if err := d.SetFriendRemark(ctx, uid, friendId, remark); err != nil {
		return err
	}

	return nil
}

// 删除好友
func (d *ContactDao) Delete(ctx context.Context, uid, friendId int) error {

	filter := bson.M{
		"user_id":   uid,
		"friend_id": friendId,
	}

	contact, err := d.FindOne(ctx, filter)
	if err != nil {
		return err
	}

	if contact.GroupId != "" {

		filter := bson.M{
			"_id":     contact.GroupId,
			"user_id": uid,
		}

		if err := ContactGroup.UpdateOne(ctx, filter, bson.M{
			"$inc": bson.M{
				"count": -1,
			},
		}); err != nil {
			return err
		}
	}

	if err := d.UpdateOne(ctx, filter, bson.M{
		"status": consts.ContactStatusDelete,
	}); err != nil {
		return err
	}

	return nil
}

// 好友列表
func (d *ContactDao) List(ctx context.Context, uid int) ([]*entity.Contact, []*entity.User, error) {

	filter := bson.M{
		"user_id": uid,
		"status":  consts.ContactStatusNormal,
	}

	contactList, err := d.Find(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	userIds := make([]int, 0)
	for _, contact := range contactList {
		userIds = append(userIds, contact.FriendId)
	}

	userList, err := User.FindUserListByUserIds(ctx, userIds)
	if err != nil {
		return nil, nil, err
	}

	return contactList, userList, nil
}

func (d *ContactDao) GetContactIds(ctx context.Context, uid int) []int {

	filter := bson.M{
		"user_id": uid,
		"status":  consts.ContactStatusNormal,
	}

	contactList, err := d.Find(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return nil
	}

	ids := make([]int, 0)
	for _, contact := range contactList {
		ids = append(ids, contact.FriendId)
	}

	return ids
}

func (d *ContactDao) MoveGroup(ctx context.Context, uid int, friendId int, groupId string) error {

	filter := bson.M{
		"user_id":   uid,
		"friend_id": friendId,
	}

	contact, err := d.FindOne(ctx, filter)
	if err != nil {
		return err
	}

	if contact.GroupId != "" {

		filter := bson.M{
			"_id":     contact.GroupId,
			"user_id": uid,
		}

		if err := ContactGroup.UpdateOne(ctx, filter, bson.M{
			"$inc": bson.M{
				"count": -1,
			},
		}); err != nil {
			return err
		}
	}

	filter = bson.M{
		"user_id":   uid,
		"friend_id": friendId,
	}

	if err := d.UpdateOne(ctx, filter, bson.M{
		"group_id": groupId,
	}); err != nil {
		return err
	}

	filter = bson.M{
		"_id":     groupId,
		"user_id": uid,
	}

	if err := ContactGroup.UpdateOne(ctx, filter, bson.M{
		"$inc": bson.M{
			"count": 1,
		},
	}); err != nil {
		return err
	}

	return nil
}
