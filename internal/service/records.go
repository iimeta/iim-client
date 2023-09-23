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
	IRecords interface {
		// GetRecords 获取会话记录
		GetRecords(ctx context.Context, params model.GetTalkRecordsReq) (*model.GetTalkRecordsRes, error)
		// SearchHistoryRecords 查询下会话记录
		SearchHistoryRecords(ctx context.Context, params model.GetTalkRecordsReq) (*model.GetTalkRecordsRes, error)
		// GetForwardRecords 获取转发记录
		GetForwardRecords(ctx context.Context, params model.GetForwardTalkRecordReq) (*model.GetTalkRecordsRes, error)
		// Download 聊天文件下载
		Download(ctx context.Context, recordId int) error
	}
)

var (
	localRecords IRecords
)

func Records() IRecords {
	if localRecords == nil {
		panic("implement not found for interface IRecords, forgot register?")
	}
	return localRecords
}

func RegisterRecords(i IRecords) {
	localRecords = i
}
