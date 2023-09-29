package group

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"slices"
)

type sGroup struct {
	RedisLock *cache.RedisLock
}

func init() {
	service.RegisterGroup(New())
}

func New() service.IGroup {
	return &sGroup{
		RedisLock: cache.NewRedisLock(redis.Client),
	}
}

func (s *sGroup) GroupAuth(ctx context.Context, auth *model.GroupAuth) error {

	// 判断对方是否是自己
	if auth.TalkType == consts.ChatPrivateMode && auth.ReceiverId == service.Session().GetUid(ctx) {
		return nil
	}

	if auth.TalkType == consts.ChatPrivateMode {
		if dao.Contact.IsFriend(ctx, auth.UserId, auth.ReceiverId, false) {
			return nil
		}
		return errors.New("暂无权限发送消息")
	}

	groupInfo, err := dao.Group.FindGroupByGroupId(ctx, auth.ReceiverId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if groupInfo.IsDismiss == 1 {
		return errors.New("此群聊已解散")
	}

	memberInfo, err := dao.GroupMember.FindByUserId(ctx, auth.ReceiverId, auth.UserId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("暂无权限发送消息")
		}

		logger.Error(ctx, err)
		return errors.New("系统繁忙, 请稍后再试")
	}

	if memberInfo.IsQuit == consts.GroupMemberQuitStatusYes {
		return errors.New("暂无权限发送消息")
	}

	if memberInfo.IsMute == consts.GroupMemberMuteStatusYes {
		return errors.New("已被群主或管理员禁言")
	}

	if auth.IsVerifyGroupMute && groupInfo.IsMute == 1 && memberInfo.Leader == 0 {
		return errors.New("此群聊已开启全员禁言")
	}

	return nil
}

// 创建群聊分组
func (s *sGroup) Create(ctx context.Context, params model.GroupCreateReq) (*model.GroupCreateRes, error) {

	gid, err := dao.Group.Create(ctx, &do.GroupCreate{
		UserId:    service.Session().GetUid(ctx),
		Name:      params.Name,
		Avatar:    params.Avatar,
		MemberIds: util.ParseIds(params.Ids),
	})
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("创建群聊失败, 请稍后再试" + err.Error())
	}

	return &model.GroupCreateRes{GroupId: gid}, nil
}

// 解散群聊
func (s *sGroup) Dismiss(ctx context.Context, params model.GroupDismissReq) error {

	uid := service.Session().GetUid(ctx)
	if !dao.GroupMember.IsMaster(ctx, params.GroupId, uid) {
		return errors.New("暂无权限解散群聊")
	}

	if err := dao.Group.Dismiss(ctx, params.GroupId, service.Session().GetUid(ctx)); err != nil {
		logger.Error(ctx, err)
		return errors.New("群聊解散失败")
	}

	_ = service.TalkMessage().SendSystemText(ctx, uid, &model.TextMessageReq{
		Content: "群聊已被群主解散",
		Receiver: &model.Receiver{
			TalkType:   consts.ChatGroupMode,
			ReceiverId: params.GroupId,
		},
	})

	return nil
}

// 邀请好友加入群聊
func (s *sGroup) Invite(ctx context.Context, params model.GroupInviteReq) error {

	key := fmt.Sprintf("group-join:%d", params.GroupId)
	if !s.RedisLock.Lock(ctx, key, 20) {
		return errors.New("网络异常, 请稍后再试")
	}

	defer s.RedisLock.UnLock(ctx, key)

	group, err := dao.Group.FindGroupByGroupId(ctx, params.GroupId)
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("网络异常, 请稍后再试")
	}

	if group != nil && group.IsDismiss == 1 {
		return errors.New("该群已解散")
	}

	uid := service.Session().GetUid(ctx)
	uids := util.Unique(util.ParseIds(params.Ids))

	if len(uids) == 0 {
		return errors.New("邀请好友列表不能为空")
	}

	if !dao.GroupMember.IsMember(ctx, params.GroupId, uid, true) {
		return errors.New("非群聊成员, 无权邀请好友")
	}

	if err := dao.Group.Invite(ctx, &do.GroupInvite{
		UserId:    uid,
		GroupId:   params.GroupId,
		MemberIds: uids,
	}); err != nil {
		logger.Error(ctx, err)
		return errors.New("邀请好友加入群聊失败" + err.Error())
	}

	return nil
}

// 退出群聊
func (s *sGroup) Secede(ctx context.Context, params model.GroupSecedeReq) error {

	uid := service.Session().GetUid(ctx)
	if err := dao.Group.Secede(ctx, params.GroupId, uid); err != nil {
		logger.Error(ctx, err)
		return err
	}

	sid, err := dao.TalkSession.FindBySessionId(ctx, uid, params.GroupId, consts.ChatGroupMode)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.TalkSession.Delete(ctx, service.Session().GetUid(ctx), sid); err != nil {
		logger.Error(ctx, err)
		return err
	}
	return nil
}

// 群设置接口(预留)
func (s *sGroup) Setting(ctx context.Context, params model.GroupSettingReq) error {

	group, err := dao.Group.FindGroupByGroupId(ctx, params.GroupId)
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("网络异常, 请稍后再试")
	}

	if group != nil && group.IsDismiss == 1 {
		return errors.New("该群已解散")
	}

	uid := service.Session().GetUid(ctx)
	if !dao.GroupMember.IsLeader(ctx, params.GroupId, uid) {
		return errors.New("无权限操作")
	}

	if err := dao.Group.Update(ctx, &do.GroupUpdate{
		GroupId: params.GroupId,
		Name:    params.GroupName,
		Avatar:  params.Avatar,
		Profile: params.Profile,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	_ = service.TalkMessage().SendSystemText(ctx, uid, &model.TextMessageReq{
		Content: "群主或管理员修改了群信息",
		Receiver: &model.Receiver{
			TalkType:   consts.ChatGroupMode,
			ReceiverId: params.GroupId,
		},
	})

	return nil
}

// 移除指定成员(群聊&管理员权限)
func (s *sGroup) RemoveMembers(ctx context.Context, params model.GroupRemoveMemberReq) error {

	uid := service.Session().GetUid(ctx)

	if !dao.GroupMember.IsLeader(ctx, params.GroupId, uid) {
		return errors.New("无权限操作")
	}

	if err := dao.Group.RemoveMember(ctx, &do.GroupMemberRemove{
		UserId:    uid,
		GroupId:   params.GroupId,
		MemberIds: util.ParseIds(params.MembersIds),
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 获取群聊信息
func (s *sGroup) Detail(ctx context.Context, params model.GroupDetailReq) (*model.GroupDetailRes, error) {

	uid := service.Session().GetUid(ctx)

	groupInfo, err := dao.Group.FindGroupByGroupId(ctx, params.GroupId)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(ctx, err)
		return nil, err
	}

	if groupInfo == nil || groupInfo.GroupId == 0 {
		return nil, errors.New("数据不存在")
	}

	resp := &model.GroupDetailRes{
		GroupId:   groupInfo.GroupId,
		GroupName: groupInfo.GroupName,
		Profile:   groupInfo.Profile,
		Avatar:    groupInfo.Avatar,
		CreatedAt: util.FormatDatetime(groupInfo.CreatedAt),
		IsManager: uid == groupInfo.CreatorId,
		IsDisturb: 0,
		IsMute:    groupInfo.IsMute,
		IsOvert:   groupInfo.IsOvert,
		VisitCard: dao.GroupMember.GetMemberRemark(ctx, params.GroupId, uid),
	}

	if dao.TalkSession.IsDisturb(ctx, uid, groupInfo.GroupId, 2) {
		resp.IsDisturb = 1
	}

	return resp, nil
}

// 修改群备注接口
func (s *sGroup) UpdateMemberRemark(ctx context.Context, params model.GroupRemarkUpdateReq) error {

	if err := dao.GroupMember.UpdateOne(ctx, bson.M{"group_id": params.GroupId, "user_id": service.Session().GetUid(ctx)}, bson.M{
		"user_card": params.VisitCard,
	}); err != nil {
		logger.Error(ctx, err)
		return errors.New("修改群备注失败")
	}

	return nil
}

func (s *sGroup) GetInviteFriends(ctx context.Context, params model.GetInviteFriendsReq) ([]*model.ContactListItem, error) {

	contactList, userList, err := dao.Contact.List(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userMap := util.ToMap(userList, func(t *entity.User) int {
		return t.UserId
	})

	items := make([]*model.ContactListItem, 0)
	for _, contact := range contactList {
		items = append(items, &model.ContactListItem{
			Id:       contact.FriendId,
			Nickname: userMap[contact.FriendId].Nickname,
			Gender:   userMap[contact.FriendId].Gender,
			Motto:    userMap[contact.FriendId].Motto,
			Avatar:   userMap[contact.FriendId].Avatar,
			Remark:   contact.Remark,
			GroupId:  contact.GroupId,
		})
	}

	if params.GroupId <= 0 {
		return items, nil
	}

	mids := dao.GroupMember.GetMemberIds(ctx, params.GroupId)
	if len(mids) == 0 {
		return items, nil
	}

	data := make([]*model.ContactListItem, 0)
	for i := 0; i < len(items); i++ {
		if !slices.Contains(mids, items[i].Id) {
			data = append(data, items[i])
		}
	}

	return data, nil
}

func (s *sGroup) List(ctx context.Context) (*model.GroupListRes, error) {

	groupList, groupMemberList, talkSessionList, err := dao.Group.List(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	groupMemberMap := make(map[int]*entity.GroupMember)
	for _, member := range groupMemberList {
		groupMemberMap[member.GroupId] = member
	}

	talkSessionMap := make(map[int]*entity.TalkSession)
	for _, session := range talkSessionList {
		talkSessionMap[session.ReceiverId] = session
	}

	items := make([]*model.Group, 0)
	for _, group := range groupList {

		groupItem := &model.Group{
			Id:        group.GroupId,
			GroupName: group.GroupName,
			Avatar:    group.Avatar,
			Profile:   group.Profile,
			CreatorId: group.CreatorId,
		}

		if groupMemberMap[group.GroupId] != nil {
			groupItem.Leader = groupMemberMap[group.GroupId].Leader
		}

		if talkSessionMap[group.GroupId] != nil {
			groupItem.IsDisturb = talkSessionMap[group.GroupId].IsDisturb
		}

		items = append(items, groupItem)
	}

	resp := &model.GroupListRes{}
	for _, item := range items {
		resp.Items = append(resp.Items, &model.Group{
			Id:        item.Id,
			GroupName: item.GroupName,
			Avatar:    item.Avatar,
			Profile:   item.Profile,
			Leader:    item.Leader,
			IsDisturb: item.IsDisturb,
			CreatorId: item.CreatorId,
		})
	}

	return resp, nil
}

// 获取群成员列表
func (s *sGroup) Members(ctx context.Context, params model.GroupMemberListReq) (*model.GroupMemberListRes, error) {

	group, err := dao.Group.FindGroupByGroupId(ctx, params.GroupId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("网络异常, 请稍后再试")
	}

	if group != nil && group.IsDismiss == 1 {
		return &model.GroupMemberListRes{}, nil
	}

	if !dao.GroupMember.IsMember(ctx, params.GroupId, service.Session().GetUid(ctx), false) {
		return nil, errors.New("非群成员无权查看成员列表")
	}

	groupMemberList, userList, err := dao.GroupMember.GetMembers(ctx, params.GroupId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userMap := make(map[int]*entity.User)
	for _, user := range userList {
		userMap[user.UserId] = user
	}

	items := make([]*model.GroupMember, 0)
	for _, member := range groupMemberList {
		items = append(items, &model.GroupMember{
			UserId:   member.UserId,
			Leader:   member.Leader,
			IsMute:   member.IsMute,
			Remark:   member.UserCard,
			Nickname: userMap[member.UserId].Nickname,
			Avatar:   userMap[member.UserId].Avatar,
			Gender:   userMap[member.UserId].Gender,
		})
	}

	return &model.GroupMemberListRes{Items: items}, nil
}

// 公开群列表
func (s *sGroup) OvertList(ctx context.Context, params model.GroupOvertListReq) (*model.GroupOvertListRes, error) {

	uid := service.Session().GetUid(ctx)

	list, err := dao.Group.SearchOvertList(ctx, &do.SearchOvert{
		Name:   params.Name,
		UserId: uid,
		Page:   int64(params.Page),
		Size:   20,
	})
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("查询异常")
	}

	resp := &model.GroupOvertListRes{}
	resp.Items = make([]*model.GroupOvert, 0)

	if len(list) == 0 {
		return resp, nil
	}

	ids := make([]int, 0)
	for _, val := range list {
		ids = append(ids, val.GroupId)
	}

	count, err := dao.GroupMember.CountGroupMemberNum(ctx, ids)
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("查询异常")
	}

	countMap := make(map[int]int)
	for _, member := range count {
		countMap[member.GroupId] = member.Count
	}

	for i, value := range list {
		if i >= 19 {
			break
		}

		resp.Items = append(resp.Items, &model.GroupOvert{
			Id:        value.GroupId,
			Type:      value.Type,
			Name:      value.GroupName,
			Avatar:    value.Avatar,
			Profile:   value.Profile,
			Count:     countMap[value.GroupId],
			MaxNum:    value.MaxNum,
			CreatedAt: util.FormatDatetime(value.CreatedAt),
		})
	}

	resp.Next = len(list) > 19

	return resp, nil
}

// 群主交接
func (s *sGroup) Handover(ctx context.Context, params model.GroupHandoverReq) error {

	uid := service.Session().GetUid(ctx)
	if !dao.GroupMember.IsMaster(ctx, params.GroupId, uid) {
		return errors.New("暂无权限")
	}

	if uid == params.UserId {
		return errors.New("暂无权限")
	}

	if err := dao.GroupMember.Handover(ctx, params.GroupId, uid, params.UserId); err != nil {
		logger.Error(ctx)
		return errors.New("转让群主失败")
	}

	members, err := dao.User.FindUserListByUserIds(ctx, []int{uid, params.UserId})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	extra := model.TalkRecordGroupTransfer{}
	for _, member := range members {
		if member.UserId == uid {
			extra.OldOwnerId = member.UserId
			extra.OldOwnerName = member.Nickname
		} else {
			extra.NewOwnerId = member.UserId
			extra.NewOwnerName = member.Nickname
		}
	}

	_ = service.TalkMessage().SendSysOther(ctx, &model.TalkRecord{
		MsgType:    consts.ChatMsgSysGroupTransfer,
		TalkType:   consts.TalkRecordTalkTypeGroup,
		UserId:     uid,
		ReceiverId: params.GroupId,
		Extra:      gjson.MustEncodeString(extra),
	})

	return nil
}

// 分配管理员
func (s *sGroup) AssignAdmin(ctx context.Context, params model.GroupAssignAdminReq) error {

	if !dao.GroupMember.IsMaster(ctx, params.GroupId, service.Session().GetUid(ctx)) {
		return errors.New("暂无权限")
	}

	leader := 0
	if params.Mode == 1 {
		leader = 1
	}

	if err := dao.GroupMember.SetLeaderStatus(ctx, params.GroupId, params.UserId, leader); err != nil {
		logger.Error(ctx)
		return errors.New("设置管理员信息失败")
	}

	return nil
}

// 禁止发言
func (s *sGroup) NoSpeak(ctx context.Context, params model.GroupNoSpeakReq) error {

	uid := service.Session().GetUid(ctx)
	if !dao.GroupMember.IsLeader(ctx, params.GroupId, uid) {
		return errors.New("暂无权限")
	}

	status := 1
	if params.Mode == 2 {
		status = 0
	}

	if err := dao.GroupMember.SetMuteStatus(ctx, params.GroupId, params.UserId, status); err != nil {
		logger.Error(ctx)
		return errors.New("设置群成员禁言状态失败")
	}

	data := &model.TalkRecord{
		TalkType:   consts.TalkRecordTalkTypeGroup,
		UserId:     uid,
		ReceiverId: params.GroupId,
	}

	members := make([]*model.TalkGroupMember, 0)
	if user, err := dao.User.FindUserByUserId(ctx, params.UserId); err != nil {
		logger.Error(ctx, err)
		return err
	} else {
		members = append(members, &model.TalkGroupMember{
			UserId:   user.UserId,
			Nickname: user.Nickname,
		})
	}

	user, err := dao.User.FindUserByUserId(ctx, uid)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if status == 1 {
		data.MsgType = consts.ChatMsgSysGroupMemberMuted
		data.Extra = gjson.MustEncodeString(model.TalkRecordGroupMemberCancelMuted{
			OwnerId:   uid,
			OwnerName: user.Nickname,
			Members:   members,
		})
	} else {
		data.MsgType = consts.ChatMsgSysGroupMemberCancelMuted
		data.Extra = gjson.MustEncodeString(model.TalkRecordGroupMemberCancelMuted{
			OwnerId:   uid,
			OwnerName: user.Nickname,
			Members:   members,
		})
	}

	_ = service.TalkMessage().SendSysOther(ctx, data)

	return nil
}

// 全员禁言
func (s *sGroup) Mute(ctx context.Context, params model.GroupMuteReq) error {

	uid := service.Session().GetUid(ctx)

	group, err := dao.Group.FindGroupByGroupId(ctx, params.GroupId)
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("网络异常, 请稍后再试")
	}

	if group.IsDismiss == 1 {
		return errors.New("此群已解散")
	}

	if !dao.GroupMember.IsLeader(ctx, params.GroupId, uid) {
		return errors.New("暂无权限")
	}

	data := make(map[string]any)
	if params.Mode == 1 {
		data["is_mute"] = 1
	} else {
		data["is_mute"] = 0
	}

	if err := dao.Group.UpdateOne(ctx, bson.M{"group_id": params.GroupId}, data); err != nil {
		logger.Error(ctx, err)
		return errors.New("服务器异常, 请稍后再试")
	}

	user, err := dao.User.FindUserByUserId(ctx, uid)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	var extra any
	var msgType int
	if params.Mode == 1 {
		msgType = consts.ChatMsgSysGroupMuted
		extra = model.TalkRecordGroupMuted{
			OwnerId:   user.UserId,
			OwnerName: user.Nickname,
		}
	} else {
		msgType = consts.ChatMsgSysGroupCancelMuted
		extra = model.TalkRecordGroupCancelMuted{
			OwnerId:   user.UserId,
			OwnerName: user.Nickname,
		}
	}

	_ = service.TalkMessage().SendSysOther(ctx, &model.TalkRecord{
		MsgType:    msgType,
		TalkType:   consts.TalkRecordTalkTypeGroup,
		UserId:     uid,
		ReceiverId: params.GroupId,
		Extra:      gjson.MustEncodeString(extra),
	})

	return nil
}

// 公开群
func (s *sGroup) Overt(ctx context.Context, params model.GroupOvertReq) error {

	uid := service.Session().GetUid(ctx)

	group, err := dao.Group.FindGroupByGroupId(ctx, params.GroupId)
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("网络异常, 请稍后再试")
	}

	if group.IsDismiss == 1 {
		return errors.New("此群已解散")
	}

	if !dao.GroupMember.IsMaster(ctx, params.GroupId, uid) {
		return errors.New("暂无权限")
	}

	data := make(map[string]any)
	if params.Mode == 1 {
		data["is_overt"] = 1
	} else {
		data["is_overt"] = 0
	}

	if err := dao.Group.UpdateOne(ctx, bson.M{"group_id": params.GroupId}, data); err != nil {
		logger.Error(ctx, err)
		return errors.New("服务器异常, 请稍后再试")
	}

	return nil
}
