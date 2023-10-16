package talk_message

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/core"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/filesystem"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/robot"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"html"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

type sTalkMessage struct {
	unreadStorage  *cache.UnreadStorage
	messageStorage *cache.MessageStorage
	sidStorage     *cache.ServerStorage
	clientStorage  *cache.ClientStorage
	Filesystem     *filesystem.Filesystem
	mapping        map[string]func(ctx context.Context) error
}

func init() {
	service.RegisterTalkMessage(New())
}

func New() service.ITalkMessage {
	return &sTalkMessage{
		unreadStorage:  cache.NewUnreadStorage(redis.Client),
		messageStorage: cache.NewMessageStorage(redis.Client),
		sidStorage:     cache.NewSidStorage(redis.Client),
		clientStorage:  cache.NewClientStorage(redis.Client, config.Cfg, cache.NewSidStorage(redis.Client)),
		Filesystem:     filesystem.NewFilesystem(config.Cfg),
	}
}

// 校验权限
func (s *sTalkMessage) VerifyPermission(ctx context.Context, info *model.VerifyInfo) error {

	// 判断对方是否是自己
	if info.TalkType == consts.ChatPrivateMode && info.ReceiverId == info.UserId {
		return nil
	}

	// 判断发送人是否为机器人 todo
	if robot, err := dao.Robot.GetRobotByUserId(ctx, info.UserId); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(ctx, err)
		return err
	} else if robot != nil {
		return nil
	}

	if info.TalkType == consts.ChatPrivateMode {
		if dao.Contact.IsFriend(ctx, info.UserId, info.ReceiverId, false) {
			return nil
		}
		return errors.New("暂无权限发送消息")
	}

	groupInfo, err := dao.Group.FindGroupByGroupId(ctx, info.ReceiverId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if groupInfo.IsDismiss == 1 {
		return errors.New("此群聊已解散")
	}

	memberInfo, err := dao.GroupMember.FindByUserId(ctx, info.ReceiverId, info.UserId)
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

	if info.IsVerifyGroupMute && groupInfo.IsMute == 1 && memberInfo.Leader == 0 {
		return errors.New("此群聊已开启全员禁言")
	}

	return nil
}

// 发送消息
func (s *sTalkMessage) SendMessage(ctx context.Context, message *model.Message) (err error) {

	if err = s.VerifyPermission(ctx, &model.VerifyInfo{
		TalkType:          message.TalkType,
		UserId:            message.Sender.Id,
		ReceiverId:        message.Receiver.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	var data *model.TalkRecord

	switch message.MsgType {
	case consts.MsgTypeText:
		data, err = s.TextMessageHandler(ctx, message)
	case consts.MsgTypeCode:
		data, err = s.CodeMessageHandler(ctx, message)
	case consts.MsgTypeImage:
		data, err = s.ImageMessageHandler(ctx, message)
	case consts.MsgTypeVoice:
		data, err = s.VoiceMessageHandler(ctx, message)
	case consts.MsgTypeVideo:
		data, err = s.VideoMessageHandler(ctx, message)
	case consts.MsgTypeFile:
		data, err = s.FileMessageHandler(ctx, message)
	case consts.MsgTypeVote:
		data, err = s.VoteMessageHandler(ctx, message)
	case consts.MsgTypeMixed:
		data, err = s.MixedMessageHandler(ctx, message)
	case consts.MsgTypeForward:
		data, err = s.ForwardMessageHandler(ctx, message)
		return err
	case consts.MsgTypeEmoticon:
		data, err = s.EmoticonMessageHandler(ctx, message)
	case consts.MsgTypeCard:
		data, err = s.CardMessageHandler(ctx, message)
	case consts.MsgTypeLocation:
		data, err = s.LocationMessageHandler(ctx, message)
	default:
		return errors.New("未知消息类型")
	}

	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return s.save(ctx, data)
}

// 发送系统消息
func (s *sTalkMessage) SendSysMessage(ctx context.Context, message *model.SysMessage) error {

	var data *model.TalkRecord

	switch message.MsgType {
	case consts.MsgSysText:
		data = &model.TalkRecord{
			TalkType:   message.TalkType,
			MsgType:    consts.ChatMsgSysText,
			UserId:     message.Sender.Id,
			ReceiverId: message.Receiver.ReceiverId,
			Content:    html.EscapeString(message.Text.Content),
			Text:       message.Text,
			Sender:     message.Sender,
			Receiver:   message.Receiver,
		}
	case consts.MsgSysGroupCreate:
	case consts.MsgSysGroupMemberJoin:
	case consts.MsgSysGroupMemberQuit:
	case consts.MsgSysGroupMemberKicked:
	case consts.MsgSysGroupMessageRevoke:
	case consts.MsgSysGroupDismissed:
	case consts.MsgSysGroupMuted:
		data = &model.TalkRecord{
			MsgType:    consts.ChatMsgSysGroupMuted,
			TalkType:   message.TalkType,
			UserId:     message.Sender.Id,
			ReceiverId: message.Receiver.ReceiverId,
			Extra:      gjson.MustEncodeString(message.GroupMuted),
			GroupMuted: message.GroupMuted,
			Sender:     message.Sender,
			Receiver:   message.Receiver,
		}
	case consts.MsgSysGroupCancelMuted:
		data = &model.TalkRecord{
			MsgType:          consts.ChatMsgSysGroupCancelMuted,
			TalkType:         message.TalkType,
			UserId:           message.Sender.Id,
			ReceiverId:       message.Receiver.ReceiverId,
			Extra:            gjson.MustEncodeString(message.GroupCancelMuted),
			GroupCancelMuted: message.GroupCancelMuted,
			Sender:           message.Sender,
			Receiver:         message.Receiver,
		}
	case consts.MsgSysGroupMemberMuted:
		data = &model.TalkRecord{
			MsgType:          consts.ChatMsgSysGroupMemberMuted,
			TalkType:         message.TalkType,
			UserId:           message.Sender.Id,
			ReceiverId:       message.Receiver.ReceiverId,
			Extra:            gjson.MustEncodeString(message.GroupMemberMuted),
			GroupMemberMuted: message.GroupMemberMuted,
			Sender:           message.Sender,
			Receiver:         message.Receiver,
		}
	case consts.MsgSysGroupMemberCancelMuted:
		data = &model.TalkRecord{
			MsgType:                consts.ChatMsgSysGroupMemberCancelMuted,
			TalkType:               message.TalkType,
			UserId:                 message.Sender.Id,
			ReceiverId:             message.Receiver.ReceiverId,
			Extra:                  gjson.MustEncodeString(message.GroupMemberCancelMuted),
			GroupMemberCancelMuted: message.GroupMemberCancelMuted,
			Sender:                 message.Sender,
			Receiver:               message.Receiver,
		}
	case consts.MsgSysGroupNotice:
		data = &model.TalkRecord{
			MsgType:     consts.ChatMsgSysGroupNotice,
			TalkType:    message.TalkType,
			UserId:      message.Sender.Id,
			ReceiverId:  message.Receiver.ReceiverId,
			Extra:       gjson.MustEncodeString(message.GroupNotice),
			GroupNotice: message.GroupNotice,
			Sender:      message.Sender,
			Receiver:    message.Receiver,
		}
	case consts.MsgSysGroupTransfer:
		data = &model.TalkRecord{
			MsgType:       consts.ChatMsgSysGroupTransfer,
			TalkType:      message.TalkType,
			UserId:        message.Sender.Id,
			ReceiverId:    message.Receiver.ReceiverId,
			Extra:         gjson.MustEncodeString(message.GroupTransfer),
			GroupTransfer: message.GroupTransfer,
			Sender:        message.Sender,
			Receiver:      message.Receiver,
		}
	default:
		return errors.New("未知消息类型")
	}

	return s.save(ctx, data)
}

// 发送通知消息
func (s *sTalkMessage) SendNoticeMessage(ctx context.Context, message *model.NoticeMessage) error {

	var data *model.TalkRecord
	switch message.MsgType {
	case consts.MsgNoticeLogin:
		data = &model.TalkRecord{
			TalkType:   message.TalkType,
			MsgType:    consts.ChatMsgTypeLogin,
			UserId:     message.Sender.Id,
			ReceiverId: message.Receiver.ReceiverId,
			Extra: gjson.MustEncodeString(&model.TalkRecordLogin{
				IP:       message.Login.IP,
				Platform: message.Login.Platform,
				Agent:    message.Login.Agent,
				Address:  message.Login.Address,
				Reason:   message.Login.Reason,
				Datetime: gtime.Datetime(),
			}),
			Login:    message.Login,
			Sender:   message.Sender,
			Receiver: message.Receiver,
		}
	default:
		return errors.New("未知消息类型")
	}

	return s.save(ctx, data)
}

// 文件消息
func (s *sTalkMessage) SendFile(ctx context.Context, uid int, req *model.MessageFileReq) error {

	file, err := dao.SplitUpload.GetFile(ctx, uid, req.UploadId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	publicUrl := ""
	filePath := fmt.Sprintf("private/files/talks/%s/%s.%s", util.DateNumber(), gmd5.MustEncryptString(util.Random(16)), file.FileExt)

	// 公开文件
	if consts.GetMediaType(file.FileExt) <= 3 {
		filePath = fmt.Sprintf("public/media/%s/%s.%s", util.DateNumber(), gmd5.MustEncryptString(util.Random(16)), file.FileExt)
		publicUrl = filesystem.NewFilesystem(config.Cfg).Default.PublicUrl(filePath)
	}

	if err := filesystem.NewFilesystem(config.Cfg).Default.Copy(file.Path, filePath); err != nil {
		logger.Error(ctx, err)
		return err
	}

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
	}

	switch consts.GetMediaType(file.FileExt) {
	case consts.MediaFileAudio:
		data.MsgType = consts.ChatMsgTypeAudio
		data.Extra = gjson.MustEncodeString(&model.TalkRecordAudio{
			Suffix:   file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
			Duration: 0,
		})
	case consts.MediaFileVideo:
		data.MsgType = consts.ChatMsgTypeVideo
		data.Extra = gjson.MustEncodeString(&model.TalkRecordVideo{
			Cover:    "",
			Suffix:   file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
			Duration: 0,
		})
	case consts.MediaFileOther:
		data.MsgType = consts.ChatMsgTypeFile
		data.Extra = gjson.MustEncodeString(&model.TalkRecordFile{
			Drive:  file.Drive,
			Name:   file.OriginalName,
			Suffix: file.FileExt,
			Size:   int(file.FileSize),
			Path:   filePath,
		})
	}

	return s.save(ctx, data)
}

// 投票消息
func (s *sTalkMessage) SendVote(ctx context.Context, uid int, req *model.MessageVoteReq) error {

	data := &model.TalkRecord{
		RecordId:   core.IncrRecordId(ctx),
		MsgId:      util.NewMsgId(),
		TalkType:   consts.ChatGroupMode,
		MsgType:    consts.ChatMsgTypeVote,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
	}

	s.loadSequence(ctx, data)

	answerOptions := make([]*model.AnswerOption, 0)
	options := make(map[string]string)
	for i, value := range req.Options {
		options[fmt.Sprintf("%c", 65+i)] = value
		answerOptions = append(answerOptions, &model.AnswerOption{
			Key:   fmt.Sprintf("%c", 65+i),
			Value: value,
		})
	}

	num := dao.GroupMember.CountMemberTotal(ctx, req.Receiver.ReceiverId)

	_, err := dao.TalkRecords.Insert(ctx, &do.TalkRecords{
		RecordId:   data.RecordId,
		MsgId:      data.MsgId,
		Sequence:   data.Sequence,
		TalkType:   data.TalkType,
		MsgType:    data.MsgType,
		UserId:     data.UserId,
		ReceiverId: data.ReceiverId,
		Vote: &model.Vote{
			Title:         req.Title,
			AnswerMode:    req.Mode,
			Anonymous:     req.Anonymous,
			AnswerOptions: answerOptions,
			AnswerNum:     int(num),
		},
		Sender: &model.Sender{
			Id: data.UserId,
		},
		Receiver: &model.Receiver{
			Id:         data.ReceiverId,
			ReceiverId: data.ReceiverId,
		},
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.TalkRecordsVote.Insert(ctx, &do.TalkRecordsVote{
		RecordId:     data.RecordId,
		UserId:       uid,
		Title:        req.Title,
		AnswerMode:   req.Mode,
		AnswerOption: gjson.MustEncodeString(options),
		AnswerNum:    int(num),
		IsAnonymous:  req.Anonymous,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	s.afterHandle(ctx, data, map[string]string{"text": "[投票消息]"})

	return nil
}

// 撤回消息
func (s *sTalkMessage) Revoke(ctx context.Context, params model.MessageRevokeReq) error {

	uid := service.Session().GetUid(ctx)

	record, err := dao.TalkRecords.FindByRecordId(ctx, params.RecordId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if record.IsRevoke == 1 {
		return nil
	}

	if record.UserId != uid {
		return errors.New("无权撤回回消息")
	}

	if time.Now().Unix() > gtime.NewFromTimeStamp(record.CreatedAt).Add(3*time.Minute).Unix() {
		return errors.New("超出有效撤回时间范围, 无法进行撤销")
	}

	if err := dao.TalkRecords.UpdateOne(ctx, bson.M{"record_id": params.RecordId}, bson.M{"is_revoke": 1}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	body := map[string]any{
		"event": consts.SubEventImMessageRevoke,
		"data": gjson.MustEncodeString(map[string]any{
			"record_id": record.RecordId,
		}),
	}

	_, err = redis.Publish(ctx, consts.ImTopicChat, gjson.MustEncodeString(body))
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (s *sTalkMessage) save(ctx context.Context, data *model.TalkRecord) error {

	if data.RecordId == 0 {
		data.RecordId = core.IncrRecordId(ctx)
	}

	if data.MsgId == "" {
		data.MsgId = util.NewMsgId()
	}

	s.loadReply(ctx, data)

	s.loadSequence(ctx, data)

	id, err := dao.TalkRecords.Insert(ctx, &do.TalkRecords{
		RecordId:   data.RecordId,
		MsgId:      data.MsgId,
		Sequence:   data.Sequence,
		TalkType:   data.TalkType,
		MsgType:    data.MsgType,
		UserId:     data.UserId,
		ReceiverId: data.ReceiverId,
		IsRevoke:   data.IsRevoke,
		IsMark:     data.IsMark,
		IsRead:     data.IsRead,
		QuoteId:    data.QuoteId,
		Content:    data.Content,
		Extra:      data.Extra,

		Sender:   data.Sender,
		Receiver: data.Receiver,
		Mention:  data.Mention,
		Reply:    data.Reply,

		Text:     data.Text,
		Code:     data.Code,
		Image:    data.Image,
		Voice:    data.Voice,
		Video:    data.Video,
		File:     data.File,
		Vote:     data.Vote,
		Mixed:    data.Mixed,
		Emoticon: data.Emoticon,
		Card:     data.Card,
		Location: data.Location,

		GroupCreate:            data.GroupCreate,
		GroupJoin:              data.GroupJoin,
		GroupTransfer:          data.GroupTransfer,
		GroupMuted:             data.GroupMuted,
		GroupCancelMuted:       data.GroupCancelMuted,
		GroupMemberMuted:       data.GroupMemberMuted,
		GroupMemberCancelMuted: data.GroupMemberCancelMuted,
		GroupDismissed:         data.GroupDismissed,
		GroupMemberQuit:        data.GroupMemberQuit,
		GroupMemberKicked:      data.GroupMemberKicked,
		GroupMessageRevoke:     data.GroupMessageRevoke,
		GroupNotice:            data.GroupNotice,

		Login: data.Login,
	})

	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	data.Id = id

	option := make(map[string]string)

	switch data.MsgType {
	case consts.ChatMsgTypeText:
		option["text"] = gstr.SubStr(util.ReplaceImgAll(data.Content), 0, 300)
	default:
		if value, ok := consts.ChatMsgTypeMapping[data.MsgType]; ok {
			option["text"] = value
		} else {
			option["text"] = "[未知消息]"
		}
	}

	s.afterHandle(ctx, data, option)

	return nil
}

func (s *sTalkMessage) loadSequence(ctx context.Context, data *model.TalkRecord) {
	if data.TalkType == consts.ChatGroupMode {
		data.Sequence = dao.Sequence.Get(ctx, 0, data.ReceiverId)
	} else {
		data.Sequence = dao.Sequence.Get(ctx, data.UserId, data.ReceiverId)
	}
}

func (s *sTalkMessage) loadReply(ctx context.Context, data *model.TalkRecord) {

	// 检测是否引用消息
	if data.QuoteId == "" {
		return
	}

	if data.Extra == "" {
		data.Extra = "{}"
	}

	extra := make(map[string]any)

	err := gjson.Unmarshal([]byte(data.Extra), &extra)
	if err != nil {
		logger.Error(ctx, "MessageService Json Decode err:", err)
		return
	}

	record, err := dao.TalkRecords.FindOne(ctx, bson.M{"msg_id": data.QuoteId})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	user, err := dao.User.FindUserByUserId(ctx, record.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	reply := model.TalkRecordReply{
		UserId:   record.UserId,
		Nickname: user.Nickname,
		MsgType:  1,
		Content:  record.Content,
		MsgId:    record.MsgId,
	}

	if record.MsgType != consts.ChatMsgTypeText {
		reply.Content = "[未知消息]"
		if value, ok := consts.ChatMsgTypeMapping[record.MsgType]; ok {
			reply.Content = value
		}
	}

	extra["reply"] = reply

	data.Extra = gjson.MustEncodeString(extra)

	data.Reply = &model.Reply{
		UserId:   record.UserId,
		Nickname: user.Nickname,
		MsgType:  consts.MsgTypeText,
		Content:  record.Content,
		MsgId:    record.MsgId,
	}
}

// 发送消息后置处理
func (s *sTalkMessage) afterHandle(ctx context.Context, record *model.TalkRecord, opt map[string]string) {

	if record.TalkType == consts.ChatPrivateMode {
		s.unreadStorage.Incr(ctx, consts.ChatPrivateMode, record.UserId, record.ReceiverId)
		if record.MsgType == consts.ChatMsgSysText {
			s.unreadStorage.Incr(ctx, 1, record.ReceiverId, record.UserId)
		}
	} else if record.TalkType == consts.ChatGroupMode {
		pipe := redis.Pipeline(ctx)
		for _, uid := range dao.GroupMember.GetMemberIds(ctx, record.ReceiverId) {
			if uid != record.UserId {
				s.unreadStorage.PipeIncr(ctx, pipe, consts.ChatGroupMode, record.ReceiverId, uid)
			}
		}
		if _, err := pipe.Exec(ctx); err != nil {
			logger.Error(ctx, err)
		}
	}

	if err := s.messageStorage.Set(ctx, record.TalkType, record.UserId, record.ReceiverId, &cache.LastCacheMessage{
		Content:  opt["text"],
		Datetime: gtime.Datetime(),
	}); err != nil {
		logger.Error(ctx, err)
	}

	content := gjson.MustEncodeString(map[string]any{
		"event": consts.SubEventImMessage,
		"data": gjson.MustEncodeString(map[string]any{
			"sender_id":   record.UserId,
			"receiver_id": record.ReceiverId,
			"talk_type":   record.TalkType,
			"record_id":   record.RecordId,
		}),
	})

	if record.TalkType == consts.ChatPrivateMode {
		sids := s.sidStorage.All(ctx, 1)

		if len(sids) > 3 {

			pipe := redis.Pipeline(ctx)

			for _, sid := range sids {
				for _, uid := range []int{record.UserId, record.ReceiverId} {
					if !s.clientStorage.IsCurrentServerOnline(ctx, sid, consts.ImChannelChat, strconv.Itoa(uid)) {
						continue
					}
					pipe.Publish(ctx, fmt.Sprintf(consts.ImTopicChatPrivate, sid), content)
				}
			}

			if _, err := pipe.Exec(ctx); err != nil {
				logger.Error(ctx, err)
				return
			}
		}
	}

	if _, err := redis.Publish(ctx, consts.ImTopicChat, content); err != nil {
		logger.Error(ctx, "[ALL]消息推送失败 err:", err)
	}
}

// 发送图片消息
func (s *sTalkMessage) Image(ctx context.Context, params model.ImageMessageReq) error {

	_, file, err := g.RequestFromCtx(ctx).Request.FormFile("image")
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("image 字段必传")
	}

	if !slices.Contains([]string{"png", "jpg", "jpeg", "gif", "webp"}, gfile.ExtName(file.Filename)) {
		return errors.New("上传文件格式不正确,仅支持 png、jpg、jpeg、gif 及 webp")
	}

	// 判断上传文件大小(20M)
	if file.Size > 20<<20 {
		return errors.New("上传文件大小不能超过20M")
	}

	if err = s.VerifyPermission(ctx, &model.VerifyInfo{
		TalkType:          params.TalkType,
		UserId:            service.Session().GetUid(ctx),
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	stream, err := filesystem.ReadMultipartStream(file)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	ext := gfile.ExtName(file.Filename)

	meta := util.ReadImageMeta(bytes.NewReader(stream))

	filePath := fmt.Sprintf("public/media/image/talk/%s/%s", util.DateNumber(), util.GenImageName(ext, meta.Width, meta.Height))

	if err = s.Filesystem.Default.Write(stream, filePath); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = s.SendMessage(ctx, &model.Message{
		MsgType:  consts.MsgTypeImage,
		TalkType: params.TalkType,
		Image: &model.Image{
			Url:    s.Filesystem.Default.PublicUrl(filePath),
			Width:  meta.Width,
			Height: meta.Height,
			Size:   int(file.Size),
		},
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			TalkType:   params.TalkType,
			ReceiverId: params.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 发送文件消息
func (s *sTalkMessage) File(ctx context.Context, params model.MessageFileReq) error {

	uid := service.Session().GetUid(ctx)
	if err := s.VerifyPermission(ctx, &model.VerifyInfo{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	//if err := s.SendFile(ctx, uid, &model.MessageFileReq{
	//	UploadId: params.UploadId,
	//	Receiver: &model.Receiver{
	//		TalkType:   params.TalkType,
	//		ReceiverId: params.ReceiverId,
	//	},
	//}); err != nil {
	//	logger.Error(ctx, err)
	//	return err
	//}

	if err := s.onSendFile(ctx, &model.MessageFileReq{
		UploadId: params.UploadId,
		Receiver: &model.Receiver{
			TalkType:   params.TalkType,
			ReceiverId: params.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 发送投票消息
func (s *sTalkMessage) Vote(ctx context.Context, params model.MessageVoteReq) error {

	if len(params.Options) <= 1 {
		return errors.New("options 选项必须大于1")
	}

	if len(params.Options) > 6 {
		return errors.New("options 选项不能超过6个")
	}

	uid := service.Session().GetUid(ctx)
	if err := s.VerifyPermission(ctx, &model.VerifyInfo{
		TalkType:          consts.ChatGroupMode,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := s.SendVote(ctx, uid, &model.MessageVoteReq{
		Mode:      params.Mode,
		Title:     params.Title,
		Options:   params.Options,
		Anonymous: params.Anonymous,
		Receiver: &model.Receiver{
			TalkType:   consts.ChatGroupMode,
			ReceiverId: params.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 投票处理
func (s *sTalkMessage) HandleVote(ctx context.Context, params model.MessageVoteHandleReq) (*model.VoteStatistics, error) {

	uid := service.Session().GetUid(ctx)

	talkRecords, err := dao.TalkRecords.FindByRecordId(ctx, params.RecordId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	talkRecordsVote, err := dao.TalkRecordsVote.FindOne(ctx, bson.M{"record_id": talkRecords.RecordId})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if talkRecords.MsgType != consts.ChatMsgTypeVote {
		return nil, errors.Newf("当前记录不属于投票信息[%d]", talkRecords.MsgType)
	}

	if talkRecords.TalkType == consts.ChatGroupMode {
		count, err := dao.GroupMember.CountDocuments(ctx, bson.M{"group_id": talkRecords.ReceiverId, "user_id": uid, "is_quit": 0})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if count == 0 {
			return nil, errors.New("暂无投票权限")
		}
	}

	count, err := dao.CountDocuments(ctx, dao.TalkRecordsVote.Database, do.TALK_RECORDS_VOTE_ANSWER_COLLECTION, bson.M{"vote_id": talkRecordsVote.Id, "user_id": uid})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if count > 0 {
		return nil, errors.Newf("重复投票[%d]", talkRecordsVote.Id)
	}

	options := strings.Split(params.Options, ",")
	sort.Strings(options)

	var answerOptions map[string]any
	if err := gjson.Unmarshal([]byte(talkRecordsVote.AnswerOption), &answerOptions); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, option := range options {
		if _, ok := answerOptions[option]; !ok {
			return nil, errors.Newf("投票选项不合法[%s]", option)
		}
	}

	if talkRecordsVote.AnswerMode == consts.VoteAnswerModeSingleChoice {
		options = options[:1]
	}

	answers := make([]interface{}, 0, len(options))
	for _, option := range options {
		answers = append(answers, &do.TalkRecordsVoteAnswer{
			VoteId: talkRecordsVote.Id,
			UserId: uid,
			Option: option,
		})
	}

	if err = dao.TalkRecordsVote.UpdateById(ctx, talkRecordsVote.Id, bson.M{
		"$inc": bson.M{
			"answered_num": 1,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if err = dao.TalkRecords.UpdateById(ctx, talkRecords.Id, bson.M{
		"$inc": bson.M{
			"vote.answered_num": 1,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if talkRecordsVote, err := dao.TalkRecordsVote.FindById(ctx, talkRecordsVote.Id); err != nil {
		logger.Error(ctx, err)
		return nil, err
	} else if talkRecordsVote.AnsweredNum >= talkRecordsVote.AnswerNum {

		if err = dao.TalkRecordsVote.UpdateById(ctx, talkRecordsVote.Id, bson.M{
			"status": 1,
		}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if err = dao.TalkRecords.UpdateById(ctx, talkRecords.Id, bson.M{
			"vote.status": 1,
		}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if _, err = dao.Inserts(ctx, dao.TalkRecordsVote.Database, answers); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if _, err = dao.TalkRecordsVote.SetVoteAnswerUser(ctx, talkRecordsVote.Id); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if _, err = dao.TalkRecordsVote.SetVoteStatistics(ctx, talkRecordsVote.Id); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	info, err := dao.TalkRecordsVote.GetVoteStatistics(ctx, talkRecordsVote.Id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return info, nil
}

// 删除消息记录
func (s *sTalkMessage) Delete(ctx context.Context, params model.MessageDeleteReq) error {

	if err := dao.TalkRecords.DeleteRecord(ctx, &do.RemoveRecord{
		UserId:     service.Session().GetUid(ctx),
		TalkType:   params.TalkType,
		ReceiverId: params.ReceiverId,
		RecordIds:  params.RecordIds,
	}); err != nil {
		logger.Error(ctx)
		return err
	}

	return nil
}

// 验证转发消息合法性
func (s *sTalkMessage) Verify(ctx context.Context, uid int, params *model.ForwardMessageReq) error {

	filter := bson.M{
		"record_id": bson.M{
			"$in": params.MessageIds,
		},
		"talk_type": params.Receiver.TalkType,
		"msg_type": bson.M{
			"$nin": []int{consts.ChatMsgTypeLogin, consts.ChatMsgTypeVote},
		},
		"is_revoke": bson.M{
			"$ne": 1,
		},
	}

	if params.Receiver.TalkType == consts.ChatPrivateMode {
		filter["$or"] = bson.A{
			bson.M{"user_id": uid, "receiver_id": params.Receiver.ReceiverId},
			bson.M{"user_id": params.Receiver.ReceiverId, "receiver_id": uid},
		}
	}

	count, err := dao.TalkRecords.CountDocuments(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if int(count) != len(params.MessageIds) {
		logger.Error(ctx, err)
		return errors.New("转发消息异常")
	}

	return nil
}

// 批量合并转发
func (s *sTalkMessage) MultiMergeForward(ctx context.Context, uid int, params *model.ForwardMessageReq) ([]*model.ForwardRecord, error) {

	receives := make([]map[string]int, 0)

	for _, userId := range params.Uids {
		receives = append(receives, map[string]int{"receiver_id": userId, "talk_type": 1})
	}

	for _, gid := range params.Gids {
		receives = append(receives, map[string]int{"receiver_id": gid, "talk_type": 2})
	}

	tmpRecords, err := aggregation(ctx, params)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	ids := make([]int, 0)
	for _, id := range params.MessageIds {
		ids = append(ids, id)
	}

	forwardItems := make([]*model.ForwardItem, 0)
	for _, record := range tmpRecords {
		forwardItems = append(forwardItems, &model.ForwardItem{
			Nickname: record["nickname"],
			Text:     record["text"],
		})
	}

	extra := gjson.MustEncodeString(model.TalkRecordForward{
		RecordsIds: ids,
		Records:    tmpRecords,
		TalkType:   params.Receiver.TalkType,
	})

	records := make([]interface{}, 0, len(receives))
	recordIds := make([]int, 0)
	for _, item := range receives {
		data := &do.TalkRecords{
			RecordId:   core.IncrRecordId(ctx),
			MsgId:      util.NewMsgId(),
			TalkType:   item["talk_type"],
			MsgType:    consts.ChatMsgTypeForward,
			UserId:     uid,
			ReceiverId: item["receiver_id"],
			Extra:      extra,
			Forward: &model.Forward{
				TalkType:   item["talk_type"],
				RecordsIds: ids,
				Items:      forwardItems,
			},
			Sender: &model.Sender{
				Id: uid,
			},
			Receiver: &model.Receiver{
				Id:         item["receiver_id"],
				ReceiverId: item["receiver_id"],
			},
		}

		if data.TalkType == consts.ChatGroupMode {
			data.Sequence = dao.Sequence.Get(ctx, 0, data.ReceiverId)
		} else {
			data.Sequence = dao.Sequence.Get(ctx, uid, data.ReceiverId)
		}

		records = append(records, data)
		recordIds = append(recordIds, data.RecordId)
	}

	_, err = dao.TalkRecords.Inserts(ctx, records)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	list := make([]*model.ForwardRecord, 0, len(records))
	for i, record := range records {
		r := record.(*do.TalkRecords)
		list = append(list, &model.ForwardRecord{
			RecordId:   recordIds[i],
			ReceiverId: r.ReceiverId,
			TalkType:   r.TalkType,
		})
	}

	return list, nil
}

// 批量逐条转发
func (s *sTalkMessage) MultiSplitForward(ctx context.Context, uid int, params *model.ForwardMessageReq) ([]*model.ForwardRecord, error) {

	receives := make([]map[string]int, 0)

	for _, userId := range params.Uids {
		receives = append(receives, map[string]int{"receiver_id": userId, "talk_type": consts.TalkRecordTalkTypePrivate})
	}

	for _, gid := range params.Gids {
		receives = append(receives, map[string]int{"receiver_id": gid, "talk_type": consts.TalkRecordTalkTypeGroup})
	}

	records, err := dao.TalkRecords.Find(ctx, bson.M{"record_id": bson.M{"$in": params.MessageIds}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]interface{}, 0, len(receives)*len(records))

	recordsLen := int64(len(records))
	recordIds := make([]int, 0)
	for _, v := range receives {
		var sequences []int64

		if v["talk_type"] == consts.TalkRecordTalkTypeGroup {
			sequences = dao.Sequence.BatchGet(ctx, 0, v["receiver_id"], recordsLen)
		} else {
			sequences = dao.Sequence.BatchGet(ctx, uid, v["receiver_id"], recordsLen)
		}

		for i, item := range records {
			records := &do.TalkRecords{
				RecordId:   core.IncrRecordId(ctx),
				MsgId:      util.NewMsgId(),
				TalkType:   v["talk_type"],
				MsgType:    item.MsgType,
				UserId:     uid,
				ReceiverId: v["receiver_id"],
				Content:    item.Content,
				Sequence:   sequences[i],
				Extra:      item.Extra,
				Text:       item.Text,
				Code:       item.Code,
				Image:      item.Image,
				Voice:      item.Voice,
				Video:      item.Video,
				File:       item.File,
				Vote:       item.Vote,
				Mixed:      item.Mixed,
				Forward:    item.Forward,
				Emoticon:   item.Emoticon,
				Card:       item.Card,
				Location:   item.Location,
				Login:      item.Login,
				Sender:     item.Sender,
				Receiver:   item.Receiver,
			}
			items = append(items, records)
			recordIds = append(recordIds, records.RecordId)
		}
	}

	_, err = dao.TalkRecords.Inserts(ctx, items)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	list := make([]*model.ForwardRecord, 0, len(items))
	for i, item := range items {
		r := item.(*do.TalkRecords)
		list = append(list, &model.ForwardRecord{
			RecordId:   recordIds[i],
			ReceiverId: r.ReceiverId,
			TalkType:   r.TalkType,
		})
	}

	return list, nil
}

type forwardItem struct {
	MsgType  int    `json:"msg_type"`
	Content  string `json:"content"`
	Nickname string `json:"nickname"`
}

// 聚合转发数据
func aggregation(ctx context.Context, params *model.ForwardMessageReq) ([]map[string]string, error) {

	ids := params.MessageIds
	if len(ids) > 3 {
		ids = ids[:3]
	}

	talkRecordsList, err := dao.TalkRecords.Find(ctx, bson.M{"record_id": bson.M{"$in": ids}}, "created_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userIds := make([]int, 0)
	for _, talkRecords := range talkRecordsList {
		userIds = append(userIds, talkRecords.UserId)
	}

	userList, err := dao.User.FindUserListByUserIds(ctx, userIds)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userMap := util.ToMap(userList, func(t *entity.User) int {
		return t.UserId
	})

	rows := make([]*forwardItem, 0, 3)
	for _, talkRecords := range talkRecordsList {
		rows = append(rows, &forwardItem{
			MsgType:  talkRecords.MsgType,
			Content:  talkRecords.Content,
			Nickname: userMap[talkRecords.UserId].Nickname,
		})
	}

	data := make([]map[string]string, 0)
	for _, row := range rows {
		item := map[string]string{
			"nickname": row.Nickname,
		}

		switch row.MsgType {
		case consts.ChatMsgTypeText:
			item["text"] = gstr.SubStr(strings.TrimSpace(row.Content), 0, 30)
		case consts.ChatMsgTypeCode:
			item["text"] = "【代码消息】"
		case consts.ChatMsgTypeImage:
			item["text"] = "【图片消息】"
		case consts.ChatMsgTypeAudio:
			item["text"] = "【语音消息】"
		case consts.ChatMsgTypeVideo:
			item["text"] = "【视频消息】"
		case consts.ChatMsgTypeFile:
			item["text"] = "【文件消息】"
		case consts.ChatMsgTypeLocation:
			item["text"] = "【位置消息】"
		case consts.ChatMsgTypeForward:
			item["text"] = "【转发消息】"
		}

		data = append(data, item)
	}

	return data, nil
}

// 收藏表情包
func (s *sTalkMessage) Collect(ctx context.Context, params model.MessageCollectReq) error {

	uid := service.Session().GetUid(ctx)

	record, err := dao.TalkRecords.FindByRecordId(ctx, params.RecordId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if record.MsgType != consts.ChatMsgTypeImage {
		return errors.New("当前消息不支持收藏")
	}

	if record.IsRevoke == 1 {
		return errors.New("当前消息不支持收藏")
	}

	if record.TalkType == consts.ChatPrivateMode {
		if record.UserId != uid && record.ReceiverId != uid {
			return errors.ERR_PERMISSION_DENIED
		}
	} else if record.TalkType == consts.ChatGroupMode {
		if !dao.GroupMember.IsMember(ctx, record.ReceiverId, uid, true) {
			return errors.ERR_PERMISSION_DENIED
		}
	}

	var file model.TalkRecordImage
	if err = gjson.Unmarshal([]byte(record.Extra), &file); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.Insert(ctx, dao.Emoticon.Database, &do.EmoticonItem{
		UserId:     uid,
		Url:        file.Url,
		FileSuffix: file.Suffix,
		FileSize:   file.Size,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

//////////////// todo

// 发送消息接口
func (s *sTalkMessage) Publish(ctx context.Context, params model.MessagePublishReq) error {

	if err := s.VerifyPermission(ctx, &model.VerifyInfo{
		TalkType:          params.Receiver.TalkType,
		UserId:            service.Session().GetUid(ctx),
		ReceiverId:        params.Receiver.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return s.transfer(ctx, params.Type)
}

// 文本消息
func (s *sTalkMessage) onSendText(ctx context.Context) error {

	textMessageReq := &model.TextMessageReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), textMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = s.SendMessage(ctx, &model.Message{
		TalkType: textMessageReq.Receiver.TalkType,
		MsgType:  consts.MsgTypeText,
		QuoteId:  textMessageReq.QuoteId,
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			Id:         textMessageReq.Receiver.ReceiverId,
			ReceiverId: textMessageReq.Receiver.ReceiverId,
		},
		Text: &model.Text{
			Content: util.EscapeHtml(textMessageReq.Content),
		},
		Mention: textMessageReq.Mention,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	uid := service.Session().GetUid(ctx)

	// todo 检查是否需要机器人回复
	_ = grpool.AddWithRecover(gctx.New(), func(ctx context.Context) {

		senderId := textMessageReq.Receiver.ReceiverId
		receiverId := uid

		if textMessageReq.Receiver.TalkType == 2 {
			if len(textMessageReq.Mention.Uids) == 0 {
				return
			}
			senderId = textMessageReq.Mention.Uids[0]
			receiverId = textMessageReq.Receiver.ReceiverId
		}

		robotInfo, isNeed := robot.IsNeedRobotReply(ctx, senderId, textMessageReq.Mention.Uids)
		if isNeed {
			mentions := make([]string, 0)
			if textMessageReq.Receiver.TalkType == 2 {
				user, err := dao.User.FindUserByUserId(ctx, uid)
				if err != nil {
					logger.Error(ctx, err)
					return
				}
				mentions = append(mentions, user.Nickname)
			}

			session, err := service.TalkSession().FindBySession(ctx, uid, textMessageReq.Receiver.ReceiverId, textMessageReq.Receiver.TalkType)
			if err != nil {
				return
			}

			// 机器人回复
			robot.RobotReply(ctx, robotInfo, senderId, receiverId, textMessageReq.Receiver.TalkType, textMessageReq.Content, session.IsOpenContext, mentions...)
		}
	}, nil)

	return nil
}

// 图片消息
func (s *sTalkMessage) onSendImage(ctx context.Context) error {

	imageMessageReq := &model.ImageMessageReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), imageMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = s.SendMessage(ctx, &model.Message{
		TalkType: imageMessageReq.Receiver.TalkType,
		MsgType:  consts.MsgTypeImage,
		QuoteId:  imageMessageReq.QuoteId,
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			Id:         imageMessageReq.Receiver.ReceiverId,
			ReceiverId: imageMessageReq.Receiver.ReceiverId,
		},
		Image: &model.Image{
			Url:    imageMessageReq.Url,
			Width:  imageMessageReq.Width,
			Height: imageMessageReq.Height,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 语音消息
func (s *sTalkMessage) onSendVoice(ctx context.Context) error {

	voiceMessageReq := &model.VoiceMessageReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), voiceMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = s.SendMessage(ctx, &model.Message{
		TalkType: voiceMessageReq.Receiver.TalkType,
		MsgType:  consts.MsgTypeVoice,
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			Id:         voiceMessageReq.Receiver.ReceiverId,
			ReceiverId: voiceMessageReq.Receiver.ReceiverId,
		},
		Voice: &model.Voice{
			Url:      voiceMessageReq.Url,
			Duration: voiceMessageReq.Duration,
			Size:     voiceMessageReq.Size,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 视频消息
func (s *sTalkMessage) onSendVideo(ctx context.Context) error {

	videoMessageReq := &model.VideoMessageReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), videoMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = s.SendMessage(ctx, &model.Message{
		TalkType: videoMessageReq.Receiver.TalkType,
		MsgType:  consts.MsgTypeVideo,
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			Id:         videoMessageReq.Receiver.ReceiverId,
			ReceiverId: videoMessageReq.Receiver.ReceiverId,
		},
		Video: &model.Video{
			Cover:    videoMessageReq.Cover,
			Size:     videoMessageReq.Size,
			Url:      videoMessageReq.Url,
			Duration: videoMessageReq.Duration,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 文件消息
func (s *sTalkMessage) onSendFile(ctx context.Context, req ...*model.MessageFileReq) error {

	fileMessageReq := &model.MessageFileReq{}
	if len(req) > 0 {
		fileMessageReq = req[0]
	} else {
		err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), fileMessageReq)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	//err = s.SendFile(ctx, service.Session().GetUid(ctx), fileMessageReq)
	//if err != nil {
	//	logger.Error(ctx, err)
	//	return err
	//}
	uid := service.Session().GetUid(ctx)
	file, err := dao.SplitUpload.GetFile(ctx, uid, fileMessageReq.UploadId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	publicUrl := ""
	filePath := fmt.Sprintf("private/files/talks/%s/%s.%s", util.DateNumber(), gmd5.MustEncryptString(util.Random(16)), file.FileExt)

	// 公开文件
	if consts.GetMediaType(file.FileExt) <= 3 {
		filePath = fmt.Sprintf("public/media/%s/%s.%s", util.DateNumber(), gmd5.MustEncryptString(util.Random(16)), file.FileExt)
		publicUrl = filesystem.NewFilesystem(config.Cfg).Default.PublicUrl(filePath)
	}

	if err := filesystem.NewFilesystem(config.Cfg).Default.Copy(file.Path, filePath); err != nil {
		logger.Error(ctx, err)
		return err
	}

	message := &model.Message{
		TalkType: fileMessageReq.Receiver.TalkType,
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			Id:         fileMessageReq.Receiver.ReceiverId,
			ReceiverId: fileMessageReq.Receiver.ReceiverId,
		},
	}

	switch consts.GetMediaType(file.FileExt) {
	case consts.MediaFileAudio:
		message.MsgType = consts.MsgTypeVoice
		message.Voice = &model.Voice{
			Suffix:   file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
			Duration: 0,
		}
	case consts.MediaFileVideo:
		message.MsgType = consts.MsgTypeVideo
		message.Video = &model.Video{
			Cover:    "",
			Suffix:   file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
			Duration: 0,
		}
	case consts.MediaFileOther:
		message.MsgType = consts.MsgTypeFile
		message.File = &model.File{
			Drive:  file.Drive,
			Name:   file.OriginalName,
			Suffix: file.FileExt,
			Size:   int(file.FileSize),
			Path:   filePath,
		}
	}

	if err = s.SendMessage(ctx, message); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 代码消息
func (s *sTalkMessage) onSendCode(ctx context.Context) error {

	codeMessageReq := &model.CodeMessageReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), codeMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = s.SendMessage(ctx, &model.Message{
		TalkType: codeMessageReq.Receiver.TalkType,
		MsgType:  consts.MsgTypeCode,
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			Id:         codeMessageReq.Receiver.ReceiverId,
			ReceiverId: codeMessageReq.Receiver.ReceiverId,
		},
		Code: &model.Code{
			Lang: codeMessageReq.Lang,
			Code: codeMessageReq.Code,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 位置消息
func (s *sTalkMessage) onSendLocation(ctx context.Context) error {

	locationMessageReq := &model.LocationMessageReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), locationMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = s.SendMessage(ctx, &model.Message{
		TalkType: locationMessageReq.Receiver.TalkType,
		MsgType:  consts.MsgTypeLocation,
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			Id:         locationMessageReq.Receiver.ReceiverId,
			ReceiverId: locationMessageReq.Receiver.ReceiverId,
		},
		Location: &model.Location{
			Longitude:   locationMessageReq.Longitude,
			Latitude:    locationMessageReq.Latitude,
			Description: locationMessageReq.Description,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 转发消息
func (s *sTalkMessage) onSendForward(ctx context.Context) error {

	forwardMessageReq := &model.ForwardMessageReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), forwardMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, uid := range forwardMessageReq.Uids {
		if err := s.VerifyPermission(ctx, &model.VerifyInfo{
			TalkType:          consts.ChatPrivateMode,
			UserId:            service.Session().GetUid(ctx),
			ReceiverId:        uid,
			IsVerifyGroupMute: true,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	for _, gid := range forwardMessageReq.Gids {
		if err := s.VerifyPermission(ctx, &model.VerifyInfo{
			TalkType:          consts.ChatGroupMode,
			UserId:            service.Session().GetUid(ctx),
			ReceiverId:        gid,
			IsVerifyGroupMute: true,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if err = s.SendMessage(ctx, &model.Message{
		TalkType: forwardMessageReq.Receiver.TalkType,
		MsgType:  consts.MsgTypeForward,
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			Id:         forwardMessageReq.Receiver.ReceiverId,
			ReceiverId: forwardMessageReq.Receiver.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 表情消息
func (s *sTalkMessage) onSendEmoticon(ctx context.Context) error {

	emoticonMessageReq := &model.EmoticonMessageReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), emoticonMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	emoticon := new(entity.EmoticonItem)
	if err := dao.FindOne(ctx, dao.Emoticon.Database, do.EMOTICON_ITEM_COLLECTION, bson.M{"_id": emoticonMessageReq.EmoticonId, "user_id": service.Session().GetUid(ctx)}, &emoticon); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("表情信息不存在")
		}

		logger.Error(ctx, err)
		return err
	}

	if err = s.SendMessage(ctx, &model.Message{
		TalkType: emoticonMessageReq.Receiver.TalkType,
		MsgType:  consts.MsgTypeEmoticon,
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			Id:         emoticonMessageReq.Receiver.ReceiverId,
			ReceiverId: emoticonMessageReq.Receiver.ReceiverId,
		},
		Emoticon: &model.Emoticon{
			Url:    emoticon.Url,
			Width:  0,
			Height: 0,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 投票消息
func (s *sTalkMessage) onSendVote(ctx context.Context) error {

	voteMessageReq := &model.MessageVoteReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), voteMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if len(voteMessageReq.Options) <= 1 {
		return errors.New("options 选项必须大于1")
	}

	if len(voteMessageReq.Options) > 6 {
		return errors.New("options 选项不能超过6个")
	}

	err = s.SendVote(ctx, service.Session().GetUid(ctx), voteMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 名片消息
func (s *sTalkMessage) onSendCard(ctx context.Context) error {

	cardMessageReq := &model.CardMessageReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), cardMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = s.SendMessage(ctx, &model.Message{
		TalkType: cardMessageReq.Receiver.TalkType,
		MsgType:  consts.MsgTypeCard,
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			Id:         cardMessageReq.Receiver.ReceiverId,
			ReceiverId: cardMessageReq.Receiver.ReceiverId,
		},
		Card: &model.Card{
			UserId: cardMessageReq.UserId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 图文消息
func (s *sTalkMessage) onMixedMessage(ctx context.Context) error {

	mixedMessageReq := &model.MixedMessageReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), mixedMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	items := make([]*model.MixedItem, 0)

	for _, item := range mixedMessageReq.Items {
		if item.Type == 1 {
			items = append(items, &model.MixedItem{
				MsgType: consts.MsgTypeText,
				Text: &model.Text{
					Content: item.Content,
				},
			})
		} else if item.Type == 3 {
			items = append(items, &model.MixedItem{
				MsgType: consts.MsgTypeImage,
				Image: &model.Image{
					Url: item.Content,
				},
			})
		}
	}

	if err = s.SendMessage(ctx, &model.Message{
		TalkType: mixedMessageReq.Receiver.TalkType,
		MsgType:  consts.MsgTypeMixed,
		QuoteId:  mixedMessageReq.QuoteId,
		Sender: &model.Sender{
			Id: service.Session().GetUid(ctx),
		},
		Receiver: &model.Receiver{
			Id:         mixedMessageReq.Receiver.ReceiverId,
			ReceiverId: mixedMessageReq.Receiver.ReceiverId,
		},
		Mixed: &model.Mixed{Items: items},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (s *sTalkMessage) transfer(ctx context.Context, typeValue string) error {

	if s.mapping == nil {
		s.mapping = make(map[string]func(ctx context.Context) error)
		s.mapping["text"] = s.onSendText
		s.mapping["code"] = s.onSendCode
		s.mapping["location"] = s.onSendLocation
		s.mapping["emoticon"] = s.onSendEmoticon
		s.mapping["vote"] = s.onSendVote
		s.mapping["image"] = s.onSendImage
		s.mapping["voice"] = s.onSendVoice
		s.mapping["video"] = s.onSendVideo
		//s.mapping["file"] = s.onSendFile todo
		s.mapping["card"] = s.onSendCard
		s.mapping["forward"] = s.onSendForward
		s.mapping["mixed"] = s.onMixedMessage
	}

	if call, ok := s.mapping[typeValue]; ok {
		return call(ctx)
	}

	return nil
}

func (s *sTalkMessage) TextMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	data := &model.TalkRecord{
		TalkType:   message.TalkType,
		MsgType:    consts.ChatMsgTypeText,
		QuoteId:    message.QuoteId,
		UserId:     message.Sender.Id,
		ReceiverId: message.Receiver.ReceiverId,
		Content:    util.EscapeHtml(message.Text.Content),
		Text:       message.Text,
		Sender:     message.Sender,
		Receiver:   message.Receiver,
	}

	if message.Mention != nil && len(message.Mention.Uids) > 0 {
		data.Mention = message.Mention
	}

	return data, nil
}

func (s *sTalkMessage) CodeMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	data := &model.TalkRecord{
		TalkType:   message.TalkType,
		MsgType:    consts.ChatMsgTypeCode,
		QuoteId:    message.QuoteId,
		UserId:     message.Sender.Id,
		ReceiverId: message.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordCode{
			Lang: message.Code.Lang,
			Code: message.Code.Code,
		}),
		Code:     message.Code,
		Sender:   message.Sender,
		Receiver: message.Receiver,
	}

	return data, nil
}

func (s *sTalkMessage) ImageMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	data := &model.TalkRecord{
		TalkType:   message.TalkType,
		MsgType:    consts.ChatMsgTypeImage,
		QuoteId:    message.QuoteId,
		UserId:     message.Sender.Id,
		ReceiverId: message.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordImage{
			Url:    message.Image.Url,
			Width:  message.Image.Width,
			Height: message.Image.Height,
		}),
		Image:    message.Image,
		Sender:   message.Sender,
		Receiver: message.Receiver,
	}

	return data, nil
}

func (s *sTalkMessage) VoiceMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	data := &model.TalkRecord{
		TalkType:   message.TalkType,
		MsgType:    consts.ChatMsgTypeAudio,
		QuoteId:    message.QuoteId,
		UserId:     message.Sender.Id,
		ReceiverId: message.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordAudio{
			Suffix:   gfile.ExtName(message.Voice.Url),
			Size:     message.Voice.Size,
			Url:      message.Voice.Url,
			Duration: 0,
		}),
		Voice:    message.Voice,
		Sender:   message.Sender,
		Receiver: message.Receiver,
	}

	return data, nil
}

func (s *sTalkMessage) VideoMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	data := &model.TalkRecord{
		TalkType:   message.TalkType,
		MsgType:    consts.ChatMsgTypeVideo,
		QuoteId:    message.QuoteId,
		UserId:     message.Sender.Id,
		ReceiverId: message.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordVideo{
			Cover:    message.Video.Cover,
			Size:     message.Video.Size,
			Url:      message.Video.Url,
			Duration: message.Video.Duration,
		}),
		Video:    message.Video,
		Sender:   message.Sender,
		Receiver: message.Receiver,
	}

	return data, nil
}

func (s *sTalkMessage) FileMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	data := &model.TalkRecord{
		TalkType:   message.TalkType,
		MsgType:    consts.ChatMsgTypeFile,
		QuoteId:    message.QuoteId,
		UserId:     message.Sender.Id,
		ReceiverId: message.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordFile{
			Drive:  message.File.Drive,
			Name:   message.File.Name,
			Suffix: message.File.Suffix,
			Size:   message.File.Size,
			Path:   message.File.Path,
		}),
		File:     message.File,
		Sender:   message.Sender,
		Receiver: message.Receiver,
	}

	return data, nil
}

func (s *sTalkMessage) VoteMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	data := &model.TalkRecord{
		TalkType:   consts.ChatGroupMode,
		MsgType:    consts.ChatMsgTypeVote,
		QuoteId:    message.QuoteId,
		RecordId:   core.IncrRecordId(ctx),
		MsgId:      util.NewMsgId(),
		UserId:     message.Sender.Id,
		ReceiverId: message.Receiver.ReceiverId,
		Vote:       message.Vote,
		Sender:     message.Sender,
		Receiver:   message.Receiver,
	}

	return data, nil
}

func (s *sTalkMessage) MixedMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	items := make([]*model.MixedMessage, 0)

	for _, item := range message.Mixed.Items {
		if item.MsgType == consts.MsgTypeText {
			items = append(items, &model.MixedMessage{
				Type:    1,
				Content: item.Text.Content,
			})
		} else if item.MsgType == consts.MsgTypeImage {
			items = append(items, &model.MixedMessage{
				Type:    3,
				Content: item.Image.Url,
			})
		}
	}

	data := &model.TalkRecord{
		TalkType:   message.TalkType,
		MsgType:    consts.ChatMsgTypeMixed,
		QuoteId:    message.QuoteId,
		UserId:     message.Sender.Id,
		ReceiverId: message.Receiver.ReceiverId,
		Extra:      gjson.MustEncodeString(model.TalkRecordMixed{Items: items}),
		Mixed:      message.Mixed,
		Sender:     message.Sender,
		Receiver:   message.Receiver,
	}

	return data, nil
}

// todo
func (s *sTalkMessage) ForwardMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	req := &model.ForwardMessageReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), req)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	uid := message.Sender.Id

	// 验证转发消息合法性
	if err = s.Verify(ctx, uid, req); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	var items []*model.ForwardRecord
	// 发送方式 1:逐条发送 2:合并发送
	if req.Mode == 1 {
		items, err = s.MultiSplitForward(ctx, uid, req)
	} else {
		items, err = s.MultiMergeForward(ctx, uid, req)
	}

	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, record := range items {
		if record.TalkType == consts.ChatPrivateMode {
			s.unreadStorage.Incr(ctx, consts.ChatPrivateMode, uid, record.ReceiverId)
		} else if record.TalkType == consts.ChatGroupMode {
			pipe := redis.Pipeline(ctx)
			for _, uid := range dao.GroupMember.GetMemberIds(ctx, record.ReceiverId) {
				s.unreadStorage.PipeIncr(ctx, pipe, consts.ChatGroupMode, record.ReceiverId, uid)
			}
			if _, err := pipe.Exec(ctx); err != nil {
				logger.Error(ctx, err)
			}
		}

		_ = s.messageStorage.Set(ctx, record.TalkType, uid, record.ReceiverId, &cache.LastCacheMessage{
			Content:  "[转发消息]",
			Datetime: gtime.Datetime(),
		})
	}

	pipe := redis.Pipeline(ctx)

	for _, item := range items {
		data := gjson.MustEncodeString(map[string]any{
			"event": consts.SubEventImMessage,
			"data": gjson.MustEncodeString(map[string]any{
				"sender_id":   uid,
				"receiver_id": item.ReceiverId,
				"talk_type":   item.TalkType,
				"record_id":   item.RecordId,
			}),
		})

		pipe.Publish(ctx, consts.ImTopicChat, data)
	}

	_, _ = redis.Pipelined(ctx, pipe)

	return nil, nil
}

func (s *sTalkMessage) EmoticonMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	data := &model.TalkRecord{
		TalkType:   message.TalkType,
		MsgType:    consts.ChatMsgTypeImage,
		QuoteId:    message.QuoteId,
		UserId:     message.Sender.Id,
		ReceiverId: message.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordImage{
			Url:    message.Emoticon.Url,
			Width:  0,
			Height: 0,
		}),
		Emoticon: message.Emoticon,
		Sender:   message.Sender,
		Receiver: message.Receiver,
	}

	return data, nil
}

func (s *sTalkMessage) CardMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	data := &model.TalkRecord{
		TalkType:   message.TalkType,
		MsgType:    consts.ChatMsgTypeCard,
		QuoteId:    message.QuoteId,
		UserId:     message.Sender.Id,
		ReceiverId: message.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordCard{
			UserId: message.Card.UserId,
		}),
		Card:     message.Card,
		Sender:   message.Sender,
		Receiver: message.Receiver,
	}

	return data, nil
}

func (s *sTalkMessage) LocationMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error) {

	data := &model.TalkRecord{
		TalkType:   message.TalkType,
		MsgType:    consts.ChatMsgTypeLocation,
		QuoteId:    message.QuoteId,
		UserId:     message.Sender.Id,
		ReceiverId: message.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordLocation{
			Longitude:   message.Location.Longitude,
			Latitude:    message.Location.Latitude,
			Description: message.Location.Description,
		}),
		Sender:   message.Sender,
		Receiver: message.Receiver,
	}

	return data, nil
}
