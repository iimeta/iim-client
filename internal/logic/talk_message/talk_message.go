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

var mapping map[string]func(ctx context.Context) error

// 系统文本消息
func (s *sTalkMessage) SendSystemText(ctx context.Context, uid int, req *model.TextMessageReq) error {

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgSysText,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Content:    html.EscapeString(req.Content),
	}

	return s.save(ctx, data)
}

// 文本消息
func (s *sTalkMessage) SendText(ctx context.Context, uid int, req *model.TextMessageReq) error {

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeText,
		QuoteId:    req.QuoteId,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Content:    util.EscapeHtml(req.Content),
	}

	return s.save(ctx, data)
}

// 图片文件消息
func (s *sTalkMessage) SendImage(ctx context.Context, uid int, req *model.ImageMessageReq) error {

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeImage,
		QuoteId:    req.QuoteId,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordImage{
			Url:    req.Url,
			Width:  req.Width,
			Height: req.Height,
		}),
	}

	return s.save(ctx, data)
}

// 语音文件消息
func (s *sTalkMessage) SendVoice(ctx context.Context, uid int, req *model.VoiceMessageReq) error {

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeAudio,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordAudio{
			Suffix:   gfile.ExtName(req.Url),
			Size:     req.Size,
			Url:      req.Url,
			Duration: 0,
		}),
	}

	return s.save(ctx, data)
}

// 视频文件消息
func (s *sTalkMessage) SendVideo(ctx context.Context, uid int, req *model.VideoMessageReq) error {

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeVideo,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordVideo{
			Cover:    req.Cover,
			Size:     req.Size,
			Url:      req.Url,
			Duration: req.Duration,
		}),
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
		MsgId:      gmd5.MustEncryptString(req.UploadId),
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

// 代码消息
func (s *sTalkMessage) SendCode(ctx context.Context, uid int, req *model.CodeMessageReq) error {

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeCode,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordCode{
			Lang: req.Lang,
			Code: req.Code,
		}),
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

	options := make(map[string]string)
	for i, value := range req.Options {
		options[fmt.Sprintf("%c", 65+i)] = value
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

// 表情消息
func (s *sTalkMessage) SendEmoticon(ctx context.Context, uid int, req *model.EmoticonMessageReq) error {

	emoticon := new(entity.EmoticonItem)
	if err := dao.FindOne(ctx, dao.Emoticon.Database, do.EMOTICON_ITEM_COLLECTION, bson.M{"_id": req.EmoticonId, "user_id": uid}, &emoticon); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("表情信息不存在")
		}

		logger.Error(ctx, err)
		return err
	}

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeImage,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordImage{
			Url:    emoticon.Url,
			Width:  0,
			Height: 0,
		}),
	}

	return s.save(ctx, data)
}

// 转发消息
func (s *sTalkMessage) SendForward(ctx context.Context, uid int, req *model.ForwardMessageReq) error {

	// 验证转发消息合法性
	if err := s.Verify(ctx, uid, req); err != nil {
		logger.Error(ctx, err)
		return err
	}

	var (
		err   error
		items []*model.ForwardRecord
	)

	// 发送方式 1:逐条发送 2:合并发送
	if req.Mode == 1 {
		items, err = s.MultiSplitForward(ctx, uid, req)
	} else {
		items, err = s.MultiMergeForward(ctx, uid, req)
	}

	if err != nil {
		logger.Error(ctx, err)
		return err
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

	return nil
}

// 位置消息
func (s *sTalkMessage) SendLocation(ctx context.Context, uid int, req *model.LocationMessageReq) error {

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeLocation,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordLocation{
			Longitude:   req.Longitude,
			Latitude:    req.Latitude,
			Description: req.Description,
		}),
	}

	return s.save(ctx, data)
}

// 推送用户名片消息
func (s *sTalkMessage) SendBusinessCard(ctx context.Context, uid int, req *model.CardMessageReq) error {

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeCard,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordCard{
			UserId: req.UserId,
		}),
	}

	return s.save(ctx, data)
}

// 推送用户登录消息
func (s *sTalkMessage) SendLogin(ctx context.Context, uid int, req *model.LoginMessageReq) error {

	robot, err := dao.Robot.GetLoginRobot(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	data := &model.TalkRecord{
		TalkType:   consts.ChatPrivateMode,
		MsgType:    consts.ChatMsgTypeLogin,
		UserId:     robot.UserId,
		ReceiverId: uid,
		Extra: gjson.MustEncodeString(&model.TalkRecordLogin{
			IP:       req.Ip,
			Platform: req.Platform,
			Agent:    req.Agent,
			Address:  req.Address,
			Reason:   req.Reason,
			Datetime: gtime.Datetime(),
		}),
	}

	return s.save(ctx, data)
}

// 图文消息
func (s *sTalkMessage) SendMixedMessage(ctx context.Context, uid int, req *model.MixedMessageReq) error {

	items := make([]*model.TalkRecordMixedItem, 0)

	for _, item := range req.Items {
		items = append(items, &model.TalkRecordMixedItem{
			Type:    item.Type,
			Content: item.Content,
		})
	}

	data := &model.TalkRecord{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeMixed,
		QuoteId:    req.QuoteId,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra:      gjson.MustEncodeString(model.TalkRecordMixed{Items: items}),
	}

	return s.save(ctx, data)
}

// 推送其它消息
func (s *sTalkMessage) SendSysOther(ctx context.Context, data *model.TalkRecord) error {
	return s.save(ctx, data)
}

// 撤回消息
func (s *sTalkMessage) Revoke(ctx context.Context, params model.MessageRevokeReq) error {

	uid := service.Session().GetUid(ctx)

	record, err := dao.TalkRecords.FindById(ctx, params.RecordId)
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

	if err := dao.TalkRecords.UpdateById(ctx, params.RecordId, bson.M{"is_revoke": 1}); err != nil {
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

// 发送文本消息
func (s *sTalkMessage) Text(ctx context.Context, params model.TextMessageReq) error {

	uid := service.Session().GetUid(ctx)

	// todo auth
	if err := service.Group().GroupAuth(ctx, &model.GroupAuth{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendText(ctx, uid, &model.TextMessageReq{
		Content: params.Text,
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

// 发送代码块消息
func (s *sTalkMessage) Code(ctx context.Context, params model.CodeMessageReq) error {

	uid := service.Session().GetUid(ctx)
	if err := service.Group().GroupAuth(ctx, &model.GroupAuth{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendCode(ctx, uid, &model.CodeMessageReq{
		Lang: params.Lang,
		Code: params.Code,
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

	// 判断上传文件大小(5M)
	if file.Size > 5<<20 {
		return errors.New("上传文件大小不能超过5M")
	}

	if err = service.Group().GroupAuth(ctx, &model.GroupAuth{
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

	if err = service.TalkMessage().SendImage(ctx, service.Session().GetUid(ctx), &model.ImageMessageReq{
		Url:    s.Filesystem.Default.PublicUrl(filePath),
		Width:  meta.Width,
		Height: meta.Height,
		Size:   int(file.Size),
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
	if err := service.Group().GroupAuth(ctx, &model.GroupAuth{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendFile(ctx, uid, &model.MessageFileReq{
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
	if err := service.Group().GroupAuth(ctx, &model.GroupAuth{
		TalkType:          consts.ChatGroupMode,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendVote(ctx, uid, &model.MessageVoteReq{
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

// 发送表情包消息
func (s *sTalkMessage) Emoticon(ctx context.Context, params model.EmoticonMessageReq) error {

	uid := service.Session().GetUid(ctx)
	if err := service.Group().GroupAuth(ctx, &model.GroupAuth{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendEmoticon(ctx, uid, &model.EmoticonMessageReq{
		EmoticonId: params.EmoticonId,
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

// 发送转发消息
func (s *sTalkMessage) Forward(ctx context.Context, params model.ForwardMessageReq) error {

	if params.ReceiveGroupIds == "" && params.ReceiveUserIds == "" {
		return errors.New("receive_user_ids 和 receive_group_ids 不能都为空")
	}

	uid := service.Session().GetUid(ctx)
	if err := service.Group().GroupAuth(ctx, &model.GroupAuth{
		TalkType:   params.TalkType,
		UserId:     service.Session().GetUid(ctx),
		ReceiverId: params.ReceiverId,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	data := &model.ForwardMessageReq{
		Mode:       params.ForwardMode,
		MessageIds: make([]int, 0),
		Gids:       make([]int, 0),
		Uids:       make([]int, 0),
		Receiver: &model.Receiver{
			TalkType:   params.TalkType,
			ReceiverId: params.ReceiverId,
		},
	}

	for _, id := range util.ParseIds(params.RecordsIds) {
		data.MessageIds = append(data.MessageIds, id)
	}

	for _, id := range util.ParseIds(params.ReceiveUserIds) {
		data.Uids = append(data.Uids, id)
	}

	for _, id := range util.ParseIds(params.ReceiveGroupIds) {
		data.Gids = append(data.Gids, id)
	}

	if err := service.TalkMessage().SendForward(ctx, uid, data); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 发送用户名片消息
func (s *sTalkMessage) Card(ctx context.Context, params model.CardMessageReq) error {

	uid := service.Session().GetUid(ctx)
	if err := service.Group().GroupAuth(ctx, &model.GroupAuth{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendBusinessCard(ctx, uid, &model.CardMessageReq{
		UserId: params.ReceiverId,
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

// 发送位置消息
func (s *sTalkMessage) Location(ctx context.Context, params model.LocationMessageReq) error {

	uid := service.Session().GetUid(ctx)
	if err := service.Group().GroupAuth(ctx, &model.GroupAuth{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendLocation(ctx, uid, &model.LocationMessageReq{
		Longitude:   params.Longitude,
		Latitude:    params.Latitude,
		Description: "", // todo 需完善
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

// 验证转发消息合法性
func (s *sTalkMessage) Verify(ctx context.Context, uid int, params *model.ForwardMessageReq) error {

	filter := bson.M{
		"record_id": bson.M{
			"$in": params.MessageIds,
		},
		"talk_type": params.Receiver.TalkType,
		"msg_type": bson.M{
			"$in": []int{1, 2, 3, 4, 5, 6, 7, 8, consts.ChatMsgTypeForward},
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

	extra := gjson.MustEncodeString(model.TalkRecordForward{
		MsgIds:  ids,
		Records: tmpRecords,
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
func aggregation(ctx context.Context, params *model.ForwardMessageReq) ([]map[string]any, error) {

	ids := params.MessageIds
	if len(ids) > 3 {
		ids = ids[:3]
	}

	talkRecordsList, err := dao.TalkRecords.Find(ctx, bson.M{"record_id": bson.M{"$in": ids}})
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

	data := make([]map[string]any, 0)
	for _, row := range rows {
		item := map[string]any{
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
		}

		data = append(data, item)
	}

	return data, nil
}

// 收藏表情包
func (s *sTalkMessage) Collect(ctx context.Context, params model.MessageCollectReq) error {

	uid := service.Session().GetUid(ctx)

	record, err := dao.TalkRecords.FindById(ctx, params.RecordId)
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
			return consts.ErrPermissionDenied
		}
	} else if record.TalkType == consts.ChatGroupMode {
		if !dao.GroupMember.IsMember(ctx, record.ReceiverId, uid, true) {
			return consts.ErrPermissionDenied
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

////////////////

// 发送消息接口
func (s *sTalkMessage) Publish(ctx context.Context, params model.MessagePublishReq) error {

	// todo
	if err := service.Group().GroupAuth(ctx, &model.GroupAuth{
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

	err = service.TalkMessage().SendText(ctx, service.Session().GetUid(ctx), textMessageReq)
	if err != nil {
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

	err = service.TalkMessage().SendImage(ctx, service.Session().GetUid(ctx), imageMessageReq)
	if err != nil {
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

	err = service.TalkMessage().SendVoice(ctx, service.Session().GetUid(ctx), voiceMessageReq)
	if err != nil {
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

	err = service.TalkMessage().SendVideo(ctx, service.Session().GetUid(ctx), videoMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 文件消息
func (s *sTalkMessage) onSendFile(ctx context.Context) error {

	fileMessageReq := &model.MessageFileReq{}
	err := gjson.Unmarshal(g.RequestFromCtx(ctx).GetBody(), fileMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	err = service.TalkMessage().SendFile(ctx, service.Session().GetUid(ctx), fileMessageReq)
	if err != nil {
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

	err = service.TalkMessage().SendCode(ctx, service.Session().GetUid(ctx), codeMessageReq)
	if err != nil {
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

	err = service.TalkMessage().SendLocation(ctx, service.Session().GetUid(ctx), locationMessageReq)
	if err != nil {
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

	err = service.TalkMessage().SendForward(ctx, service.Session().GetUid(ctx), forwardMessageReq)
	if err != nil {
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

	err = service.TalkMessage().SendEmoticon(ctx, service.Session().GetUid(ctx), emoticonMessageReq)
	if err != nil {
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

	err = service.TalkMessage().SendVote(ctx, service.Session().GetUid(ctx), voteMessageReq)
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

	err = service.TalkMessage().SendBusinessCard(ctx, service.Session().GetUid(ctx), cardMessageReq)
	if err != nil {
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

	err = service.TalkMessage().SendMixedMessage(ctx, service.Session().GetUid(ctx), mixedMessageReq)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (s *sTalkMessage) transfer(ctx context.Context, typeValue string) error {

	if mapping == nil {
		mapping = make(map[string]func(ctx context.Context) error)
		mapping["text"] = s.onSendText
		mapping["code"] = s.onSendCode
		mapping["location"] = s.onSendLocation
		mapping["emoticon"] = s.onSendEmoticon
		mapping["vote"] = s.onSendVote
		mapping["image"] = s.onSendImage
		mapping["voice"] = s.onSendVoice
		mapping["video"] = s.onSendVideo
		mapping["file"] = s.onSendFile
		mapping["card"] = s.onSendCard
		mapping["forward"] = s.onSendForward
		mapping["mixed"] = s.onMixedMessage
	}

	if call, ok := mapping[typeValue]; ok {
		return call(ctx)
	}

	return nil
}
