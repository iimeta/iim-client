package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 会话创建接口请求参数
type SessionCreateReq struct {
	g.Meta `path:"/create" tags:"talk_session" method:"post" summary:"会话创建接口"`
	model.SessionCreateReq
}

// 会话创建接口响应参数
type SessionCreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SessionCreateRes
}

// 会话删除接口请求参数
type SessionDeleteReq struct {
	g.Meta `path:"/delete" tags:"talk_session" method:"post" summary:"会话删除接口"`
	model.SessionDeleteReq
}

// 会话删除接口响应参数
type SessionDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 会话置顶接口请求参数
type SessionTopReq struct {
	g.Meta `path:"/topping" tags:"talk_session" method:"post" summary:"会话置顶接口"`
	model.SessionTopReq
}

// 会话置顶接口响应参数
type SessionTopRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 会话免打扰接口请求参数
type SessionDisturbReq struct {
	g.Meta `path:"/disturb" tags:"talk_session" method:"post" summary:"会话免打扰接口"`
	model.SessionDisturbReq
}

// 会话免打扰接口响应参数
type SessionDisturbRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 会话列表接口请求参数
type SessionListReq struct {
	g.Meta `path:"/list" tags:"talk_session" method:"get" summary:"会话列表接口"`
}

// 会话列表接口响应参数
type SessionListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SessionListRes
}

// 会话未读数清除接口请求参数
type SessionClearUnreadNumReq struct {
	g.Meta `path:"/unread/clear" tags:"talk_session" method:"post" summary:"会话未读数清除接口"`
	model.SessionClearUnreadNumReq
}

// 会话未读数清除接口响应参数
type SessionClearUnreadNumRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 清空上下文接口请求参数
type SessionClearContextReq struct {
	g.Meta `path:"/clear/context" tags:"talk_session" method:"post" summary:"清空上下文接口"`
	model.SessionClearContextReq
}

// 清空上下文接口响应参数
type SessionClearContextRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 开启/关闭上下文接口请求参数
type SessionOpenContextReq struct {
	g.Meta `path:"/open/context" tags:"talk_session" method:"post" summary:"开启/关闭上下文接口"`
	model.SessionOpenContextReq
}

// 开启/关闭上下文接口响应参数
type SessionOpenContextRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
