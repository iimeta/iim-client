package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 文章编辑接口请求参数
type ArticleEditReq struct {
	g.Meta `path:"/article/editor" tags:"note_tag" method:"post" summary:"文章编辑接口"`
	model.ArticleEditReq
}

// 文章编辑接口响应参数
type ArticleEditRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ArticleEditRes
}

// 文章详情接口请求参数
type ArticleDetailReq struct {
	g.Meta `path:"/article/detail" tags:"note_tag" method:"get" summary:"文章详情接口"`
	model.ArticleDetailReq
}

// 文章详情接口响应参数
type ArticleDetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ArticleDetailRes
}

// 文章列表接口请求参数
type ArticleListReq struct {
	g.Meta `path:"/article/list" tags:"note_tag" method:"get" summary:"文章列表接口"`
	model.ArticleListReq
}

// 文章列表请求接口响应参数
type ArticleListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ArticleListRes
}

// 文章删除接口请求参数
type ArticleDeleteReq struct {
	g.Meta `path:"/article/delete" tags:"note_tag" method:"post" summary:"文章删除接口"`
	model.ArticleDeleteReq
}

// 文章删除接口响应参数
type ArticleDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 恢复文章接口请求参数
type ArticleRecoverReq struct {
	g.Meta `path:"/article/recover" tags:"note_tag" method:"post" summary:"恢复文章接口"`
	model.ArticleRecoverReq
}

// 恢复文章接口响应参数
type ArticleRecoverRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 文章移动分类接口请求参数
type ArticleMoveReq struct {
	g.Meta `path:"/article/move" tags:"note_tag" method:"post" summary:"文章移动分类接口"`
	model.ArticleMoveReq
}

// 文章移动分类接口响应参数
type ArticleMoveRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 标记文章接口请求参数
type ArticleAsteriskReq struct {
	g.Meta `path:"/article/asterisk" tags:"note_tag" method:"post" summary:"标记文章接口"`
	model.ArticleAsteriskReq
}

// 标记文章接口响应参数
type ArticleAsteriskRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 文章标签接口请求参数
type ArticleTagsReq struct {
	g.Meta `path:"/article/tag" tags:"note_tag" method:"post" summary:"文章标签接口"`
	model.ArticleTagsReq
}

// 文章标签接口响应参数
type ArticleTagsRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 永久删除文章接口请求参数
type ArticleForeverDeleteReq struct {
	g.Meta `path:"/article/forever/delete" tags:"note_tag" method:"post" summary:"永久删除文章接口"`
	model.ArticleForeverDeleteReq
}

// 永久删除文章接口响应参数
type ArticleForeverDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 文章图片上传接口请求参数
type ArticleUploadImageReq struct {
	g.Meta `path:"/article/upload/image" tags:"note_tag" method:"post" summary:"文章图片上传接口"`
	model.ArticleUploadImageReq
}

// 文章图片上传接口响应参数
type ArticleUploadImageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ArticleUploadImageRes
}
