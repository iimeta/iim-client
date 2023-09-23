package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 文章标签列表接口请求参数
type ArticleTagListReq struct {
	g.Meta `path:"/tag/list" tags:"note_tag" method:"get" summary:"文章标签列表接口"`
}

// 文章标签列表接口响应参数
type ArticleTagListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ArticleTagListRes
}

// 文章标签编辑接口请求参数
type ArticleTagEditReq struct {
	g.Meta `path:"/tag/editor" tags:"note_tag" method:"post" summary:"文章标签编辑接口"`
	model.ArticleTagEditReq
}

// 文章标签编辑接口请求参数
type ArticleTagEditRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ArticleTagEditRes
}

// 文章标签删除接口请求参数
type ArticleTagDeleteReq struct {
	g.Meta `path:"/tag/delete" tags:"note_tag" method:"post" summary:"文章标签删除接口"`
	model.ArticleTagDeleteReq
}

// 文章标签删除接口响应参数
type ArticleTagDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
