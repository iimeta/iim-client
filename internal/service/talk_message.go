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
		SendMessage(ctx context.Context, message *model.Message) (err error)
		// 发送系统消息
		SendSysMessage(ctx context.Context, message *model.SysMessage) error
		// 发送通知消息
		SendNoticeMessage(ctx context.Context, message *model.NoticeMessage) error
		// 文本消息
		SendText(ctx context.Context, uid int, req *model.TextMessageReq) error
		// 图片文件消息
		SendImage(ctx context.Context, uid int, req *model.ImageMessageReq) error
		// 文件消息
		SendFile(ctx context.Context, uid int, req *model.MessageFileReq) error
		// 投票消息
		SendVote(ctx context.Context, uid int, req *model.MessageVoteReq) error
		// 撤回消息
		Revoke(ctx context.Context, params model.MessageRevokeReq) error
		// 发送图片消息
		Image(ctx context.Context, params model.ImageMessageReq) error
		// 发送文件消息
		File(ctx context.Context, params model.MessageFileReq) error
		// 发送投票消息
		Vote(ctx context.Context, params model.MessageVoteReq) error
		// 投票处理
		HandleVote(ctx context.Context, params model.MessageVoteHandleReq) (*model.VoteStatistics, error)
		// 删除消息记录
		Delete(ctx context.Context, params model.MessageDeleteReq) error
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
		TextMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
		CodeMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
		ImageMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
		VoiceMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
		VideoMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
		FileMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
		VoteMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
		MixedMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
		// todo
		ForwardMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
		EmoticonMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
		CardMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
		LocationMessageHandler(ctx context.Context, message *model.Message) (*model.TalkRecord, error)
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
