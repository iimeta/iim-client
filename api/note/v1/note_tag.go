package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 笔记标签列表接口请求参数
type TagListReq struct {
	g.Meta `path:"/tag/list" tags:"note_tag" method:"get" summary:"笔记标签列表接口"`
}

// 笔记标签列表接口响应参数
type TagListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TagListRes
}

// 笔记标签编辑接口请求参数
type TagEditReq struct {
	g.Meta `path:"/tag/editor" tags:"note_tag" method:"post" summary:"笔记标签编辑接口"`
	model.TagEditReq
}

// 笔记标签编辑接口请求参数
type TagEditRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TagEditRes
}

// 笔记标签删除接口请求参数
type TagDeleteReq struct {
	g.Meta `path:"/tag/delete" tags:"note_tag" method:"post" summary:"笔记标签删除接口"`
	model.TagDeleteReq
}

// 笔记标签删除接口响应参数
type TagDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
