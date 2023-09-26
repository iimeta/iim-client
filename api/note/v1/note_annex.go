package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 笔记附件上传接口请求参数
type AnnexUploadReq struct {
	g.Meta `path:"/annex/upload" tags:"note_annex" method:"post" summary:"笔记附件上传接口"`
	model.AnnexUploadReq
}

// 笔记附件上传接口响应参数
type AnnexUploadRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AnnexUploadRes
}

// 笔记附件删除接口请求参数
type AnnexDeleteReq struct {
	g.Meta `path:"/annex/delete" tags:"note_annex" method:"post" summary:"笔记附件删除接口"`
	model.AnnexDeleteReq
}

// 笔记附件删除接口响应参数
type AnnexDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 笔记附件恢复删除接口请求参数
type AnnexRecoverReq struct {
	g.Meta `path:"/annex/recover" tags:"note_annex" method:"post" summary:"笔记附件恢复删除接口"`
	model.AnnexRecoverReq
}

// 笔记附件恢复删除接口响应参数
type AnnexRecoverRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 笔记附件永久删除接口请求参数
type AnnexForeverDeleteReq struct {
	g.Meta `path:"/annex/forever/delete" tags:"note_annex" method:"post" summary:"笔记附件永久删除接口"`
	model.AnnexForeverDeleteReq
}

// 笔记附件永久删除接口响应参数
type AnnexForeverDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 笔记附件下载接口请求参数
type AnnexDownloadReq struct {
	g.Meta `path:"/annex/download" tags:"note_annex" method:"get" summary:"笔记附件下载接口"`
	model.AnnexDownloadReq
}

// 笔记附件下载接口响应参数
type AnnexDownloadRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 笔记附件回收站列表接口请求参数
type AnnexRecoverListReq struct {
	g.Meta `path:"/annex/recover/list" tags:"note_annex" method:"get" summary:"笔记附件回收站列表接口"`
}

// 笔记附件回收站列表接口响应参数
type AnnexRecoverListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AnnexRecoverListRes
}
