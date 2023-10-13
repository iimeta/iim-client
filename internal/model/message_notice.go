package model

type NoticeMessage struct {
	MsgId    string    `json:"msg_id,omitempty" bson:"msg_id,omitempty"`
	MsgType  string    `json:"msg_type,omitempty" bson:"msg_type,omitempty"`
	TalkType int       `json:"talk_type,omitempty" bson:"talk_type,omitempty"` // 对话类型 1:私聊 2:群聊
	Sender   *Sender   `json:"sender,omitempty" bson:"sender,omitempty"`       // 发送者
	Receiver *Receiver `json:"receiver,omitempty" bson:"receiver,omitempty"`   // 接收者

	Login *Login `json:"login,omitempty" bson:"login,omitempty"`
}

// 登录消息
type Login struct {
	IP       string `json:"ip,omitempty" bson:"ip,omitempty"`             // 登录IP
	Address  string `json:"address,omitempty" bson:"address,omitempty"`   // 登录地址
	Platform string `json:"platform,omitempty" bson:"platform,omitempty"` // 登录平台
	Agent    string `json:"agent,omitempty" bson:"agent,omitempty"`       // 登录设备
	Reason   string `json:"reason,omitempty" bson:"reason,omitempty"`     // 登录原因
	Datetime string `json:"datetime,omitempty" bson:"datetime,omitempty"` // 登录时间
}
