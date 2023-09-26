package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 笔记编辑接口请求参数
type NoteEditReq struct {
	g.Meta `path:"/editor" tags:"note" method:"post" summary:"笔记编辑接口"`
	model.NoteEditReq
}

// 笔记编辑接口响应参数
type NoteEditRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoteEditRes
}

// 笔记详情接口请求参数
type NoteDetailReq struct {
	g.Meta `path:"/detail" tags:"note" method:"get" summary:"笔记详情接口"`
	model.NoteDetailReq
}

// 笔记详情接口响应参数
type NoteDetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoteDetailRes
}

// 笔记列表接口请求参数
type NoteListReq struct {
	g.Meta `path:"/list" tags:"note" method:"get" summary:"笔记列表接口"`
	model.NoteListReq
}

// 笔记列表请求接口响应参数
type NoteListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoteListRes
}

// 笔记删除接口请求参数
type NoteDeleteReq struct {
	g.Meta `path:"/delete" tags:"note" method:"post" summary:"笔记删除接口"`
	model.NoteDeleteReq
}

// 笔记删除接口响应参数
type NoteDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 恢复笔记接口请求参数
type NoteRecoverReq struct {
	g.Meta `path:"/recover" tags:"note" method:"post" summary:"恢复笔记接口"`
	model.NoteRecoverReq
}

// 恢复笔记接口响应参数
type NoteRecoverRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 笔记移动分类接口请求参数
type NoteMoveReq struct {
	g.Meta `path:"/move" tags:"note" method:"post" summary:"笔记移动分类接口"`
	model.NoteMoveReq
}

// 笔记移动分类接口响应参数
type NoteMoveRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 标记笔记接口请求参数
type NoteAsteriskReq struct {
	g.Meta `path:"/asterisk" tags:"note" method:"post" summary:"标记笔记接口"`
	model.NoteAsteriskReq
}

// 标记笔记接口响应参数
type NoteAsteriskRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 笔记标签接口请求参数
type NoteTagsReq struct {
	g.Meta `path:"/tag" tags:"note" method:"post" summary:"笔记标签接口"`
	model.NoteTagsReq
}

// 笔记标签接口响应参数
type NoteTagsRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 永久删除笔记接口请求参数
type NoteForeverDeleteReq struct {
	g.Meta `path:"/forever/delete" tags:"note" method:"post" summary:"永久删除笔记接口"`
	model.NoteForeverDeleteReq
}

// 永久删除笔记接口响应参数
type NoteForeverDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 笔记图片上传接口请求参数
type NoteUploadImageReq struct {
	g.Meta `path:"/upload/image" tags:"note" method:"post" summary:"笔记图片上传接口"`
}

// 笔记图片上传接口响应参数
type NoteUploadImageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoteUploadImageRes
}
