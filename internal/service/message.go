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
	IMessage interface {
		// 发送文本消息
		Text(ctx context.Context, params model.TextMessageReq) error
		// 发送代码块消息
		Code(ctx context.Context, params model.CodeMessageReq) error
		// 发送图片消息
		Image(ctx context.Context, params model.ImageMessageReq) error
		// 发送文件消息
		File(ctx context.Context, params model.FileMessageReq) error
		// 发送投票消息
		Vote(ctx context.Context, params model.VoteMessageReq) error
		// 发送表情包消息
		Emoticon(ctx context.Context, params model.EmoticonMessageReq) error
		// 发送转发消息
		Forward(ctx context.Context, params model.ForwardMessageReq) error
		// 发送用户名片消息
		Card(ctx context.Context, params model.CardMessageReq) error
		// 收藏聊天图片
		Collect(ctx context.Context, params model.CollectMessageReq) error
		// 撤销聊天记录
		Revoke(ctx context.Context, params model.RevokeMessageReq) error
		// 删除聊天记录
		Delete(ctx context.Context, params model.DeleteMessageReq) error
		// 投票处理
		HandleVote(ctx context.Context, params model.VoteMessageHandleReq) (*model.VoteStatistics, error)
		// 发送位置消息
		Location(ctx context.Context, params model.LocationMessageReq) error
	}
)

var (
	localMessage IMessage
)

func Message() IMessage {
	if localMessage == nil {
		panic("implement not found for interface IMessage, forgot register?")
	}
	return localMessage
}

func RegisterMessage(i IMessage) {
	localMessage = i
}
