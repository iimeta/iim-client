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
		// 校验权限
		VerifyPermission(ctx context.Context, info *model.VerifyInfo) error
		// 发送消息
		SendMessage(ctx context.Context, message *model.Message) error
		// 发送系统消息
		SendSysMessage(ctx context.Context, message *model.SysMessage) error
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
		SendFile(ctx context.Context, uid int, req *model.MessageFileReq) error
		// 代码消息
		SendCode(ctx context.Context, uid int, req *model.CodeMessageReq) error
		// 投票消息
		SendVote(ctx context.Context, uid int, req *model.MessageVoteReq) error
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
		SendSysOther(ctx context.Context, data *model.TalkRecord) error
		// 撤回消息
		Revoke(ctx context.Context, params model.MessageRevokeReq) error
		// 投票处理
		HandleVote(ctx context.Context, params model.MessageVoteHandleReq) (*model.VoteStatistics, error)
		// 发送文本消息
		Text(ctx context.Context, params model.TextMessageReq) error
		// 发送代码块消息
		Code(ctx context.Context, params model.CodeMessageReq) error
		// 发送图片消息
		Image(ctx context.Context, params model.ImageMessageReq) error
		// 发送文件消息
		File(ctx context.Context, params model.MessageFileReq) error
		// 发送投票消息
		Vote(ctx context.Context, params model.MessageVoteReq) error
		// 发送表情包消息
		Emoticon(ctx context.Context, params model.EmoticonMessageReq) error
		// 发送转发消息
		Forward(ctx context.Context, params model.ForwardMessageReq) error
		// 发送用户名片消息
		Card(ctx context.Context, params model.CardMessageReq) error
		// 删除消息记录
		Delete(ctx context.Context, params model.MessageDeleteReq) error
		// 发送位置消息
		Location(ctx context.Context, params model.LocationMessageReq) error
		// 验证转发消息合法性
		Verify(ctx context.Context, uid int, params *model.ForwardMessageReq) error
		// 批量合并转发
		MultiMergeForward(ctx context.Context, uid int, params *model.ForwardMessageReq) ([]*model.ForwardRecord, error)
		// 批量逐条转发
		MultiSplitForward(ctx context.Context, uid int, params *model.ForwardMessageReq) ([]*model.ForwardRecord, error)
		// 收藏表情包
		Collect(ctx context.Context, params model.MessageCollectReq) error
		// 发送消息接口
		Publish(ctx context.Context, params model.MessagePublishReq) error
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
