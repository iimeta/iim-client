package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 会话面板记录接口请求参数
type GetRecordsReq struct {
	g.Meta `path:"/records" tags:"talk" method:"get" summary:"会话面板记录接口"`
	model.GetTalkRecordsReq
}

// 会话面板记录接口响应参数
type GetRecordsRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GetTalkRecordsRes
}

// 历史会话记录接口请求参数
type SearchHistoryRecordsReq struct {
	g.Meta `path:"/records/history" tags:"talk" method:"get" summary:"历史会话记录接口"`
	model.GetTalkRecordsReq
}

// 历史会话记录接口响应参数
type SearchHistoryRecordsRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GetTalkRecordsRes
}

// 会话转发记录接口请求参数
type GetForwardRecordsReq struct {
	g.Meta `path:"/records/forward" tags:"talk" method:"get" summary:"会话转发记录接口"`
	model.GetForwardTalkRecordReq
}

// 会话转发记录接口响应参数
type GetForwardRecordsRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GetTalkRecordsRes
}

// 会话文件下载接口请求参数
type RecordsFileDownloadReq struct {
	g.Meta `path:"/records/file/download" tags:"talk" method:"get" summary:"会话文件下载接口"`
	model.DownloadChatFileReq
}

// 会话文件下载接口响应参数
type RecordsFileDownloadRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
