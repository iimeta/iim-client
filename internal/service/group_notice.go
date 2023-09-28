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
		// 群公告列表
		List(ctx context.Context, params model.NoticeListReq) (*model.NoticeListRes, error)
		// 发布或更新群公告
		CreateAndUpdate(ctx context.Context, params model.NoticeEditReq) error
		// 删除群公告
		Delete(ctx context.Context, params model.NoticeDeleteReq) error
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
