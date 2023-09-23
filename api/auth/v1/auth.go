package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 登录接口请求参数
type AuthLoginReq struct {
	g.Meta `path:"/login" tags:"auth" method:"post" summary:"登录接口"`
	model.AuthLoginReq
}

// 登录接口响应参数
type AuthLoginRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AuthLoginRes
}

// 登出接口请求参数
type AuthLogoutReq struct {
	g.Meta `path:"/logout" tags:"auth" method:"post" summary:"登出接口"`
}

// 登出接口响应参数
type AuthLogoutRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 注册接口请求参数
type AuthRegisterReq struct {
	g.Meta `path:"/register" tags:"auth" method:"post" summary:"注册接口"`
	model.AuthRegisterReq
}

// 注册接口响应参数
type AuthRegisterRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// Token 刷新接口请求参数
type AuthRefreshReq struct {
	g.Meta `path:"/refresh" tags:"auth" method:"post" summary:"刷新Token接口"`
}

// Token 刷新接口响应参数
type AuthRefreshRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.AuthRefreshRes
}

// 找回密码接口请求参数
type AuthForgetReq struct {
	g.Meta `path:"/forget" tags:"auth" method:"post" summary:"找回密码接口"`
	model.AuthForgetReq
}

// 找回密码接口响应参数
type AuthForgetRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
