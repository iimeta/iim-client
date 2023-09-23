package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 提交入群申请接口请求参数
type GroupApplyCreateReq struct {
	g.Meta `path:"/apply/create" tags:"group" method:"post" summary:"提交入群申请接口"`
	model.GroupApplyCreateReq
}

// 提交入群申请接口响应参数
type GroupApplyCreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 拒绝入群申请接口请求参数
type GroupApplyDeleteReq struct {
	g.Meta  `path:"/apply/delete" tags:"group" method:"post" summary:"拒绝入群申请接口"`
	ApplyId string `json:"apply_id,omitempty" v:"required"`
}

// 同意入群申请接口响应参数
type GroupApplyDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 同意入群申请接口请求参数
type GroupApplyAgreeReq struct {
	g.Meta `path:"/apply/agree" tags:"group" method:"post" summary:"同意入群申请接口"`
	model.GroupApplyAgreeReq
}

// 同意入群申请接口响应参数
type GroupApplyAgreeRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 拒绝入群申请接口请求参数
type GroupApplyDeclineReq struct {
	g.Meta `path:"/apply/decline" tags:"group" method:"post" summary:"拒绝入群申请接口"`
	model.GroupApplyDeclineReq
}

// 拒绝入群申请接口响应参数
type GroupApplyDeclineRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 入群申请列表接口请求参数
type GroupApplyListReq struct {
	g.Meta `path:"/apply/list" tags:"group" method:"get" summary:"入群申请列表接口"`
	model.GroupApplyListReq
}

// 入群申请列表接口响应参数
type GroupApplyListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupApplyListRes
}

// 申请管理-所有入群申请列表接口请求参数
type GroupApplyAllReq struct {
	g.Meta `path:"/apply/all" tags:"group" method:"get" summary:"申请管理-所有入群申请列表接口"`
}

// 申请管理-所有入群申请列表接口响应参数
type GroupApplyAllRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupApplyAllRes
}

// 入群申请未读接口请求参数
type GroupApplyUnreadReq struct {
	g.Meta `path:"/apply/unread" tags:"group" method:"get" summary:"入群申请未读接口"`
}

// 入群申请未读接口响应参数
type GroupApplyUnreadRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ApplyUnreadNumRes
}
