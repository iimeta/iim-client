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
		// 创建好友申请
		Create(ctx context.Context, apply *model.Apply) (string, error)
		// 同意好友申请
		Accept(ctx context.Context, apply *model.Apply) (*model.Apply, error)
		// 拒绝好友申请
		Decline(ctx context.Context, apply *model.Apply) error
		// 好友申请列表
		List(ctx context.Context, uid int) ([]*model.ApplyItem, error)
		// 获取申请未读数
		GetApplyUnreadNum(ctx context.Context, uid int) int
		// 清除申请未读数
		ClearApplyUnreadNum(ctx context.Context, uid int)
		// 获取好友申请未读数
		ApplyUnreadNum(ctx context.Context) (int, error)
		// 创建好友申请
		ApplyCreate(ctx context.Context, params model.ApplyCreateReq) error
		// 同意好友添加申请
		ApplyAccept(ctx context.Context, params model.ApplyAcceptReq) error
		// 拒绝好友添加申请
		ApplyDecline(ctx context.Context, params model.ApplyDeclineReq) error
		// 获取好友申请列表
		ApplyList(ctx context.Context) (*model.ApplyListRes, error)
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
