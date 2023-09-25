package talk

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/robot"
)

type sPublish struct{}

func init() {
	service.RegisterPublish(New())
}

func New() service.IPublish {
	return &sPublish{}
}

var mapping map[string]func(ctx context.Context) error

// 发送消息接口
func (s *sPublish) Publish(ctx context.Context, params model.PublishBaseMessageReq) error {

	// todo
	if err := service.Group().IsAuth(ctx, &model.AuthOption{
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
func (s *sPublish) onSendText(ctx context.Context) error {

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

			session, err := service.Session().FindBySession(ctx, uid, textMessageReq.Receiver.ReceiverId, textMessageReq.Receiver.TalkType)
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
func (s *sPublish) onSendImage(ctx context.Context) error {

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
func (s *sPublish) onSendVoice(ctx context.Context) error {

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
func (s *sPublish) onSendVideo(ctx context.Context) error {

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
func (s *sPublish) onSendFile(ctx context.Context) error {

	fileMessageReq := &model.FileMessageReq{}
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
func (s *sPublish) onSendCode(ctx context.Context) error {

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
func (s *sPublish) onSendLocation(ctx context.Context) error {

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
func (s *sPublish) onSendForward(ctx context.Context) error {

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
func (s *sPublish) onSendEmoticon(ctx context.Context) error {

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
func (s *sPublish) onSendVote(ctx context.Context) error {

	voteMessageReq := &model.VoteMessageReq{}
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
func (s *sPublish) onSendCard(ctx context.Context) error {

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
func (s *sPublish) onMixedMessage(ctx context.Context) error {

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

func (s *sPublish) transfer(ctx context.Context, typeValue string) error {

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
