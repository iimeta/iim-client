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
	IContactApply interface {
		// 好友申请列表
		List(ctx context.Context) ([]*model.Apply, error)
		// 创建好友申请
		Create(ctx context.Context, params model.ApplyCreateReq) (string, error)
		// 同意好友申请
		Accept(ctx context.Context, params model.ApplyAcceptReq) (*model.ContactApply, error)
		// 拒绝好友申请
		Decline(ctx context.Context, params model.ApplyDeclineReq) error
		// 获取好友申请未读数
		ApplyUnreadNum(ctx context.Context) (int, error)
		// 清除申请未读数
		ClearApplyUnreadNum(ctx context.Context, uid int)
	}
)

var (
	localContactApply IContactApply
)

func ContactApply() IContactApply {
	if localContactApply == nil {
		panic("implement not found for interface IContactApply, forgot register?")
	}
	return localContactApply
}

func RegisterContactApply(i IContactApply) {
	localContactApply = i
}
