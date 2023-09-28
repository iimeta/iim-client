package dao

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/core"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/db"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var Group = NewGroupDao()

type GroupDao struct {
	*MongoDB[entity.Group]
	relation *cache.Relation
}

func NewGroupDao(database ...string) *GroupDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &GroupDao{
		MongoDB:  NewMongoDB[entity.Group](database[0], do.GROUP_COLLECTION),
		relation: cache.NewRelation(redis.Client),
	}
}

func (d *GroupDao) SearchOvertList(ctx context.Context, overt *do.SearchOvert) ([]*entity.Group, error) {

	groupIds, err := GroupMember.FindGroupIds(ctx, overt.UserId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"is_overt": 1,
		"is_dismiss": bson.M{
			"$ne": 1,
		},
	}

	if len(groupIds) > 0 {
		filter["group_id"] = bson.M{
			"$nin": groupIds,
		}
	}

	if overt.Name != "" {
		filter["group_name"] = bson.M{
			"$regex": overt.Name,
		}
	}

	groupList, err := d.FindByPage(ctx, &db.Paging{Page: overt.Page, PageSize: overt.Size}, filter, "-created_at")
	if err != nil {
		return nil, err
	}

	return groupList, nil
}

// 创建群聊
func (d *GroupDao) Create(ctx context.Context, create *do.GroupCreate) (int, error) {

	// 群成员用户ID
	uids := util.Unique(append(create.MemberIds, create.UserId))

	group := &do.Group{
		GroupId:   core.IncrGroupId(ctx),
		CreatorId: create.UserId,
		Name:      create.Name,
		Profile:   create.Profile,
		Avatar:    create.Avatar,
		MaxNum:    consts.GroupMemberMaxNum,
	}

	if _, err := d.Insert(ctx, &group); err != nil {
		return 0, err
	}

	userList, err := User.FindUserListByUserIds(ctx, create.MemberIds)
	if err != nil {
		return 0, err
	}

	addMembers := make([]*model.TalkGroupMember, 0, len(create.MemberIds))
	for _, user := range userList {
		addMembers = append(addMembers, &model.TalkGroupMember{
			UserId:   user.UserId,
			Nickname: user.Nickname,
		})
	}

	joinTime := gtime.Timestamp()
	members := make([]interface{}, 0)
	talkList := make([]interface{}, 0)

	for _, val := range uids {

		leader := 0
		if create.UserId == val {
			leader = 2
		}

		members = append(members, &do.GroupMember{
			GroupId:  group.GroupId,
			UserId:   val,
			Leader:   leader,
			JoinTime: joinTime,
		})

		talkList = append(talkList, &do.TalkSession{
			TalkType:   2,
			UserId:     val,
			ReceiverId: group.GroupId,
		})
	}

	if _, err = GroupMember.Inserts(ctx, members); err != nil {
		return 0, err
	}

	if _, err = TalkSession.Inserts(ctx, talkList); err != nil {
		return 0, err
	}

	user, err := User.FindUserByUserId(ctx, create.UserId)
	if err != nil {
		return 0, err
	}

	record := &do.TalkRecords{
		RecordId:   core.IncrRecordId(ctx),
		MsgId:      util.NewMsgId(),
		TalkType:   consts.ChatGroupMode,
		ReceiverId: group.GroupId,
		MsgType:    consts.ChatMsgSysGroupCreate,
		Sequence:   Sequence.Get(ctx, 0, group.GroupId),
		Extra: gjson.MustEncodeString(model.TalkRecordGroupCreate{
			OwnerId:   user.UserId,
			OwnerName: user.Nickname,
			Members:   addMembers,
		}),
	}

	if _, err = TalkRecords.Insert(ctx, &record); err != nil {
		return 0, err
	}

	// 广播网关将在线的用户加入房间
	body := g.Map{
		"event": consts.SubEventGroupJoin,
		"data": gjson.MustEncodeString(g.Map{
			"group_id": group.GroupId,
			"uids":     uids,
		}),
	}

	_, err = redis.Publish(ctx, consts.ImTopicChat, gjson.MustEncodeString(body))
	if err != nil {
		logger.Error(ctx, err)
	}

	return group.GroupId, err
}

// 更新群信息
func (d *GroupDao) Update(ctx context.Context, update *do.GroupUpdate) error {

	if err := d.UpdateOne(ctx, bson.M{"group_id": update.GroupId}, bson.M{
		"group_name": update.Name,
		"avatar":     update.Avatar,
		"profile":    update.Profile,
	}); err != nil {
		return err
	}

	return nil
}

// 解散群聊[群主权限]
func (d *GroupDao) Dismiss(ctx context.Context, groupId int, uid int) error {

	if err := d.UpdateOne(ctx, bson.M{"group_id": groupId, "creator_id": uid}, bson.M{
		"is_dismiss": 1,
	}); err != nil {
		return err
	}

	if err := GroupMember.UpdateMany(ctx, bson.M{"group_id": groupId}, bson.M{
		"is_quit": 1,
	}); err != nil {
		return err
	}

	return nil
}

// 退出群聊[仅管理员及群成员]
func (d *GroupDao) Secede(ctx context.Context, groupId int, uid int) error {

	groupMember, err := GroupMember.FindByUserId(ctx, groupId, uid)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("群聊不存在或已被解散")
		}
		return err
	}

	if groupMember.Leader == 2 {
		return errors.New("群主不能退出群聊")
	}

	user, err := User.FindUserByUserId(ctx, uid)
	if err != nil {
		return err
	}

	record := &do.TalkRecords{
		RecordId:   core.IncrRecordId(ctx),
		MsgId:      util.NewMsgId(),
		TalkType:   consts.ChatGroupMode,
		ReceiverId: groupId,
		MsgType:    consts.ChatMsgSysGroupMemberQuit,
		Sequence:   Sequence.Get(ctx, 0, groupId),
		Extra: gjson.MustEncodeString(&model.TalkRecordGroupMemberQuit{
			OwnerId:   user.UserId,
			OwnerName: user.Nickname,
		}),
	}

	if err = GroupMember.UpdateOne(ctx, bson.M{"group_id": groupId, "user_id": uid}, bson.M{
		"is_quit": 1,
	}); err != nil {
		return err
	}

	_, err = TalkRecords.Insert(ctx, record)
	if err != nil {
		return err
	}

	d.relation.DelGroupRelation(ctx, uid, groupId)

	if _, err = redis.Publish(ctx, consts.ImTopicChat, gjson.MustEncodeString(g.Map{
		"event": consts.SubEventGroupJoin,
		"data": gjson.MustEncodeString(g.Map{
			"type":     2,
			"group_id": groupId,
			"uids":     []int{uid},
		}),
	})); err != nil {
		return err
	}

	if _, err = redis.Publish(ctx, consts.ImTopicChat, gjson.MustEncodeString(g.Map{
		"event": consts.SubEventImMessage,
		"data": gjson.MustEncodeString(g.Map{
			"sender_id":   record.UserId,
			"receiver_id": record.ReceiverId,
			"talk_type":   record.TalkType,
			"record_id":   record.RecordId,
		}),
	})); err != nil {
		return err
	}

	return nil
}

// 邀请加入群聊
func (d *GroupDao) Invite(ctx context.Context, invite *do.GroupInvite) error {

	addMembers := make([]interface{}, 0)
	addTalkList := make([]interface{}, 0)
	updateTalkList := make([]int, 0)

	m := make(map[int]struct{})
	for _, value := range GroupMember.GetMemberIds(ctx, invite.GroupId) {
		m[value] = struct{}{}
	}

	talkList, err := TalkSession.Find(ctx, bson.M{"user_id": bson.M{"$in": invite.MemberIds}, "receiver_id": invite.GroupId, "talk_type": 2})
	if err != nil {
		return err
	}

	listHash := make(map[int]*entity.TalkSession)
	for _, item := range talkList {
		listHash[item.UserId] = item
	}

	mids := make([]int, 0)
	mids = append(mids, invite.MemberIds...)
	mids = append(mids, invite.UserId)

	memberItems, err := User.FindUserListByUserIds(ctx, mids)
	if err != nil {
		return err
	}

	memberMaps := make(map[int]*entity.User)
	for _, item := range memberItems {
		memberMaps[item.UserId] = item
	}

	members := make([]*model.TalkGroupMember, 0)
	for _, value := range invite.MemberIds {
		members = append(members, &model.TalkGroupMember{
			UserId:   value,
			Nickname: memberMaps[value].Nickname,
		})

		if _, ok := m[value]; !ok {
			addMembers = append(addMembers, &do.GroupMember{
				GroupId:  invite.GroupId,
				UserId:   value,
				JoinTime: gtime.Timestamp(),
			})
		}

		if item, ok := listHash[value]; !ok {
			addTalkList = append(addTalkList, &do.TalkSession{
				TalkType:   consts.ChatGroupMode,
				UserId:     value,
				ReceiverId: invite.GroupId,
			})
		} else if item.IsDelete == 1 {
			updateTalkList = append(updateTalkList, item.UserId)
		}
	}

	if len(addMembers) == 0 {
		return errors.New("邀请的好友, 都已成为群成员")
	}

	record := &do.TalkRecords{
		RecordId:   core.IncrRecordId(ctx),
		MsgId:      util.NewMsgId(),
		TalkType:   consts.ChatGroupMode,
		ReceiverId: invite.GroupId,
		MsgType:    consts.ChatMsgSysGroupMemberJoin,
		Sequence:   Sequence.Get(ctx, 0, invite.GroupId),
	}

	record.Extra = gjson.MustEncodeString(&model.TalkRecordGroupJoin{
		OwnerId:   memberMaps[invite.UserId].UserId,
		OwnerName: memberMaps[invite.UserId].Nickname,
		Members:   members,
	})

	// 删除已存在成员记录
	if _, err = GroupMember.DeleteMany(ctx, bson.M{"group_id": invite.GroupId, "user_id": bson.M{"$in": invite.MemberIds}, "is_quit": consts.GroupMemberQuitStatusYes}); err != nil {
		return err
	}

	if _, err = GroupMember.Inserts(ctx, addMembers); err != nil {
		return err
	}

	// 添加用户的对话列表
	if len(addTalkList) > 0 {
		if _, err = TalkSession.Inserts(ctx, addTalkList); err != nil {
			return err
		}
	}

	// 更新用户的对话列表
	if len(updateTalkList) > 0 {
		if err = TalkSession.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": updateTalkList}}, bson.M{
			"is_delete":  0,
			"created_at": gtime.Datetime(),
		}); err != nil {
			return err
		}
	}

	_, err = TalkRecords.Insert(ctx, record)
	if err != nil {
		return err
	}

	// 广播网关将在线的用户加入房间
	if _, err = redis.Publish(ctx, consts.ImTopicChat, gjson.MustEncodeString(g.Map{
		"event": consts.SubEventGroupJoin,
		"data": gjson.MustEncodeString(g.Map{
			"type":     1,
			"group_id": invite.GroupId,
			"uids":     invite.MemberIds,
		}),
	})); err != nil {
		return err
	}

	if _, err = redis.Publish(ctx, consts.ImTopicChat, gjson.MustEncodeString(g.Map{
		"event": consts.SubEventImMessage,
		"data": gjson.MustEncodeString(g.Map{
			"sender_id":   record.UserId,
			"receiver_id": record.ReceiverId,
			"talk_type":   record.TalkType,
			"record_id":   record.RecordId,
		}),
	})); err != nil {
		return err
	}

	return nil
}

// 群成员移除群聊
func (d *GroupDao) RemoveMember(ctx context.Context, remove *do.GroupMemberRemove) error {

	num, err := GroupMember.CountDocuments(ctx, bson.M{"group_id": remove.GroupId, "user_id": bson.M{"$in": remove.MemberIds}, "is_quit": 0})
	if err != nil {
		return err
	}

	if int(num) != len(remove.MemberIds) {
		return errors.New("删除失败")
	}

	mids := make([]int, 0)
	mids = append(mids, remove.MemberIds...)
	mids = append(mids, remove.UserId)

	memberItems, err := User.FindUserListByUserIds(ctx, mids)
	if err != nil {
		return err
	}

	memberMaps := make(map[int]*entity.User)
	for _, item := range memberItems {
		memberMaps[item.UserId] = item
	}

	members := make([]*model.TalkGroupMember, 0)
	for _, value := range remove.MemberIds {
		members = append(members, &model.TalkGroupMember{
			UserId:   value,
			Nickname: memberMaps[value].Nickname,
		})
	}

	record := &do.TalkRecords{
		RecordId:   core.IncrRecordId(ctx),
		MsgId:      util.NewMsgId(),
		Sequence:   Sequence.Get(ctx, 0, remove.GroupId),
		TalkType:   consts.ChatGroupMode,
		ReceiverId: remove.GroupId,
		MsgType:    consts.ChatMsgSysGroupMemberKicked,
		Extra: gjson.MustEncodeString(&model.TalkRecordGroupMemberKicked{
			OwnerId:   memberMaps[remove.UserId].UserId,
			OwnerName: memberMaps[remove.UserId].Nickname,
			Members:   members,
		}),
	}

	if err = GroupMember.UpdateMany(ctx, bson.M{"group_id": remove.GroupId, "user_id": bson.M{"$in": remove.MemberIds}, "is_quit": 0}, bson.M{
		"is_quit":    1,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}

	_, err = TalkRecords.Insert(ctx, record)
	if err != nil {
		return err
	}

	d.relation.BatchDelGroupRelation(ctx, remove.MemberIds, remove.GroupId)

	pipe := redis.Pipeline(ctx)

	pipe.Publish(ctx, consts.ImTopicChat, gjson.MustEncodeString(g.Map{
		"event": consts.SubEventGroupJoin,
		"data": gjson.MustEncodeString(g.Map{
			"type":     2,
			"group_id": remove.GroupId,
			"uids":     remove.MemberIds,
		}),
	}))

	pipe.Publish(ctx, consts.ImTopicChat, gjson.MustEncodeString(g.Map{
		"event": consts.SubEventImMessage,
		"data": gjson.MustEncodeString(g.Map{
			"sender_id":   int64(record.UserId),
			"receiver_id": int64(record.ReceiverId),
			"talk_type":   record.TalkType,
			"record_id":   record.RecordId,
		}),
	}))

	_, _ = redis.Pipelined(ctx, pipe)

	return nil
}

func (d *GroupDao) List(ctx context.Context, userId int) ([]*entity.Group, []*entity.GroupMember, []*entity.TalkSession, error) {

	groupMemberList, err := GroupMember.FindGroupMemberListByUserId(ctx, userId)
	if err != nil {
		return nil, nil, nil, err
	}

	groupIds := make([]int, 0)
	for _, member := range groupMemberList {
		groupIds = append(groupIds, member.GroupId)
	}

	groupList, err := d.FindGroupListByGroupIds(ctx, groupIds)
	if err != nil {
		return nil, nil, nil, err
	}

	ids := make([]int, 0)
	for _, group := range groupList {
		ids = append(ids, group.GroupId)
	}

	talkSessionList, err := TalkSession.Find(ctx, bson.M{"talk_type": 2, "receiver_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, nil, nil, err
	}

	return groupList, groupMemberList, talkSessionList, nil
}

// 根据groupIds查询群列表
func (d *GroupDao) FindGroupListByGroupIds(ctx context.Context, groupIds []int) ([]*entity.Group, error) {

	groupList, err := d.Find(ctx, bson.M{"group_id": bson.M{"$in": groupIds}})
	if err != nil {
		return nil, err
	}

	return groupList, nil
}

// 根据groupId查询群信息
func (d *GroupDao) FindGroupByGroupId(ctx context.Context, groupId int) (*entity.Group, error) {

	group, err := d.FindOne(ctx, bson.M{"group_id": groupId})
	if err != nil {
		return nil, err
	}

	return group, nil
}
