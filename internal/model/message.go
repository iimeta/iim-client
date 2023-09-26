package model

// 登录消息
type LoginMessageReq struct {
	Ip       string `json:"ip,omitempty"`
	Address  string `json:"address,omitempty"`
	Platform string `json:"platform,omitempty"`
	Agent    string `json:"agent,omitempty"`
	Reason   string `json:"reason,omitempty"`
}

// 表情消息
type EmoticonMessageReq struct {
	Type       string           `json:"type,omitempty"`
	Receiver   *MessageReceiver `json:"receiver,omitempty"` // 消息接收者
	TalkType   int              `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int              `json:"receiver_id" v:"required"`
	EmoticonId string           `json:"emoticon_id" v:"required"`
}

// 位置消息
type CardMessageReq struct {
	Type       string           `json:"type,omitempty"`
	UserId     int              `json:"user_id,omitempty" v:"required"`
	Receiver   *MessageReceiver `json:"receiver,omitempty"`
	TalkType   int              `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int              `json:"receiver_id" v:"required"`
}

// 图文消息
type MixedMessageReq struct {
	Type     string                      `json:"type,omitempty"`
	Items    []*MixedMessageRequest_Item `json:"items"`
	Receiver *MessageReceiver            `json:"receiver,omitempty"`
	QuoteId  string                      `json:"quote_id,omitempty"` // 引用的消息ID
}

type MixedMessageRequest_Item struct {
	Type    int    `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
}

// 发送文件消息接口请求参数
type FileMessageReq struct {
	Type       string           `json:"type,omitempty"`
	Receiver   *MessageReceiver `json:"receiver,omitempty"` // 消息接收者
	TalkType   int              `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int              `json:"receiver_id" v:"required"`
	UploadId   string           `json:"upload_id" v:"required"`
}

// 代码消息
type CodeMessageReq struct {
	Type       string           `json:"type,omitempty"`
	Receiver   *MessageReceiver `json:"receiver,omitempty"` // 消息接收者
	TalkType   int              `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int              `json:"receiver_id" v:"required"`
	Lang       string           `json:"lang" v:"required"`
	Code       string           `json:"code" v:"required|max-length:65535"`
}

// 投票消息接口请求参数
type VoteMessageReq struct {
	Type       string           `json:"type,omitempty"`
	Receiver   *MessageReceiver `json:"receiver,omitempty"` // 消息接收者
	ReceiverId int              `json:"receiver_id" v:"required"`
	Mode       int              `json:"mode" v:"in:0,1"`
	Anonymous  int              `json:"anonymous" v:"in:0,1"`
	Title      string           `json:"title" v:"required"`
	Options    []string         `json:"options"`
}

// 位置消息
type LocationMessageReq struct {
	Type        string           `json:"type,omitempty"`
	Longitude   string           `json:"longitude,omitempty" v:"required"`   // 地理位置 经度
	Latitude    string           `json:"latitude,omitempty" v:"required"`    // 地理位置 纬度
	Description string           `json:"description,omitempty" v:"required"` // 位置描述
	Receiver    *MessageReceiver `json:"receiver,omitempty"`                 // 消息接收者
	TalkType    int              `json:"talk_type" v:"required|in:1,2"`
	ReceiverId  int              `json:"receiver_id" v:"required"`
}

// // 转发消息
type ForwardMessageReq struct {
	Type            string           `json:"type,omitempty"`
	Mode            int              `json:"mode,omitempty" v:"required"`        // 转发模式
	MessageIds      []int            `json:"message_ids,omitempty" v:"required"` // 消息ID
	Gids            []int            `json:"gids,omitempty"`                     // 群ID列表
	Uids            []int            `json:"uids,omitempty"`                     // 好友ID列表
	Receiver        *MessageReceiver `json:"receiver,omitempty"`                 // 消息接收者
	TalkType        int              `json:"talk_type" v:"required|in:1,2"`      // 对话类型
	ReceiverId      int              `json:"receiver_id" v:"required"`           // 接收者ID
	MsgType         int              `json:"msg_type"`                           // 消息类型
	RecordId        int              `json:"record_id" v:"min:0"`                // 上次查询的最小消息ID
	Limit           int              `json:"limit" v:"required|max:100"`         // 数据行数
	ForwardMode     int              `json:"forward_mode" v:"required|in:1,2"`
	RecordsIds      string           `json:"records_ids" v:"required"`
	ReceiveUserIds  string           `json:"receive_user_ids"`
	ReceiveGroupIds string           `json:"receive_group_ids"`
}

type MessageReceiver struct {
	TalkType   int `json:"talk_type,omitempty"`   // 对话类型
	ReceiverId int `json:"receiver_id,omitempty"` // 接受者ID
}

// 文本消息
type TextMessageReq struct {
	Type       string                      `json:"type,omitempty"` // 消息类型
	Content    string                      `json:"content,omitempty" v:"required"`
	Mention    *TextMessageRequest_Mention `json:"mention,omitempty"`
	QuoteId    string                      `json:"quote_id,omitempty"` // 引用的消息ID
	Receiver   *MessageReceiver            `json:"receiver,omitempty"` // 消息接收者
	TalkType   int                         `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int                         `json:"receiver_id" v:"required"`
	Text       string                      `json:"text" v:"required|max-length:3000"`
}

type TextMessageRequest_Mention struct {
	All  int   `json:"all,omitempty"`
	Uids []int `json:"uids,omitempty"`
}

// 图片消息
type ImageMessageReq struct {
	Type       string           `json:"type,omitempty"`
	Url        string           `json:"url,omitempty" v:"required"`    // 图片地址
	Width      int              `json:"width,omitempty" v:"required"`  // 图片宽度
	Height     int              `json:"height,omitempty" v:"required"` // 图片高度
	Size       int              `json:"size,omitempty" v:"required"`   // 图片大小
	Receiver   *MessageReceiver `json:"receiver,omitempty"`            // 消息接收者
	QuoteId    string           `json:"quote_id,omitempty"`            // 引用的消息ID
	TalkType   int              `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int              `json:"receiver_id" v:"required"`
}

// 语音消息
type VoiceMessageReq struct {
	Type     string           `json:"type,omitempty"`
	Url      string           `json:"url,omitempty" v:"required"`
	Duration int              `json:"duration,omitempty" v:"required"`
	Size     int              `json:"size,omitempty" v:"required"` // 语音大小
	Receiver *MessageReceiver `json:"receiver,omitempty"`          // 消息接收者
}

// 视频文件消息
type VideoMessageReq struct {
	Type     string           `json:"type,omitempty"`
	Url      string           `json:"url,omitempty" v:"required"`
	Duration int              `json:"duration,omitempty" v:"required"`
	Size     int              `json:"size,omitempty" v:"required"` // 视频大小
	Receiver *MessageReceiver `json:"receiver,omitempty"`          // 消息接收者
	Cover    string           `json:"cover,omitempty"`             //封面图
}

type ForwardRecord struct {
	RecordId   int
	ReceiverId int
	TalkType   int
}

type TextMessage struct {
	AckId   string         `json:"ack_id"`
	Event   string         `json:"event"`
	Content TextMessageReq `json:"content"`
}

type CodeMessage struct {
	AckId   string         `json:"ack_id"`
	Event   string         `json:"event"`
	Content CodeMessageReq `json:"content"`
}

type EmoticonMessage struct {
	MsgId   string             `json:"msg_id"`
	Event   string             `json:"event"`
	Content EmoticonMessageReq `json:"content"`
}

type ImageMessage struct {
	MsgId   string          `json:"msg_id"`
	Event   string          `json:"event"`
	Content ImageMessageReq `json:"content"`
}

type FileMessage struct {
	MsgId   string          `json:"msg_id"`
	Event   string          `json:"event"`
	Content ImageMessageReq `json:"content"`
}

type LocationMessage struct {
	MsgId   string             `json:"msg_id"`
	Event   string             `json:"event"`
	Content LocationMessageReq `json:"content"`
}

type VoteMessage struct {
	MsgId   string         `json:"msg_id"`
	Event   string         `json:"event"`
	Content VoteMessageReq `json:"content"`
}

type KeyboardMessage struct {
	Event   string `json:"event"`
	Content struct {
		SenderID   int `json:"sender_id"`
		ReceiverID int `json:"receiver_id"`
	} `json:"content"`
}

type TalkReadMessage struct {
	Event   string `json:"event"`
	Content struct {
		MsgIds     []int `json:"msg_id"`
		ReceiverId int   `json:"receiver_id"`
	} `json:"content"`
}

type AuthorityOption struct {
	TalkType   int // 对话类型
	UserId     int // 发送者ID
	ReceiverId int // 接收者ID
}

type CollectMessageReq struct {
	RecordId int `json:"record_id" v:"required"`
}

type RevokeMessageReq struct {
	RecordId int `json:"record_id" v:"required"`
}

type DeleteMessageReq struct {
	TalkType   int    `json:"talk_type" v:"required|in:1,2"`
	ReceiverId int    `json:"receiver_id" v:"required"`
	RecordIds  string `json:"record_id" v:"record_id@required"`
}

type VoteMessageHandleReq struct {
	RecordId int    `json:"record_id" v:"required"`
	Options  string `json:"options" v:"required"`
}

type PublishBaseMessageReq struct {
	Type     string `json:"type" v:"required"`
	Receiver struct {
		TalkType   int `json:"talk_type" v:"required"`   // 对话类型 1:私聊 2:群聊
		ReceiverId int `json:"receiver_id" v:"required"` // 好友ID或群ID
	} `json:"receiver" v:"required"`
}
