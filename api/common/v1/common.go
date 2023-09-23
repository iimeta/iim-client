package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 发送短信验证码接口请求参数
type CommonSendSmsReq struct {
	g.Meta `path:"/sms-code" tags:"common" method:"post" summary:"发送短信验证码接口"`
	model.CommonSendSmsReq
}

// 发送短信验证码接口响应参数
type CommonSendSmsRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.CommonSendSmsRes
}

// 发送邮件验证码接口请求参数
type CommonSendEmailReq struct {
	g.Meta `path:"/email-code" tags:"common" method:"post" summary:"发送邮件验证码接口"`
	model.CommonSendEmailReq
}

// 发送邮件验证码接口响应参数
type CommonSendEmailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.CommonSendEmailRes
}
