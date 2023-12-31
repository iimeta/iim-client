package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 添加好友分组接口请求参数
type GroupCreateReq struct {
	g.Meta `path:"/group/create" tags:"contact_group" method:"post" summary:"添加好友分组接口"`
	model.ContactGroupCreateReq
}

// 添加好友分组接口响应参数
type GroupCreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ContactGroupCreateRes
}

// 更新好友分组接口请求参数
type GroupUpdateReq struct {
	g.Meta `path:"/group/update" tags:"contact_group" method:"post" summary:"更新好友分组接口"`
	model.GroupUpdateReq
}

// 更新好友分组接口响应参数
type GroupUpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupUpdateRes
}

// 删除好友分组接口请求参数
type GroupDeleteReq struct {
	g.Meta `path:"/group/delete" tags:"contact_group" method:"post" summary:"删除好友分组接口"`
	model.GroupDeleteReq
}

// 删除好友分组接口响应参数
type GroupDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupDeleteRes
}

// 排序好友分组接口请求参数
type GroupSortReq struct {
	g.Meta `path:"/group/sort" tags:"contact_group" method:"post" summary:"排序好友分组接口"`
	model.GroupSortReq
}

// 排序好友分组接口响应参数
type GroupSortRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 好友分组列表接口请求参数
type GroupListReq struct {
	g.Meta `path:"/group/list" tags:"contact_group" method:"get" summary:"好友分组列表接口"`
}

// 好友分组列表接口响应参数
type GroupListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ContactGroupListRes
}

// 保存好友分组列表接口请求参数
type GroupSaveReq struct {
	g.Meta `path:"/group/save" tags:"contact_group" method:"post" summary:"保存好友分组列表接口"`
	model.GroupSaveReq
}

// 保存好友分组列表接口响应参数
type GroupSaveRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
