package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 联系人列表接口请求参数
type ContactListReq struct {
	g.Meta `path:"/list" tags:"contact" method:"get" summary:"联系人列表接口"`
}

// 联系人列表接口响应参数
type ContactListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ContactListRes
}

// 联系人删除接口请求参数
type ContactDeleteReq struct {
	g.Meta `path:"/delete" tags:"contact" method:"post" summary:"联系人删除接口"`
	model.ContactDeleteReq
}

// 联系人删除接口响应参数
type ContactDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 联系人备注修改接口请求参数
type ContactEditRemarkReq struct {
	g.Meta `path:"/edit-remark" tags:"contact" method:"post" summary:"联系人备注修改接口"`
	model.ContactEditRemarkReq
}

// 联系人备注修改接口响应参数
type ContactEditRemarkRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 联系人详情接口请求参数
type ContactDetailReq struct {
	g.Meta `path:"/detail" tags:"contact" method:"get" summary:"联系人详情接口"`
	model.ContactDetailReq
}

// 联系人详情接口响应参数
type ContactDetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ContactDetailRes
}

// 联系人搜索接口请求参数
type ContactSearchReq struct {
	g.Meta `path:"/search" tags:"contact" method:"get" summary:"联系人搜索接口"`
	model.ContactSearchReq
}

// 联系人搜索接口响应参数
type ContactSearchRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ContactSearchRes
}

// 修改联系人分组接口请求参数
type ContactChangeGroupReq struct {
	g.Meta `path:"/move-group" tags:"contact" method:"post" summary:"修改联系人分组接口"`
	model.ContactChangeGroupReq
}

// 修改联系人分组接口响应参数
type ContactChangeGroupRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
