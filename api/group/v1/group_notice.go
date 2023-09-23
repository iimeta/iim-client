package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 删除群公告接口请求参数
type GroupNoticeDeleteReq struct {
	g.Meta `path:"/notice/delete" tags:"group" method:"post" summary:"删除群公告接口"`
	model.GroupNoticeDeleteReq
}

// 删除群公告接口响应参数
type GroupNoticeDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 添加或编辑群公告接口请求参数
type GroupNoticeEditReq struct {
	g.Meta `path:"/notice/edit" tags:"group" method:"post" summary:"添加或编辑群公告接口"`
	model.GroupNoticeEditReq
}

// 添加或编辑群公告接口响应参数
type GroupNoticeEditRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 群公告列表接口请求参数
type GroupNoticeListReq struct {
	g.Meta `path:"/notice/list" tags:"group" method:"get" summary:"群公告列表接口"`
	model.GroupNoticeListReq
}

// 群公告列表接口响应参数
type GroupNoticeListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupNoticeListRes
}
