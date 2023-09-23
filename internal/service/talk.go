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
	ITalk interface {
		// DeleteRecordList 删除消息记录
		DeleteRecordList(ctx context.Context, opt *model.RemoveRecordListOpt) error
		// Collect 收藏表情包
		Collect(ctx context.Context, uid int, recordId int) error
	}
)

var (
	localTalk ITalk
)

func Talk() ITalk {
	if localTalk == nil {
		panic("implement not found for interface ITalk, forgot register?")
	}
	return localTalk
}

func RegisterTalk(i ITalk) {
	localTalk = i
}
