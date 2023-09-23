package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 添加或移出表情包接口请求参数
type EmoticonSetSystemReq struct {
	g.Meta `path:"/system/install" tags:"emoticon" method:"post" summary:"添加或移出表情包接口"`
	model.EmoticonSetSystemReq
}

// 添加或移出表情包接口响应参数
type EmoticonSetSystemRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.EmoticonSetSystemRes
}

// 删除表情包接口请求参数
type EmoticonDeleteReq struct {
	g.Meta `path:"/customize/delete" tags:"emoticon" method:"post" summary:"删除表情包接口"`
	model.EmoticonDeleteReq
}

// 删除表情包接口响应参数
type EmoticonDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 系统表情包列表接口请求参数
type EmoticonSysListReq struct {
	g.Meta `path:"/system/list" tags:"emoticon" method:"get" summary:"系统表情包列表接口"`
}

// 系统表情包列表接口响应参数
type EmoticonSysListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.EmoticonSysListRes
}

// 用户表情包列表接口请求参数
type EmoticonListReq struct {
	g.Meta `path:"/list" tags:"emoticon" method:"get" summary:"用户表情包列表接口"`
}

// 用户表情包列表接口响应参数
type EmoticonListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.EmoticonListRes
}

// 表情包上传接口请求参数
type EmoticonUploadReq struct {
	g.Meta `path:"/customize/create" tags:"emoticon" method:"post" summary:"表情包上传接口"`
}

// 表情包上传接口响应参数
type EmoticonUploadRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.EmoticonUploadRes
}
