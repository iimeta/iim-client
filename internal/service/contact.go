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
	IContact interface {
		// List 联系人列表
		List(ctx context.Context) (*model.ContactListRes, error)
		// Delete 删除联系人
		Delete(ctx context.Context, params model.ContactDeleteReq) error
		// Search 查找联系人
		Search(ctx context.Context, params model.ContactSearchReq) (*model.ContactSearchRes, error)
		// Remark 编辑联系人备注
		Remark(ctx context.Context, params model.ContactEditRemarkReq) error
		// Detail 联系人详情信息
		Detail(ctx context.Context, params model.ContactDetailReq) (*model.ContactDetailRes, error)
		// MoveGroup 移动好友分组
		MoveGroup(ctx context.Context, params model.ContactChangeGroupReq) error
	}
)

var (
	localContact IContact
)

func Contact() IContact {
	if localContact == nil {
		panic("implement not found for interface IContact, forgot register?")
	}
	return localContact
}

func RegisterContact(i IContact) {
	localContact = i
}
