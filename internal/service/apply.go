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
	IApply interface {
		// ApplyUnreadNum 获取好友申请未读数
		ApplyUnreadNum(ctx context.Context) (int, error)
		// Create 创建联系人申请
		Create(ctx context.Context, params model.ContactApplyCreateReq) error
		// Accept 同意联系人添加申请
		Accept(ctx context.Context, params model.ContactApplyAcceptReq) error
		// Decline 拒绝联系人添加申请
		Decline(ctx context.Context, params model.ContactApplyDeclineReq) error
		// List 获取联系人申请列表
		List(ctx context.Context) (*model.ContactApplyListRes, error)
	}
)

var (
	localApply IApply
)

func Apply() IApply {
	if localApply == nil {
		panic("implement not found for interface IApply, forgot register?")
	}
	return localApply
}

func RegisterApply(i IApply) {
	localApply = i
}
