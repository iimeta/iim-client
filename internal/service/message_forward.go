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
	IMessageForward interface {
		// Verify 验证转发消息合法性
		Verify(ctx context.Context, uid int, params *model.ForwardMessageReq) error
		// MultiMergeForward 批量合并转发
		MultiMergeForward(ctx context.Context, uid int, params *model.ForwardMessageReq) ([]*model.ForwardRecord, error)
		// MultiSplitForward 批量逐条转发
		MultiSplitForward(ctx context.Context, uid int, params *model.ForwardMessageReq) ([]*model.ForwardRecord, error)
	}
)

var (
	localMessageForward IMessageForward
)

func MessageForward() IMessageForward {
	if localMessageForward == nil {
		panic("implement not found for interface IMessageForward, forgot register?")
	}
	return localMessageForward
}

func RegisterMessageForward(i IMessageForward) {
	localMessageForward = i
}
