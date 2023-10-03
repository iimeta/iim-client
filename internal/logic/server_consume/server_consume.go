package server_consume

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/socket"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

type sServerConsume struct {
	handlers      map[string]func(ctx context.Context, data []byte)
	clientStorage *cache.ClientStorage
	roomStorage   *cache.RoomStorage
}

func init() {
	service.RegisterServerConsume(New())
}

func New() service.IServerConsume {
	return &sServerConsume{
		clientStorage: cache.NewClientStorage(redis.Client, config.Cfg, cache.NewSidStorage(redis.Client)),
		roomStorage:   cache.NewRoomStorage(redis.Client),
	}
}

// 触发回调事件
func (s *sServerConsume) Call(ctx context.Context, event string, data []byte) {

	if s.handlers == nil {
		s.init()
	}

	if call, ok := s.handlers[event]; ok {
		call(ctx, data)
	} else {
		logger.Errorf(ctx, "consume chat event: [%s]未注册回调事件", event)
	}
}

func (s *sServerConsume) init() {

	s.handlers = make(map[string]func(ctx context.Context, data []byte))

	s.handlers[consts.SubEventImMessage] = s.onConsumeTalk
	s.handlers[consts.SubEventImMessageKeyboard] = s.onConsumeTalkKeyboard
	s.handlers[consts.SubEventImMessageRead] = s.onConsumeTalkRead
	s.handlers[consts.SubEventImMessageRevoke] = s.onConsumeTalkRevoke
	s.handlers[consts.SubEventContactStatus] = s.onConsumeContactStatus
	s.handlers[consts.SubEventContactApply] = s.onConsumeContactApply
	s.handlers[consts.SubEventGroupJoin] = s.onConsumeGroupJoin
	s.handlers[consts.SubEventGroupApply] = s.onConsumeGroupApply
}

// 用户上线或下线消息
func (s *sServerConsume) onConsumeContactStatus(ctx context.Context, body []byte) {

	var in model.ConsumeContactStatus
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Error(ctx, "[ChatSubscribe] onConsumeContactStatus Unmarshal err: ", err.Error())
		return
	}

	contactIds := dao.Contact.GetContactIds(ctx, in.UserId)

	clientIds := make([]int64, 0)
	for _, uid := range util.Unique(contactIds) {
		ids := s.clientStorage.GetUidFromClientIds(ctx, config.Cfg.ServerId(), socket.Session.Chat.Name(), gconv.String(uid))
		if len(ids) > 0 {
			clientIds = append(clientIds, ids...)
		}
	}

	if len(clientIds) == 0 {
		return
	}

	c := socket.NewSenderContent()
	c.SetReceive(clientIds...)
	c.SetMessage(consts.PushEventContactStatus, in)

	socket.Session.Chat.Write(c)
}

// 好友申请消息
func (s *sServerConsume) onConsumeContactApply(ctx context.Context, body []byte) {

	var in model.ConsumeContactApply
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Error(gctx.New(), "[ChatSubscribe] onConsumeContactApply Unmarshal err: ", err.Error())
		return
	}

	apply := new(entity.ContactApply)
	if err := dao.FindById(ctx, dao.Contact.Database, do.CONTACT_APPLY_COLLECTION, in.ApplyId, &apply); err != nil {
		return
	}

	clientIds := s.clientStorage.GetUidFromClientIds(ctx, config.Cfg.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(apply.FriendId))
	if len(clientIds) == 0 {
		return
	}

	user, err := dao.User.FindUserByUserId(ctx, apply.FriendId)
	if err != nil {
		return
	}

	data := g.Map{}
	data["sender_id"] = apply.UserId
	data["receiver_id"] = apply.FriendId
	data["remark"] = apply.Remark
	data["friend"] = g.Map{
		"nickname":   user.Nickname,
		"remark":     apply.Remark,
		"created_at": util.FormatDatetime(apply.CreatedAt),
	}

	c := socket.NewSenderContent()
	c.SetAck(true)
	c.SetReceive(clientIds...)
	c.SetMessage(consts.PushEventContactApply, data)

	socket.Session.Chat.Write(c)
}

// 加入群房间
func (s *sServerConsume) onConsumeGroupJoin(ctx context.Context, body []byte) {

	var in model.ConsumeGroupJoin
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Error(ctx, "[ChatSubscribe] onConsumeGroupJoin Unmarshal err: ", err.Error())
		return
	}

	sid := config.Cfg.ServerId()
	for _, uid := range in.Uids {
		ids := s.clientStorage.GetUidFromClientIds(ctx, sid, socket.Session.Chat.Name(), strconv.Itoa(uid))

		for _, cid := range ids {
			opt := &cache.RoomOption{
				Channel:  socket.Session.Chat.Name(),
				RoomType: consts.RoomImGroup,
				Number:   strconv.Itoa(in.Gid),
				Sid:      config.Cfg.ServerId(),
				Cid:      cid,
			}

			if in.Type == 2 {
				_ = s.roomStorage.Del(ctx, opt)
			} else {
				_ = s.roomStorage.Add(ctx, opt)
			}
		}
	}
}

// 入群申请通知
func (s *sServerConsume) onConsumeGroupApply(ctx context.Context, body []byte) {

	var in model.ConsumeGroupApply
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Error(ctx, "[ChatSubscribe] onConsumeGroupApply Unmarshal err: ", err.Error())
		return
	}

	groupMember, err := dao.GroupMember.FindOne(ctx, bson.M{"group_id": in.GroupId, "leader": 2})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	groupDetail, err := dao.Group.FindGroupByGroupId(ctx, in.GroupId)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	user, err := dao.User.FindUserByUserId(ctx, in.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	data := make(g.Map)
	data["group_name"] = groupDetail.GroupName
	data["nickname"] = user.Nickname

	clientIds := s.clientStorage.GetUidFromClientIds(ctx, config.Cfg.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(groupMember.UserId))

	c := socket.NewSenderContent()
	c.SetReceive(clientIds...)
	c.SetMessage(consts.PushEventGroupApply, data)

	socket.Session.Chat.Write(c)
}

// 键盘输入事件消息
func (s *sServerConsume) onConsumeTalkKeyboard(ctx context.Context, body []byte) {

	var in model.ConsumeTalkKeyboard
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Error(ctx, "[ChatSubscribe] onConsumeTalkKeyboard Unmarshal err: ", err.Error())
		return
	}

	ids := s.clientStorage.GetUidFromClientIds(ctx, config.Cfg.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(in.ReceiverID))
	if len(ids) == 0 {
		return
	}

	c := socket.NewSenderContent()
	c.SetReceive(ids...)
	c.SetMessage(consts.PushEventImMessageKeyboard, g.Map{
		"sender_id":   in.SenderID,
		"receiver_id": in.ReceiverID,
	})

	socket.Session.Chat.Write(c)
}

// 聊天消息事件
func (s *sServerConsume) onConsumeTalk(ctx context.Context, body []byte) {

	var in model.ConsumeTalk
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Error(ctx, "[ChatSubscribe] onConsumeTalk Unmarshal err: ", err.Error())
		return
	}

	var clientIds []int64
	if in.TalkType == consts.ChatPrivateMode {
		for _, val := range [2]int{in.SenderID, in.ReceiverID} {
			ids := s.clientStorage.GetUidFromClientIds(ctx, config.Cfg.ServerId(), socket.Session.Chat.Name(), gconv.String(val))

			clientIds = append(clientIds, ids...)
		}
	} else if in.TalkType == consts.ChatGroupMode {
		ids := s.roomStorage.All(ctx, &cache.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: consts.RoomImGroup,
			Number:   gconv.String(in.ReceiverID),
			Sid:      config.Cfg.ServerId(),
		})

		clientIds = append(clientIds, ids...)
	}

	if len(clientIds) == 0 {
		return
	}

	data, err := service.TalkRecords().GetTalkRecord(ctx, in.RecordID)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	c := socket.NewSenderContent()
	c.SetReceive(clientIds...)
	c.SetAck(true)
	c.SetMessage(consts.PushEventImMessage, g.Map{
		"sender_id":   in.SenderID,
		"receiver_id": in.ReceiverID,
		"talk_type":   in.TalkType,
		"data":        data,
	})

	socket.Session.Chat.Write(c)
}

// 消息已读事件
func (s *sServerConsume) onConsumeTalkRead(ctx context.Context, body []byte) {

	var in model.ConsumeTalkRead
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Error(ctx, "[ChatSubscribe] onConsumeContactApply Unmarshal err: ", err.Error())
		return
	}

	clientIds := s.clientStorage.GetUidFromClientIds(ctx, config.Cfg.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(in.ReceiverId))
	if len(clientIds) == 0 {
		return
	}

	c := socket.NewSenderContent()
	c.SetAck(true)
	c.SetReceive(clientIds...)
	c.SetMessage(consts.PushEventImMessageRead, g.Map{
		"sender_id":   in.SenderId,
		"receiver_id": in.ReceiverId,
		"ids":         in.Ids,
	})

	socket.Session.Chat.Write(c)
}

// 撤销聊天消息
func (s *sServerConsume) onConsumeTalkRevoke(ctx context.Context, body []byte) {

	var in model.ConsumeTalkRevoke
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Error(ctx, "[ChatSubscribe] onConsumeTalkRevoke Unmarshal err: ", err.Error())
		return
	}

	record, err := dao.TalkRecords.FindByRecordId(ctx, in.RecordId)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	var clientIds []int64
	if record.TalkType == consts.ChatPrivateMode {
		for _, uid := range [2]int{record.UserId, record.ReceiverId} {
			ids := s.clientStorage.GetUidFromClientIds(ctx, config.Cfg.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(uid))
			clientIds = append(clientIds, ids...)
		}
	} else if record.TalkType == consts.ChatGroupMode {
		clientIds = s.roomStorage.All(ctx, &cache.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: consts.RoomImGroup,
			Number:   strconv.Itoa(record.ReceiverId),
			Sid:      config.Cfg.ServerId(),
		})
	}

	if len(clientIds) == 0 {
		return
	}

	c := socket.NewSenderContent()
	c.SetAck(true)
	c.SetReceive(clientIds...)
	c.SetMessage(consts.PushEventImMessageRevoke, g.Map{
		"talk_type":   record.TalkType,
		"sender_id":   record.UserId,
		"receiver_id": record.ReceiverId,
		"record_id":   record.RecordId,
	})

	socket.Session.Chat.Write(c)
}
