package model

type Message struct {
	MsgType  string    `json:"msg_type"`
	Sender   *Sender   `json:"sender,omitempty"`    // 消息发送者
	Receiver *Receiver `json:"receiver,omitempty"`  // 消息接收者
	TalkType int       `json:"talk_type,omitempty"` // 对话类型 1:私聊 2:群聊
	Mention  *Mention  `json:"mention,omitempty"`
	QuoteId  string    `json:"quote_id,omitempty"` // 引用的消息ID

	Text     *Text     `json:"text,omitempty"`
	Image    *Image    `json:"image,omitempty"`
	Voice    *Voice    `json:"voice,omitempty"`
	Video    *Video    `json:"video,omitempty"`
	File     *File     `json:"file,omitempty"`
	Vote     *Vote     `json:"vote,omitempty"`
	Mixed    *Mixed    `json:"mixed,omitempty"`
	Card     *Card     `json:"card,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type SysMessage struct {
	MsgType  string    `json:"msg_type"`
	Sender   *Sender   `json:"sender,omitempty"`   // 消息发送者
	Receiver *Receiver `json:"receiver,omitempty"` // 消息接收者

	Text *Text
}

type NoticeMessage struct {
	MsgType string `json:"msg_type"`
	Login   *Login `json:"login,omitempty"`
}

type Sender struct {
	SenderId int `json:"sender_id,omitempty"` // 发送者ID
}

type Receiver struct {
	TalkType   int `json:"talk_type,omitempty"`   // 对话类型 1:私聊 2:群聊 todo
	ReceiverId int `json:"receiver_id,omitempty"` // 接收者ID, 好友ID或群ID
}

type Mention struct {
	All  int   `json:"all,omitempty"`
	Uids []int `json:"uids,omitempty"`
}

// 登录消息
type Login struct {
	Ip       string `json:"ip,omitempty"`
	Address  string `json:"address,omitempty"`
	Platform string `json:"platform,omitempty"`
	Agent    string `json:"agent,omitempty"`
	Reason   string `json:"reason,omitempty"`
}

// 表情消息
type Emoticon struct {
	EmoticonId string `json:"emoticon_id" v:"required"`
}

// 位置消息
type Card struct {
	UserId   int `json:"user_id,omitempty" v:"required"`
	TalkType int `json:"talk_type" v:"required|in:1,2"`
}

// 图文消息
type Mixed struct {
	Items []*MixedItem `json:"items"`
}

type MixedItem struct {
	MsgType string `json:"msg_type"`
	Text    *Text  `json:"text,omitempty"`
	Image   *Image `json:"image,omitempty"`
}

// 代码消息
type Code struct {
	Lang string `json:"lang" v:"required"`
	Code string `json:"code" v:"required"`
}

// 投票消息接口请求参数
type Vote struct {
	Mode      int      `json:"mode" v:"in:0,1"`
	Anonymous int      `json:"anonymous" v:"in:0,1"`
	Title     string   `json:"title" v:"required"`
	Options   []string `json:"options"`
}

// 位置消息
type Location struct {
	Longitude   string `json:"longitude,omitempty" v:"required"`   // 地理位置 经度
	Latitude    string `json:"latitude,omitempty" v:"required"`    // 地理位置 纬度
	Description string `json:"description,omitempty" v:"required"` // 位置描述
}

// 转发消息
type Forward struct {
	Mode            int    `json:"mode,omitempty" v:"required"`        // 转发模式
	MessageIds      []int  `json:"message_ids,omitempty" v:"required"` // 消息ID
	Gids            []int  `json:"gids,omitempty"`                     // 群ID列表
	Uids            []int  `json:"uids,omitempty"`                     // 好友ID列表
	MsgType         int    `json:"msg_type"`                           // 消息类型
	RecordId        int    `json:"record_id" v:"min:0"`                // 上次查询的最小消息ID
	Limit           int    `json:"limit" v:"required|max:100"`         // 数据行数
	ForwardMode     int    `json:"forward_mode" v:"required|in:1,2"`
	RecordsIds      string `json:"records_ids" v:"required"`
	ReceiveUserIds  string `json:"receive_user_ids"`
	ReceiveGroupIds string `json:"receive_group_ids"`
}

// 文本消息
type Text struct {
	Content string `json:"content,omitempty" v:"required"`
}

// 图片消息
type Image struct {
	Url    string `json:"url,omitempty" v:"required"`    // 图片地址
	Width  int    `json:"width,omitempty" v:"required"`  // 图片宽度
	Height int    `json:"height,omitempty" v:"required"` // 图片高度
	Size   int    `json:"size,omitempty" v:"required"`   // 图片大小
}

// 语音消息
type Voice struct {
	Url      string `json:"url,omitempty" v:"required"`
	Duration int    `json:"duration,omitempty" v:"required"`
	Size     int    `json:"size,omitempty" v:"required"` // 语音大小
}

// 视频文件消息
type Video struct {
	Url      string `json:"url,omitempty" v:"required"`
	Duration int    `json:"duration,omitempty" v:"required"`
	Size     int    `json:"size,omitempty" v:"required"` // 视频大小
	Cover    string `json:"cover,omitempty"`             // 封面图
}

// 文件消息
type File struct {
	UploadId string `json:"upload_id" v:"required"`
	Name     string `json:"name"`   // 文件名称
	Drive    int    `json:"drive"`  // 文件存储方式
	Suffix   string `json:"suffix"` // 文件后缀
	Size     int    `json:"size"`   // 文件大小
	Path     string `json:"path"`   // 文件路径
}
