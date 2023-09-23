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
	IGroupNotice interface {
		// 创建群公告
		GroupNoticeCreate(ctx context.Context, edit *model.GroupNoticeEdit) error
		// 更新群公告
		GroupNoticeUpdate(ctx context.Context, edit *model.GroupNoticeEdit) error
		GroupNoticeDelete(ctx context.Context, groupId int, noticeId string) error
		// 添加或编辑群公告
		CreateAndUpdate(ctx context.Context, params model.GroupNoticeEditReq) (string, error)
		// 删除群公告
		Delete(ctx context.Context, params model.GroupNoticeDeleteReq) (string, error)
		// 获取群公告列表(所有)
		List(ctx context.Context, params model.GroupNoticeListReq) (*model.GroupNoticeListRes, error)
	}
)

var (
	localGroupNotice IGroupNotice
)

func GroupNotice() IGroupNotice {
	if localGroupNotice == nil {
		panic("implement not found for interface IGroupNotice, forgot register?")
	}
	return localGroupNotice
}

func RegisterGroupNotice(i IGroupNotice) {
	localGroupNotice = i
}
