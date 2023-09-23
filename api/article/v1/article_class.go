package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 文章分类列表接口请求参数
type ArticleClassListReq struct {
	g.Meta `path:"/class/list" tags:"note_tag" method:"get" summary:"文章分类列表接口"`
}

// 文章分类列表接口响应参数
type ArticleClassListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ArticleClassListRes
}

// 文章分类编辑接口请求参数
type ArticleClassEditReq struct {
	g.Meta `path:"/class/editor" tags:"note_tag" method:"post" summary:"文章分类编辑接口"`
	model.ArticleClassEditReq
}

// 文章分类编辑接口响应参数
type ArticleClassEditRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ArticleClassEditRes
}

// 文章分类删除接口请求参数
type ArticleClassDeleteReq struct {
	g.Meta `path:"/class/delete" tags:"note_tag" method:"post" summary:"文章分类删除接口"`
	model.ArticleClassDeleteReq
}

// 文章分类删除接口响应参数
type ArticleClassDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 文章分类排序接口请求参数
type ArticleClassSortReq struct {
	g.Meta `path:"/class/sort" tags:"note_tag" method:"post" summary:"文章分类排序接口"`
	model.ArticleClassSortReq
}

// 文章分类排序接口响应参数
type ArticleClassSortRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
