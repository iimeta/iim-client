package server_event

import (
	"context"
	"encoding/json"
	gjson2 "github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/socket"
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

type sServerEvent struct {
	handlers    map[string]func(ctx context.Context, client socket.IClient, data []byte)
	RoomStorage *cache.RoomStorage
}

func init() {
	service.RegisterServerEvent(New())
}

func New() service.IServerEvent {
	return &sServerEvent{
		RoomStorage: cache.NewRoomStorage(redis.Client),
	}
}

func (s *sServerEvent) Call(ctx context.Context, client socket.IClient, event string, data []byte) {

	if s.handlers == nil {
		s.init()
	}

	if call, ok := s.handlers[event]; ok {
		call(ctx, client, data)
	} else {
		logger.Errorf(ctx, "Chat Event: [%s]未注册回调事件", event)
	}
}

func (s *sServerEvent) init() {

	s.handlers = make(map[string]func(ctx context.Context, client socket.IClient, data []byte))

	// 注册自定义绑定事件
	s.handlers["im.message.publish"] = s.onPublish
	s.handlers["im.message.revoke"] = s.onRevokeMessage
	s.handlers["im.message.delete"] = s.onDeleteMessage
	s.handlers["im.message.read"] = s.onReadMessage
	s.handlers["im.message.keyboard"] = s.onKeyboardMessage
}

// 连接成功回调事件
func (s *sServerEvent) OnOpen(client socket.IClient) {

	ctx := gctx.New()

	// 1.查询用户群列表
	ids := dao.GroupMember.GetUserGroupIds(ctx, client.Uid())

	// 2.客户端加入群房间
	rooms := make([]*cache.RoomOption, 0, len(ids))
	for _, id := range ids {
		rooms = append(rooms, &cache.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: consts.RoomImGroup,
			Number:   strconv.Itoa(id),
			Sid:      config.Cfg.ServerId(),
			Cid:      client.Cid(),
		})
	}

	if err := s.RoomStorage.BatchAdd(ctx, rooms); err != nil {
		logger.Error(ctx, "加入群聊失败", err)
	}

	// 推送上线消息
	redis.Client.Publish(ctx, consts.ImTopicChat, gjson2.MustEncodeString(map[string]any{
		"event": consts.SubEventContactStatus,
		"data": gjson2.MustEncodeString(map[string]any{
			"user_id": client.Uid(),
			"status":  1,
		}),
	}))
}

// 消息回调事件
func (s *sServerEvent) OnMessage(client socket.IClient, message []byte) {

	// 获取事件名 todo 更改替换gjson
	event := gjson.GetBytes(message, "event").String()
	if event != "" {
		// 触发事件
		s.Call(gctx.New(), client, event, message)
	}
}

// 连接关闭回调事件
func (s *sServerEvent) OnClose(client socket.IClient, code int, text string) {

	ctx := gctx.New()

	logger.Infof(ctx, "client close uid: %d, cid: %d, channel: %s, code: %d, text: %s", client.Uid(), client.Cid(), client.Channel().Name(), code, text)

	// 1.判断用户是否是多点登录

	// 2.查询用户群列表
	ids := dao.GroupMember.GetUserGroupIds(ctx, client.Uid())

	// 3.客户端退出群房间
	rooms := make([]*cache.RoomOption, 0, len(ids))
	for _, id := range ids {
		rooms = append(rooms, &cache.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: consts.RoomImGroup,
			Number:   strconv.Itoa(id),
			Sid:      config.Cfg.ServerId(),
			Cid:      client.Cid(),
		})
	}

	if err := s.RoomStorage.BatchDel(ctx, rooms); err != nil {
		logger.Error(ctx, "退出群聊失败", err)
	}

	// 推送下线消息
	redis.Client.Publish(
		ctx,
		consts.ImTopicChat,
		gjson2.MustEncodeString(map[string]any{
			"event": consts.SubEventContactStatus,
			"data": gjson2.MustEncodeString(map[string]any{
				"user_id": client.Uid(),
				"status":  0,
			}),
		}),
	)
}

// 键盘输入事件
func (s *sServerEvent) onKeyboardMessage(ctx context.Context, _ socket.IClient, data []byte) {

	var in model.KeyboardMessage
	if err := json.Unmarshal(data, &in); err != nil {
		logger.Error(ctx, "Chat onKeyboardMessage Err:", err)
		return
	}

	redis.Client.Publish(ctx, consts.ImTopicChat, gjson2.MustEncodeString(map[string]any{
		"event": consts.SubEventImMessageKeyboard,
		"data": gjson2.MustEncodeString(map[string]any{
			"sender_id":   in.Content.SenderID,
			"receiver_id": in.Content.ReceiverID,
		}),
	}))
}

func (s *sServerEvent) onPublish(ctx context.Context, client socket.IClient, data []byte) {

	if s.handlers == nil {
		s.handlers = make(map[string]func(ctx context.Context, client socket.IClient, data []byte))
		s.handlers["text"] = s.onTextMessage
		s.handlers["code"] = s.onCodeMessage
		s.handlers["location"] = s.onLocationMessage
		s.handlers["emoticon"] = s.onEmoticonMessage
		s.handlers["vote"] = s.onVoteMessage
		s.handlers["image"] = s.onImageMessage
		s.handlers["file"] = s.onFileMessage
	}

	typeValue := gjson.GetBytes(data, "content.type").String()
	if call, ok := s.handlers[typeValue]; ok {
		call(ctx, client, data)
	} else {
		logger.Errorf(ctx, "chat event: onPublish [%s]未知的消息类型", typeValue)
	}
}

// 文本消息
func (s *sServerEvent) onTextMessage(ctx context.Context, client socket.IClient, data []byte) {

	var in model.TextMessage
	if err := json.Unmarshal(data, &in); err != nil {
		logger.Error(ctx, "Chat onTextMessage Err:", err)
		return
	}

	if in.Content.Content == "" {
		return
	}

	err := service.TalkMessage().SendText(ctx, client.Uid(), &model.TextMessageReq{
		Content: in.Content.Content,
		Receiver: &model.MessageReceiver{
			TalkType:   in.Content.Receiver.TalkType,
			ReceiverId: in.Content.Receiver.ReceiverId,
		},
	})

	if err != nil {
		logger.Error(ctx, "Chat onTextMessage err:", err)
		return
	}

	if len(in.AckId) == 0 {
		return
	}

	if err = client.Write(&socket.ClientResponse{Sid: in.AckId, Event: "ack"}); err != nil {
		logger.Error(ctx, "Chat onTextMessage ack err:", err)
	}
}

// 代码消息
func (s *sServerEvent) onCodeMessage(ctx context.Context, client socket.IClient, data []byte) {

	var m model.CodeMessage
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(ctx, "Chat onTextMessage Err:", err)
		return
	}

	if m.Content.Receiver == nil {
		return
	}

	err := service.TalkMessage().SendCode(ctx, client.Uid(), &model.CodeMessageReq{
		Lang: m.Content.Lang,
		Code: m.Content.Code,
		Receiver: &model.MessageReceiver{
			TalkType:   m.Content.Receiver.TalkType,
			ReceiverId: m.Content.Receiver.ReceiverId,
		},
	})

	if err != nil {
		logger.Error(ctx, "Chat onTextMessage err:", err)
		return
	}

	if len(m.AckId) == 0 {
		return
	}

	if err = client.Write(&socket.ClientResponse{Sid: m.AckId, Event: "ack"}); err != nil {
		logger.Error(ctx, "Chat onTextMessage ack err:", err)
	}
}

// 表情包消息
func (s *sServerEvent) onEmoticonMessage(ctx context.Context, _ socket.IClient, data []byte) {

	var m model.EmoticonMessage
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(ctx, "Chat onEmoticonMessage Err:", err)
		return
	}

	logger.Debug(ctx, "[onEmoticonMessage] 新消息 ", string(data))
}

// 图片消息
func (s *sServerEvent) onImageMessage(ctx context.Context, _ socket.IClient, data []byte) {

	var m model.ImageMessage
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(ctx, "Chat onImageMessage Err:", err)
		return
	}

	logger.Debug(ctx, "[onImageMessage] 新消息 ", string(data))
}

// 文件消息
func (s *sServerEvent) onFileMessage(ctx context.Context, _ socket.IClient, data []byte) {

	var m model.FileMessage
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(ctx, "Chat onFileMessage Err:", err)
		return
	}

	logger.Debug(ctx, "[onFileMessage] 新消息 ", string(data))
}

// 位置消息
func (s *sServerEvent) onLocationMessage(ctx context.Context, _ socket.IClient, data []byte) {

	var m model.LocationMessage
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(ctx, "Chat onLocationMessage Err:", err)
		return
	}

	logger.Debug(ctx, "[onLocationMessage] 新消息 ", string(data))
}

// 投票消息
func (s *sServerEvent) onVoteMessage(ctx context.Context, _ socket.IClient, data []byte) {

	var m model.VoteMessage
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(ctx, "Chat onVoteMessage Err:", err)
		return
	}

	logger.Debug(ctx, "[onVoteMessage] 新消息 ", string(data))
}

// 消息已读事件
func (s *sServerEvent) onReadMessage(ctx context.Context, client socket.IClient, data []byte) {

	var in model.TalkReadMessage
	if err := gjson2.Unmarshal(data, &in); err != nil {
		logger.Error(ctx, "Chat onReadMessage Err:", err)
		return
	}

	if err := dao.TalkRecords.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": in.Content.MsgIds}, "receiver_id": client.Uid(), "is_read": 0}, bson.M{
		"is_read": 1,
	}); err != nil {
		logger.Error(ctx, err)
		return
	}

	redis.Client.Publish(ctx, consts.ImTopicChat, gjson2.MustEncodeString(map[string]any{
		"event": consts.SubEventImMessageRead,
		"data": gjson2.MustEncodeString(map[string]any{
			"sender_id":   client.Uid(),
			"receiver_id": in.Content.ReceiverId,
			"ids":         in.Content.MsgIds,
		}),
	}))
}

func (s *sServerEvent) onRevokeMessage(ctx context.Context, client socket.IClient, data []byte) {

}

func (s *sServerEvent) onDeleteMessage(ctx context.Context, client socket.IClient, data []byte) {

}
