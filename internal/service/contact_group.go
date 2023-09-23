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
	IContactGroup interface {
		Delete(ctx context.Context, id int, uid int) error
		// GetUserGroup 用户联系人分组列表
		GetUserGroup(ctx context.Context, uid int) ([]*model.ContactGroup, error)
		// List 联系人分组列表
		List(ctx context.Context) (*model.ContactGroupListRes, error)
		Save(ctx context.Context, params model.ContactGroupSaveReq) error
	}
)

var (
	localContactGroup IContactGroup
)

func ContactGroup() IContactGroup {
	if localContactGroup == nil {
		panic("implement not found for interface IContactGroup, forgot register?")
	}
	return localContactGroup
}

func RegisterContactGroup(i IContactGroup) {
	localContactGroup = i
}
