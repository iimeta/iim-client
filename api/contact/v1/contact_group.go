package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 添加联系人分组接口请求参数
type ContactGroupCreateReq struct {
	g.Meta `path:"/group/create" tags:"contact" method:"post" summary:"添加联系人分组接口"`
	Name   string `json:"name,omitempty" v:"required"`
	Sort   int32  `json:"sort,omitempty" v:"required"`
}

// 添加联系人分组接口响应参数
type ContactGroupCreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
	Id     int32 `json:"id,omitempty"`
}

// 更新联系人分组接口请求参数
type ContactGroupUpdateReq struct {
	g.Meta `path:"/group/update" tags:"contact" method:"post" summary:"更新联系人分组接口"`
	Id     int32  `json:"id,omitempty" v:"required"`
	Name   string `json:"name,omitempty" v:"required"`
	Sort   int32  `json:"sort,omitempty" v:"required"`
}

// 更新联系人分组接口响应参数
type ContactGroupUpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
	Id     int32 `json:"id,omitempty"`
}

// 删除联系人分组接口请求参数
type ContactGroupDeleteReq struct {
	g.Meta `path:"/group/delete" tags:"contact" method:"post" summary:"删除联系人分组接口"`
	Id     int32 `json:"id,omitempty" v:"required"`
}

// 删除联系人分组接口响应参数
type ContactGroupDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
	Id     int32 `json:"id,omitempty"`
}

// 排序联系人分组接口请求参数
type ContactGroupSortReq struct {
	g.Meta `path:"/group/sort" tags:"contact" method:"post" summary:"排序联系人分组接口"`
	Items  []*ContactGroupSortRequest_Item `json:"items" v:"required"`
}

// 排序联系人分组接口响应参数
type ContactGroupSortRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 联系人分组列表接口请求参数
type ContactGroupListReq struct {
	g.Meta `path:"/group/list" tags:"contact" method:"get" summary:"联系人分组列表接口"`
}

// 联系人分组列表接口响应参数
type ContactGroupListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ContactGroupListRes
}

// 保存联系人分组列表接口请求参数
type ContactGroupSaveReq struct {
	g.Meta `path:"/group/save" tags:"contact" method:"post" summary:"保存联系人分组列表接口"`
	model.ContactGroupSaveReq
}

// 保存联系人分组列表接口响应参数
type ContactGroupSaveRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

type ContactGroupSortRequest_Item struct {
	Id   int32 `json:"id,omitempty" v:"required"`
	Sort int32 `json:"sort,omitempty" v:"required"`
}
