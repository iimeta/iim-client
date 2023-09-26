package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 添加或移出表情包接口请求参数
type SetSystemReq struct {
	g.Meta `path:"/system/install" tags:"emoticon" method:"post" summary:"添加或移出表情包接口"`
	model.SetSystemReq
}

// 添加或移出表情包接口响应参数
type SetSystemRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SetSystemRes
}

// 删除表情包接口请求参数
type DeleteReq struct {
	g.Meta `path:"/customize/delete" tags:"emoticon" method:"post" summary:"删除表情包接口"`
	model.DeleteReq
}

// 删除表情包接口响应参数
type DeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 系统表情包列表接口请求参数
type SysListReq struct {
	g.Meta `path:"/system/list" tags:"emoticon" method:"get" summary:"系统表情包列表接口"`
}

// 系统表情包列表接口响应参数
type SysListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SysListRes
}

// 用户表情包列表接口请求参数
type ListReq struct {
	g.Meta `path:"/list" tags:"emoticon" method:"get" summary:"用户表情包列表接口"`
}

// 用户表情包列表接口响应参数
type ListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ListRes
}

// 表情包上传接口请求参数
type UploadReq struct {
	g.Meta `path:"/customize/create" tags:"emoticon" method:"post" summary:"表情包上传接口"`
}

// 表情包上传接口响应参数
type UploadRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.UploadRes
}
