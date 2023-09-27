package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 会话面板记录接口请求参数
type RecordsReq struct {
	g.Meta `path:"/records" tags:"talk_records" method:"get" summary:"会话面板记录接口"`
	model.TalkRecordsReq
}

// 会话面板记录接口响应参数
type RecordsRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TalkRecordsRes
}

// 历史会话记录接口请求参数
type RecordsSearchHistoryReq struct {
	g.Meta `path:"/records/history" tags:"talk_records" method:"get" summary:"历史会话记录接口"`
	model.TalkRecordsReq
}

// 历史会话记录接口响应参数
type RecordsSearchHistoryRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TalkRecordsRes
}

// 会话转发记录接口请求参数
type RecordsForwardReq struct {
	g.Meta `path:"/records/forward" tags:"talk_records" method:"get" summary:"会话转发记录接口"`
	model.RecordsForwardReq
}

// 会话转发记录接口响应参数
type RecordsForwardRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TalkRecordsRes
}

// 会话文件下载接口请求参数
type RecordsFileDownloadReq struct {
	g.Meta `path:"/records/file/download" tags:"talk_records" method:"get" summary:"会话文件下载接口"`
	model.RecordsFileDownloadReq
}

// 会话文件下载接口响应参数
type RecordsFileDownloadRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
