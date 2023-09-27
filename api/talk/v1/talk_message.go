package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 发送消息接口请求参数
type MessagePublishReq struct {
	g.Meta `path:"/message/publish" tags:"talk_message" method:"post" summary:"发送消息接口"`
	model.MessagePublishReq
}

// 发送消息接口响应参数
type MessagePublishRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 发送文件消息接口请求参数
type MessageFileReq struct {
	g.Meta `path:"/message/file" tags:"talk_message" method:"post" summary:"发送文件消息接口"`
	model.MessageFileReq
}

// 发送文件消息接口响应参数
type MessageFileRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 投票消息接口请求参数
type MessageVoteReq struct {
	g.Meta `path:"/message/vote" tags:"talk_message" method:"post" summary:"投票消息接口"`
	model.MessageVoteReq
}

// 投票消息接口响应参数
type MessageVoteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 投票消息处理接口请求参数
type MessageVoteHandleReq struct {
	g.Meta `path:"/message/vote/handle" tags:"talk_message" method:"post" summary:"投票消息处理接口"`
	model.MessageVoteHandleReq
}

// 投票消息接口响应参数
type MessageVoteHandleRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.VoteStatistics
}

// 收藏会话表情图片接口请求参数
type MessageCollectReq struct {
	g.Meta `path:"/message/collect" tags:"talk_message" method:"post" summary:"收藏会话表情图片接口"`
	model.MessageCollectReq
}

// 收藏会话表情图片接口响应参数
type MessageCollectRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 撤销聊天消息接口请求参数
type MessageRevokeReq struct {
	g.Meta `path:"/message/revoke" tags:"talk_message" method:"post" summary:"撤销聊天消息接口"`
	model.MessageRevokeReq
}

// 撤销聊天消息接口响应参数
type MessageRevokeRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 删除聊天记录接口请求参数
type MessageDeleteReq struct {
	g.Meta `path:"/message/delete" tags:"talk_message" method:"post" summary:"删除聊天记录接口"`
	model.MessageDeleteReq
}

// 删除聊天记录接口响应参数
type MessageDeleteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
