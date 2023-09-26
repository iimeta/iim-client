package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 会话创建接口请求参数
type TalkSessionCreateReq struct {
	g.Meta `path:"/create" tags:"talk" method:"post" summary:"会话创建接口"`
	model.TalkSessionCreateReq
}

// 会话创建接口响应参数
type TalkSessionCreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TalkSessionCreateRes
}

// 会话删除接口请求参数
type TalkSessionDeleteReq struct {
	g.Meta `path:"/delete" tags:"talk" method:"post" summary:"会话删除接口"`
	model.TalkSessionDeleteReq
}

// 会话删除接口响应参数
type TalkSessionDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 会话置顶接口请求参数
type TalkSessionTopReq struct {
	g.Meta `path:"/topping" tags:"talk" method:"post" summary:"会话置顶接口"`
	model.TalkSessionTopReq
}

// 会话置顶接口响应参数
type TalkSessionTopRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 会话免打扰接口请求参数
type TalkSessionDisturbReq struct {
	g.Meta `path:"/disturb" tags:"talk" method:"post" summary:"会话免打扰接口"`
	model.TalkSessionDisturbReq
}

// 会话免打扰接口响应参数
type TalkSessionDisturbRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 会话列表接口请求参数
type TalkSessionListReq struct {
	g.Meta `path:"/list" tags:"talk" method:"get" summary:"会话列表接口"`
}

// 会话列表接口响应参数
type TalkSessionListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.TalkSessionListRes
}

// 会话未读数清除接口请求参数
type TalkSessionClearUnreadNumReq struct {
	g.Meta `path:"/unread/clear" tags:"talk" method:"post" summary:"会话未读数清除接口"`
	model.TalkSessionClearUnreadNumReq
}

// 会话未读数清除接口响应参数
type TalkSessionClearUnreadNumRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 清空上下文接口请求参数
type TalkClearContextReq struct {
	g.Meta `path:"/clear/context" tags:"talk" method:"post" summary:"清空上下文接口"`
	model.TalkClearContextReq
}

// 清空上下文接口响应参数
type TalkClearContextRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 开启/关闭上下文接口请求参数
type TalkOpenContextReq struct {
	g.Meta `path:"/open/context" tags:"talk" method:"post" summary:"开启/关闭上下文接口"`
	model.TalkOpenContextReq
}

// 开启/关闭上下文接口响应参数
type TalkOpenContextRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
