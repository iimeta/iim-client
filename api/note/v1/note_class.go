package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 笔记分类列表接口请求参数
type ClassListReq struct {
	g.Meta `path:"/class/list" tags:"note_class" method:"get" summary:"笔记分类列表接口"`
}

// 笔记分类列表接口响应参数
type ClassListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ClassListRes
}

// 笔记分类编辑接口请求参数
type ClassEditReq struct {
	g.Meta `path:"/class/editor" tags:"note_class" method:"post" summary:"笔记分类编辑接口"`
	model.ClassEditReq
}

// 笔记分类编辑接口响应参数
type ClassEditRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ClassEditRes
}

// 笔记分类删除接口请求参数
type ClassDeleteReq struct {
	g.Meta `path:"/class/delete" tags:"note_class" method:"post" summary:"笔记分类删除接口"`
	model.ClassDeleteReq
}

// 笔记分类删除接口响应参数
type ClassDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 笔记分类排序接口请求参数
type ClassSortReq struct {
	g.Meta `path:"/class/sort" tags:"note_class" method:"post" summary:"笔记分类排序接口"`
	model.ClassSortReq
}

// 笔记分类排序接口响应参数
type ClassSortRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
