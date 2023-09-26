// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/iim-client/internal/model"
)

type (
	ITalkMessage interface {
		// 系统文本消息
		SendSystemText(ctx context.Context, uid int, req *model.TextMessageReq) error
		// 文本消息
		SendText(ctx context.Context, uid int, req *model.TextMessageReq) error
		// 图片文件消息
		SendImage(ctx context.Context, uid int, req *model.ImageMessageReq) error
		// 语音文件消息
		SendVoice(ctx context.Context, uid int, req *model.VoiceMessageReq) error
		// 视频文件消息
		SendVideo(ctx context.Context, uid int, req *model.VideoMessageReq) error
		// 文件消息
		SendFile(ctx context.Context, uid int, req *model.FileMessageReq) error
		// 代码消息
		SendCode(ctx context.Context, uid int, req *model.CodeMessageReq) error
		// 投票消息
		SendVote(ctx context.Context, uid int, req *model.VoteMessageReq) error
		// 表情消息
		SendEmoticon(ctx context.Context, uid int, req *model.EmoticonMessageReq) error
		// 转发消息
		SendForward(ctx context.Context, uid int, req *model.ForwardMessageReq) error
		// 位置消息
		SendLocation(ctx context.Context, uid int, req *model.LocationMessageReq) error
		// 推送用户名片消息
		SendBusinessCard(ctx context.Context, uid int, req *model.CardMessageReq) error
		// 推送用户登录消息
		SendLogin(ctx context.Context, uid int, req *model.LoginMessageReq) error
		// 图文消息
		SendMixedMessage(ctx context.Context, uid int, req *model.MixedMessageReq) error
		// 推送其它消息
		SendSysOther(ctx context.Context, data *model.TalkRecords) error
		// 撤回消息
		Revoke(ctx context.Context, uid int, recordId int) error
		// 投票
		Vote(ctx context.Context, uid int, recordId int, optionsValue string) (*model.VoteStatistics, error)
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
