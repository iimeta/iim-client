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
	IMessage interface {
		// Text 发送文本消息
		Text(ctx context.Context, params model.TextMessageReq) error
		// Code 发送代码块消息
		Code(ctx context.Context, params model.CodeMessageReq) error
		// Image 发送图片消息
		Image(ctx context.Context, params model.ImageMessageReq) error
		// File 发送文件消息
		File(ctx context.Context, params model.FileMessageReq) error
		// Vote 发送投票消息
		Vote(ctx context.Context, params model.VoteMessageReq) error
		// Emoticon 发送表情包消息
		Emoticon(ctx context.Context, params model.EmoticonMessageReq) error
		// Forward 发送转发消息
		Forward(ctx context.Context, params model.ForwardMessageReq) error
		// Card 发送用户名片消息
		Card(ctx context.Context, params model.CardMessageReq) error
		// Collect 收藏聊天图片
		Collect(ctx context.Context, params model.CollectMessageReq) error
		// Revoke 撤销聊天记录
		Revoke(ctx context.Context, params model.RevokeMessageReq) error
		// Delete 删除聊天记录
		Delete(ctx context.Context, params model.DeleteMessageReq) error
		// HandleVote 投票处理
		HandleVote(ctx context.Context, params model.VoteMessageHandleReq) (*dao.VoteStatistics, error)
		// Location 发送位置消息
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
