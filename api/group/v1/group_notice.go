package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 删除群公告接口请求参数
type NoticeDeleteReq struct {
	g.Meta `path:"/notice/delete" tags:"group_notice" method:"post" summary:"删除群公告接口"`
	model.NoticeDeleteReq
}

// 删除群公告接口响应参数
type NoticeDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 添加或编辑群公告接口请求参数
type NoticeEditReq struct {
	g.Meta `path:"/notice/edit" tags:"group_notice" method:"post" summary:"添加或编辑群公告接口"`
	model.NoticeEditReq
}

// 添加或编辑群公告接口响应参数
type NoticeEditRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 群公告列表接口请求参数
type NoticeListReq struct {
	g.Meta `path:"/notice/list" tags:"group_notice" method:"get" summary:"群公告列表接口"`
	model.NoticeListReq
}

// 群公告列表接口响应参数
type NoticeListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.NoticeListRes
}
