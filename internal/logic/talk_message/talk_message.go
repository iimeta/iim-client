package talk_message

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
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
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"html"
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
	}
}

// 系统文本消息
func (s *sTalkMessage) SendSystemText(ctx context.Context, uid int, req *model.TextMessageReq) error {

	data := &model.TalkRecords{
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

	data := &model.TalkRecords{
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

	data := &model.TalkRecords{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeImage,
		QuoteId:    req.QuoteId,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordExtraImage{
			Url:    req.Url,
			Width:  req.Width,
			Height: req.Height,
		}),
	}

	return s.save(ctx, data)
}

// 语音文件消息
func (s *sTalkMessage) SendVoice(ctx context.Context, uid int, req *model.VoiceMessageReq) error {

	data := &model.TalkRecords{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeAudio,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordExtraAudio{
			Suffix:   util.FileSuffix(req.Url),
			Size:     req.Size,
			Url:      req.Url,
			Duration: 0,
		}),
	}

	return s.save(ctx, data)
}

// 视频文件消息
func (s *sTalkMessage) SendVideo(ctx context.Context, uid int, req *model.VideoMessageReq) error {

	data := &model.TalkRecords{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeVideo,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordExtraVideo{
			Cover:    req.Cover,
			Size:     req.Size,
			Url:      req.Url,
			Duration: req.Duration,
		}),
	}

	return s.save(ctx, data)
}

// 文件消息
func (s *sTalkMessage) SendFile(ctx context.Context, uid int, req *model.FileMessageReq) error {

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

	data := &model.TalkRecords{
		MsgId:      gmd5.MustEncryptString(req.UploadId),
		TalkType:   req.Receiver.TalkType,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
	}

	switch consts.GetMediaType(file.FileExt) {
	case consts.MediaFileAudio:
		data.MsgType = consts.ChatMsgTypeAudio
		data.Extra = gjson.MustEncodeString(&model.TalkRecordExtraAudio{
			Suffix:   file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
			Duration: 0,
		})
	case consts.MediaFileVideo:
		data.MsgType = consts.ChatMsgTypeVideo
		data.Extra = gjson.MustEncodeString(&model.TalkRecordExtraVideo{
			Cover:    "",
			Suffix:   file.FileExt,
			Size:     int(file.FileSize),
			Url:      publicUrl,
			Duration: 0,
		})
	case consts.MediaFileOther:
		data.MsgType = consts.ChatMsgTypeFile
		data.Extra = gjson.MustEncodeString(&model.TalkRecordExtraFile{
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

	data := &model.TalkRecords{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeCode,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordExtraCode{
			Lang: req.Lang,
			Code: req.Code,
		}),
	}

	return s.save(ctx, data)
}

// 投票消息
func (s *sTalkMessage) SendVote(ctx context.Context, uid int, req *model.VoteMessageReq) error {

	data := &model.TalkRecords{
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

	data := &model.TalkRecords{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeImage,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordExtraImage{
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
	if err := service.MessageForward().Verify(ctx, uid, req); err != nil {
		logger.Error(ctx, err)
		return err
	}

	var (
		err   error
		items []*model.ForwardRecord
	)

	// 发送方式 1:逐条发送 2:合并发送
	if req.Mode == 1 {
		items, err = service.MessageForward().MultiSplitForward(ctx, uid, req)
	} else {
		items, err = service.MessageForward().MultiMergeForward(ctx, uid, req)
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

	data := &model.TalkRecords{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeLocation,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordExtraLocation{
			Longitude:   req.Longitude,
			Latitude:    req.Latitude,
			Description: req.Description,
		}),
	}

	return s.save(ctx, data)
}

// 推送用户名片消息
func (s *sTalkMessage) SendBusinessCard(ctx context.Context, uid int, req *model.CardMessageReq) error {

	data := &model.TalkRecords{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeCard,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra: gjson.MustEncodeString(&model.TalkRecordExtraCard{
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

	data := &model.TalkRecords{
		TalkType:   consts.ChatPrivateMode,
		MsgType:    consts.ChatMsgTypeLogin,
		UserId:     robot.UserId,
		ReceiverId: uid,
		Extra: gjson.MustEncodeString(&model.TalkRecordExtraLogin{
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

	items := make([]*model.TalkRecordExtraMixedItem, 0)

	for _, item := range req.Items {
		items = append(items, &model.TalkRecordExtraMixedItem{
			Type:    item.Type,
			Content: item.Content,
		})
	}

	data := &model.TalkRecords{
		TalkType:   req.Receiver.TalkType,
		MsgType:    consts.ChatMsgTypeMixed,
		QuoteId:    req.QuoteId,
		UserId:     uid,
		ReceiverId: req.Receiver.ReceiverId,
		Extra:      gjson.MustEncodeString(model.TalkRecordExtraMixed{Items: items}),
	}

	return s.save(ctx, data)
}

// 推送其它消息
func (s *sTalkMessage) SendSysOther(ctx context.Context, data *model.TalkRecords) error {
	return s.save(ctx, data)
}

// 撤回消息
func (s *sTalkMessage) Revoke(ctx context.Context, uid int, recordId int) error {

	record, err := dao.TalkRecords.FindById(ctx, recordId)
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

	if err := dao.TalkRecords.UpdateById(ctx, recordId, bson.M{"is_revoke": 1}); err != nil {
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

// 投票
func (s *sTalkMessage) Vote(ctx context.Context, uid int, recordId int, optionsValue string) (*model.VoteStatistics, error) {

	talkRecords, err := dao.TalkRecords.FindByRecordId(ctx, recordId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	talkRecordsVote, err := dao.TalkRecordsVote.FindOne(ctx, bson.M{"record_id": talkRecords.RecordId})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	vote := &model.QueryVoteModel{
		ReceiverId:   talkRecords.ReceiverId,
		TalkType:     talkRecords.TalkType,
		MsgType:      talkRecords.MsgType,
		VoteId:       talkRecordsVote.Id,
		RecordId:     talkRecordsVote.RecordId,
		AnswerMode:   talkRecordsVote.AnswerMode,
		AnswerOption: talkRecordsVote.AnswerOption,
		AnswerNum:    talkRecordsVote.AnswerNum,
		VoteStatus:   talkRecordsVote.Status,
	}

	if vote.MsgType != consts.ChatMsgTypeVote {
		return nil, errors.Newf("当前记录不属于投票信息[%d]", vote.MsgType)
	}

	if vote.TalkType == consts.ChatGroupMode {
		count, err := dao.GroupMember.CountDocuments(ctx, bson.M{"group_id": vote.ReceiverId, "user_id": uid, "is_quit": 0})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if count == 0 {
			return nil, errors.New("暂无投票权限")
		}
	}

	count, err := dao.CountDocuments(ctx, dao.TalkRecordsVote.Database, do.TALK_RECORDS_VOTE_ANSWER_COLLECTION, bson.M{"vote_id": vote.VoteId, "user_id": uid})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if count > 0 {
		return nil, errors.Newf("重复投票[%d]", vote.VoteId)
	}

	options := strings.Split(optionsValue, ",")
	sort.Strings(options)

	var answerOptions map[string]any
	if err := gjson.Unmarshal([]byte(vote.AnswerOption), &answerOptions); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, option := range options {
		if _, ok := answerOptions[option]; !ok {
			return nil, errors.Newf("投票选项不合法[%s]", option)
		}
	}

	if vote.AnswerMode == consts.VoteAnswerModeSingleChoice {
		options = options[:1]
	}

	answers := make([]interface{}, 0, len(options))
	for _, option := range options {
		answers = append(answers, &do.TalkRecordsVoteAnswer{
			VoteId: vote.VoteId,
			UserId: uid,
			Option: option,
		})
	}

	if err = dao.TalkRecordsVote.UpdateById(ctx, vote.VoteId, bson.M{
		"$inc": bson.M{
			"answered_num": 1,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if talkRecordsVote, err := dao.TalkRecordsVote.FindById(ctx, vote.VoteId); err != nil {
		logger.Error(ctx, err)
		return nil, err
	} else if talkRecordsVote.AnsweredNum >= talkRecordsVote.AnswerNum {
		if err = dao.TalkRecordsVote.UpdateById(ctx, vote.VoteId, bson.M{
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

	if _, err = dao.TalkRecordsVote.SetVoteAnswerUser(ctx, vote.VoteId); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if _, err = dao.TalkRecordsVote.SetVoteStatistics(ctx, vote.VoteId); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	info, err := dao.TalkRecordsVote.GetVoteStatistics(ctx, vote.VoteId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return info, nil
}

func (s *sTalkMessage) save(ctx context.Context, data *model.TalkRecords) error {

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
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	data.Id = id

	option := make(map[string]string)

	switch data.MsgType {
	case consts.ChatMsgTypeText:
		option["text"] = util.MtSubstr(util.ReplaceImgAll(data.Content), 0, 300)
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

func (s *sTalkMessage) loadSequence(ctx context.Context, data *model.TalkRecords) {
	if data.TalkType == consts.ChatGroupMode {
		data.Sequence = dao.Sequence.Get(ctx, 0, data.ReceiverId)
	} else {
		data.Sequence = dao.Sequence.Get(ctx, data.UserId, data.ReceiverId)
	}
}

func (s *sTalkMessage) loadReply(ctx context.Context, data *model.TalkRecords) {

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

	reply := model.Reply{
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
func (s *sTalkMessage) afterHandle(ctx context.Context, record *model.TalkRecords, opt map[string]string) {

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
