package consts

// IM 渠道分组(用于业务划分, 业务间相互隔离)
const (
	// 默认分组
	ImChannelChat = "chat"
)

const (
	// 默认渠道消息订阅
	ImTopicChat        = "im:message:chat:all"
	ImTopicChatPrivate = "im:message:chat:%s"
)
