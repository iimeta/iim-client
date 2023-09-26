package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 添加好友申请接口请求参数
type ApplyCreateReq struct {
	g.Meta `path:"/apply/create" tags:"contact_apply" method:"post" summary:"添加好友申请接口"`
	model.ApplyCreateReq
}

// 添加好友申请接口响应参数
type ApplyCreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 同意好友申请接口请求参数
type ApplyAcceptReq struct {
	g.Meta `path:"/apply/accept" tags:"contact_apply" method:"post" summary:"同意好友申请接口"`
	model.ApplyAcceptReq
}

// 同意好友申请接口响应参数
type ApplyAcceptRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 拒绝好友申请接口请求参数
type ApplyDeclineReq struct {
	g.Meta `path:"/apply/decline" tags:"contact_apply" method:"post" summary:"拒绝好友申请接口"`
	model.ApplyDeclineReq
}

// 拒绝好友申请接口响应参数
type ApplyDeclineRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 好友申请列表接口请求参数
type ApplyListReq struct {
	g.Meta `path:"/apply/records" tags:"contact_apply" method:"get" summary:"好友申请列表接口"`
}

// 好友申请列表接口响应参数
type ApplyListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ApplyListRes
}

// 好友申请未读数接口请求参数
type ApplyUnreadNumReq struct {
	g.Meta `path:"/apply/unread-num" tags:"contact_apply" method:"get" summary:"好友申请未读数接口"`
}

// 好友申请未读数接口响应参数
type ApplyUnreadNumRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ApplyUnreadNumRes
}
