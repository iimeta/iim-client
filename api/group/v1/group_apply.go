package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 提交入群申请接口请求参数
type ApplyCreateReq struct {
	g.Meta `path:"/apply/create" tags:"group_apply" method:"post" summary:"提交入群申请接口"`
	model.GroupApplyCreateReq
}

// 提交入群申请接口响应参数
type ApplyCreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 拒绝入群申请接口请求参数
type ApplyDeleteReq struct {
	g.Meta `path:"/apply/delete" tags:"group_apply" method:"post" summary:"拒绝入群申请接口"`
	model.ApplyDeleteReq
}

// 同意入群申请接口响应参数
type ApplyDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 同意入群申请接口请求参数
type ApplyAgreeReq struct {
	g.Meta `path:"/apply/agree" tags:"group_apply" method:"post" summary:"同意入群申请接口"`
	model.ApplyAgreeReq
}

// 同意入群申请接口响应参数
type ApplyAgreeRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 拒绝入群申请接口请求参数
type ApplyDeclineReq struct {
	g.Meta `path:"/apply/decline" tags:"group_apply" method:"post" summary:"拒绝入群申请接口"`
	model.GroupApplyDeclineReq
}

// 拒绝入群申请接口响应参数
type ApplyDeclineRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 入群申请列表接口请求参数
type ApplyListReq struct {
	g.Meta `path:"/apply/list" tags:"group_apply" method:"get" summary:"入群申请列表接口"`
	model.ApplyListReq
}

// 入群申请列表接口响应参数
type ApplyListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupApplyListRes
}

// 申请管理-所有入群申请列表接口请求参数
type ApplyAllReq struct {
	g.Meta `path:"/apply/all" tags:"group_apply" method:"get" summary:"申请管理-所有入群申请列表接口"`
}

// 申请管理-所有入群申请列表接口响应参数
type ApplyAllRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ApplyAllRes
}

// 入群申请未读接口请求参数
type ApplyUnreadReq struct {
	g.Meta `path:"/apply/unread" tags:"group_apply" method:"get" summary:"入群申请未读接口"`
}

// 入群申请未读接口响应参数
type ApplyUnreadRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupApplyUnreadNumRes
}
