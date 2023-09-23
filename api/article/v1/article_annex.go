package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 文章附件上传接口请求参数
type ArticleAnnexUploadReq struct {
	g.Meta `path:"/annex/upload" tags:"note_tag" method:"post" summary:"文章附件上传接口"`
	model.ArticleAnnexUploadReq
}

// 文章附件上传接口响应参数
type ArticleAnnexUploadRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ArticleAnnexUploadRes
}

// 文章附件删除接口请求参数
type ArticleAnnexDeleteReq struct {
	g.Meta `path:"/annex/delete" tags:"note_tag" method:"post" summary:"文章附件删除接口"`
	model.ArticleAnnexDeleteReq
}

// 文章附件删除接口响应参数
type ArticleAnnexDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 文章附件恢复删除接口请求参数
type ArticleAnnexRecoverReq struct {
	g.Meta `path:"/annex/recover" tags:"note_tag" method:"post" summary:"文章附件恢复删除接口"`
	model.ArticleAnnexRecoverReq
}

// 文章附件恢复删除接口响应参数
type ArticleAnnexRecoverRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 文章附件永久删除接口请求参数
type ArticleAnnexForeverDeleteReq struct {
	g.Meta `path:"/annex/forever/delete" tags:"note_tag" method:"post" summary:"文章附件永久删除接口"`
	model.ArticleAnnexForeverDeleteReq
}

// 文章附件永久删除接口响应参数
type ArticleAnnexForeverDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 文章附件下载接口请求参数
type ArticleAnnexDownloadReq struct {
	g.Meta `path:"/annex/download" tags:"note_tag" method:"get" summary:"文章附件下载接口"`
	model.ArticleAnnexDownloadReq
}

// 文章附件下载接口响应参数
type ArticleAnnexDownloadRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 文章附件回收站列表接口请求参数
type ArticleAnnexRecoverListReq struct {
	g.Meta `path:"/annex/recover/list" tags:"note_tag" method:"get" summary:"文章附件回收站列表接口"`
}

// 文章附件回收站列表接口响应参数
type ArticleAnnexRecoverListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ArticleAnnexRecoverListRes
}
