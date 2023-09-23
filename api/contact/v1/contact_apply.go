package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 添加联系人申请接口请求参数
type ContactApplyCreateReq struct {
	g.Meta `path:"/apply/create" tags:"contact" method:"post" summary:"添加联系人申请接口"`
	model.ContactApplyCreateReq
}

// 添加联系人申请接口响应参数
type ContactApplyCreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 同意联系人申请接口请求参数
type ContactApplyAcceptReq struct {
	g.Meta `path:"/apply/accept" tags:"contact" method:"post" summary:"同意联系人申请接口"`
	model.ContactApplyAcceptReq
}

// 同意联系人申请接口响应参数
type ContactApplyAcceptRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 拒绝联系人申请接口请求参数
type ContactApplyDeclineReq struct {
	g.Meta `path:"/apply/decline" tags:"contact" method:"post" summary:"拒绝联系人申请接口"`
	model.ContactApplyDeclineReq
}

// 拒绝联系人申请接口响应参数
type ContactApplyDeclineRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 联系人申请列表接口请求参数
type ContactApplyListReq struct {
	g.Meta `path:"/apply/records" tags:"contact" method:"get" summary:"联系人申请列表接口"`
}

// 联系人申请列表接口响应参数
type ContactApplyListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.ContactApplyListRes
}

// 联系人申请未读数接口请求参数
type ContactApplyUnreadNumReq struct {
	g.Meta `path:"/apply/unread-num" tags:"contact" method:"get" summary:"联系人申请未读数接口"`
}

// 联系人申请未读数接口响应参数
type ContactApplyUnreadNumRes struct {
	g.Meta    `mime:"application/json" example:"json"`
	UnreadNum int `json:"unread_num"`
}
