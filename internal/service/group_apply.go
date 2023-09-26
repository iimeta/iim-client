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
	IGroupApply interface {
		Create(ctx context.Context, params model.GroupApplyCreateReq) error
		Agree(ctx context.Context, params model.ApplyAgreeReq) error
		Decline(ctx context.Context, params model.GroupApplyDeclineReq) error
		List(ctx context.Context, params model.ApplyListReq) (*model.GroupApplyListRes, error)
		All(ctx context.Context) (*model.ApplyAllRes, error)
		ApplyUnreadNum(ctx context.Context) (*model.GroupApplyUnreadNumRes, error)
		Delete(ctx context.Context, applyId string, userId int) error
	}
)

var (
	localGroupApply IGroupApply
)

func GroupApply() IGroupApply {
	if localGroupApply == nil {
		panic("implement not found for interface IGroupApply, forgot register?")
	}
	return localGroupApply
}

func RegisterGroupApply(i IGroupApply) {
	localGroupApply = i
}
