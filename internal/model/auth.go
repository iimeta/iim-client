package model

// 登录接口请求参数
type LoginReq struct {
	// 登录账号
	Account string `json:"account,omitempty" v:"required"`
	// 登录密码
	Password string `json:"password,omitempty" v:"required"`
	// 登录平台
	Platform string `json:"platform,omitempty" v:"required|in:web,h5,ios,windows,mac,file"`
}

// 登录接口响应参数
type LoginRes struct {
	// Token 类型
	Type string `json:"type,omitempty"`
	// token
	AccessToken string `json:"access_token,omitempty"`
	// 过期时间
	ExpiresIn int `json:"expires_in,omitempty"`
}

// 注册接口请求参数
type RegisterReq struct {
	// 登录账号
	Account string `json:"account,omitempty" v:"required"`
	// 登录密码
	Password string `json:"password,omitempty" v:"required|min-length:6"`
	// 昵称
	Nickname string `json:"nickname,omitempty" v:"required|length:2,30"`
	// 登录平台
	Platform string `json:"platform,omitempty" v:"required|in:web,h5,ios,windows,mac,file"`
	// 短信验证码
	Code string `json:"code,omitempty" v:"required"`
}

// 注册接口响应参数
type RegisterRes struct {
}

// Token 刷新接口响应参数
type RefreshRes struct {
	Type        string `json:"type,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
}

// 找回密码接口请求参数
type ForgetReq struct {
	*UserForget
}
