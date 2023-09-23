// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package message

import (
	"context"

	"github.com/iimeta/iim-client/api/message/v1"
)

type IMessageV1 interface {
	FileMessage(ctx context.Context, req *v1.FileMessageReq) (res *v1.FileMessageRes, err error)
	VoteMessage(ctx context.Context, req *v1.VoteMessageReq) (res *v1.VoteMessageRes, err error)
	VoteMessageHandle(ctx context.Context, req *v1.VoteMessageHandleReq) (res *v1.VoteMessageHandleRes, err error)
	PublishBaseMessage(ctx context.Context, req *v1.PublishBaseMessageReq) (res *v1.PublishBaseMessageRes, err error)
	CollectMessage(ctx context.Context, req *v1.CollectMessageReq) (res *v1.CollectMessageRes, err error)
	RevokeMessage(ctx context.Context, req *v1.RevokeMessageReq) (res *v1.RevokeMessageRes, err error)
	DeleteMessage(ctx context.Context, req *v1.DeleteMessageReq) (res *v1.DeleteMessageRes, err error)
}
