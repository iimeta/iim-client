package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
	"google.golang.org/protobuf/types/known/anypb"
)

// 会话记录
type TalkRecordItem struct {
	Id         int32      `json:"id,omitempty"`
	TalkType   int32      `json:"talk_type,omitempty"`
	ReceiverId int32      `json:"receiver_id,omitempty"`
	IsTop      int32      `json:"is_top,omitempty"`
	IsDisturb  int32      `json:"is_disturb,omitempty"`
	IsOnline   int32      `json:"is_online,omitempty"`
	IsRobot    int32      `json:"is_robot,omitempty"`
	Name       string     `json:"name,omitempty"`
	Avatar     string     `json:"avatar,omitempty"`
	RemarkName string     `json:"remark_name,omitempty"`
	UnreadNum  int32      `json:"unread_num,omitempty"`
	MsgText    string     `json:"msg_text,omitempty"`
	UpdatedAt  string     `json:"updated_at,omitempty"`
	Extra      *anypb.Any `json:"extra,omitempty"`
}

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
