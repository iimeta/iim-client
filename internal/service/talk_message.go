// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
)

type (
	ITalkMessage interface {
		// SendSystemText 系统文本消息
		SendSystemText(ctx context.Context, uid int, req *model.TextMessageReq) error
		// SendText 文本消息
		SendText(ctx context.Context, uid int, req *model.TextMessageReq) error
		// SendImage 图片文件消息
		SendImage(ctx context.Context, uid int, req *model.ImageMessageReq) error
		// SendVoice 语音文件消息
		SendVoice(ctx context.Context, uid int, req *model.VoiceMessageReq) error
		// SendVideo 视频文件消息
		SendVideo(ctx context.Context, uid int, req *model.VideoMessageReq) error
		// SendFile 文件消息
		SendFile(ctx context.Context, uid int, req *model.FileMessageReq) error
		// SendCode 代码消息
		SendCode(ctx context.Context, uid int, req *model.CodeMessageReq) error
		// SendVote 投票消息
		SendVote(ctx context.Context, uid int, req *model.VoteMessageReq) error
		// SendEmoticon 表情消息
		SendEmoticon(ctx context.Context, uid int, req *model.EmoticonMessageReq) error
		// SendForward 转发消息
		SendForward(ctx context.Context, uid int, req *model.ForwardMessageReq) error
		// SendLocation 位置消息
		SendLocation(ctx context.Context, uid int, req *model.LocationMessageReq) error
		// SendBusinessCard 推送用户名片消息
		SendBusinessCard(ctx context.Context, uid int, req *model.CardMessageReq) error
		// SendLogin 推送用户登录消息
		SendLogin(ctx context.Context, uid int, req *model.LoginMessageReq) error
		// SendMixedMessage 图文消息
		SendMixedMessage(ctx context.Context, uid int, req *model.MixedMessageReq) error
		// SendSysOther 推送其它消息
		SendSysOther(ctx context.Context, data *model.TalkRecords) error
		// Revoke 撤回消息
		Revoke(ctx context.Context, uid int, recordId int) error
		// Vote 投票
		Vote(ctx context.Context, uid int, recordId int, optionsValue string) (*dao.VoteStatistics, error)
	}
)

var (
	localTalkMessage ITalkMessage
)

func TalkMessage() ITalkMessage {
	if localTalkMessage == nil {
		panic("implement not found for interface ITalkMessage, forgot register?")
	}
	return localTalkMessage
}

func RegisterTalkMessage(i ITalkMessage) {
	localTalkMessage = i
}
