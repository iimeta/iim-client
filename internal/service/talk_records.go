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
	ITalkRecords interface {
		// 获取对话消息
		GetTalkRecords(ctx context.Context, opt *model.QueryTalkRecordsOpt) ([]*model.TalkRecordsItem, error)
		// 对话搜索消息
		SearchTalkRecords()
		GetTalkRecord(ctx context.Context, recordId int) (*model.TalkRecordsItem, error)
		// 获取转发消息记录
		GetForwardRecords(ctx context.Context, uid, recordId int) ([]*model.TalkRecordsItem, error)
		HandleTalkRecords(ctx context.Context, items []*model.TalkRecordsItem) ([]*model.TalkRecordsItem, error)
	}
)

var (
	localTalkRecords ITalkRecords
)

func TalkRecords() ITalkRecords {
	if localTalkRecords == nil {
		panic("implement not found for interface ITalkRecords, forgot register?")
	}
	return localTalkRecords
}

func RegisterTalkRecords(i ITalkRecords) {
	localTalkRecords = i
}
