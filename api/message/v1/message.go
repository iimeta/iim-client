package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
)

// 发送文件消息接口请求参数
type FileMessageReq struct {
	g.Meta `path:"/file" tags:"message" method:"post" summary:"发送文件消息接口"`
	model.FileMessageReq
}

// 发送文件消息接口响应参数
type FileMessageRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 投票消息接口请求参数
type VoteMessageReq struct {
	g.Meta `path:"/vote" tags:"message" method:"post" summary:"投票消息接口"`
	model.VoteMessageReq
}

// 投票消息接口响应参数
type VoteMessageRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 投票消息处理接口请求参数
type VoteMessageHandleReq struct {
	g.Meta `path:"/vote/handle" tags:"message" method:"post" summary:"投票消息处理接口"`
	model.VoteMessageHandleReq
}

// 投票消息接口响应参数
type VoteMessageHandleRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*dao.VoteStatistics
}

// 发送消息接口请求参数
type PublishBaseMessageReq struct {
	g.Meta `path:"/publish" tags:"message" method:"post" summary:"发送消息接口"`
	model.PublishBaseMessageReq
}

// 发送消息接口响应参数
type PublishBaseMessageRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 收藏会话表情图片接口请求参数
type CollectMessageReq struct {
	g.Meta `path:"/collect" tags:"message" method:"post" summary:"收藏会话表情图片接口"`
	model.CollectMessageReq
}

// 收藏会话表情图片接口响应参数
type CollectMessageRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 撤销聊天消息接口请求参数
type RevokeMessageReq struct {
	g.Meta `path:"/revoke" tags:"message" method:"post" summary:"撤销聊天消息接口"`
	model.RevokeMessageReq
}

// 撤销聊天消息接口响应参数
type RevokeMessageRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除聊天记录接口请求参数
type DeleteMessageReq struct {
	g.Meta `path:"/delete" tags:"message" method:"post" summary:"删除聊天记录接口"`
	model.DeleteMessageReq
}

// 删除聊天记录接口响应参数
type DeleteMessageRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
