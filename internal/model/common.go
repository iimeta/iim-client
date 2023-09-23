package model

// 发送短信验证码接口请求参数
type CommonSendSmsReq struct {
	Mobile  string `json:"mobile,omitempty" v:"required|length:0,11"`
	Channel string `json:"channel,omitempty" v:"required|in:login,register,forget_account,change_email,change_mobile"`
}

// 发送短信验证码接口响应参数
type CommonSendSmsRes struct {
	IsDebug bool   `json:"is_debug"`
	SmsCode string `json:"code"`
}

// 发送邮件验证码接口请求参数
type CommonSendEmailReq struct {
	Email   string `json:"email,omitempty" v:"required"`
	Channel string `json:"channel,omitempty" v:"required|in:login,register,forget_account,change_email,change_mobile"`
}

// 发送邮件验证码接口响应参数
type CommonSendEmailRes struct {
	IsDebug bool   `json:"is_debug"`
	Code    string `json:"code"`
}
