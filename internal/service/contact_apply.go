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
		Create(ctx context.Context, apply *model.ContactApply) (string, error)
		// 同意好友申请
		Accept(ctx context.Context, apply *model.ContactApply) (*model.ContactApply, error)
		// 拒绝好友申请
		Decline(ctx context.Context, apply *model.ContactApply) error
		// 好友申请列表
		List(ctx context.Context, uid int) ([]*model.ApplyItem, error)
		// 获取申请未读数
		GetApplyUnreadNum(ctx context.Context, uid int) int
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
