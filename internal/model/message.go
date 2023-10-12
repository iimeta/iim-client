package model

type Message struct {
	MsgType  string    `json:"msg_type,omitempty" bson:"msg_type,omitempty"`
	Sender   *Sender   `json:"sender,omitempty" bson:"sender,omitempty"`       // 发送者
	Receiver *Receiver `json:"receiver,omitempty" bson:"receiver,omitempty"`   // 接收者
	TalkType int       `json:"talk_type,omitempty" bson:"talk_type,omitempty"` // 对话类型 1:私聊 2:群聊
	Mention  *Mention  `json:"mention,omitempty" bson:"mention,omitempty"`
	QuoteId  string    `json:"quote_id,omitempty" bson:"quote_id,omitempty"` // 引用的消息ID
	Reply    *Reply    `json:"reply,omitempty" bson:"reply,omitempty"`

	Text     *Text     `json:"text,omitempty" bson:"text,omitempty"`
	Code     *Code     `json:"code,omitempty" bson:"code,omitempty"`
	Image    *Image    `json:"image,omitempty" bson:"image,omitempty"`
	Voice    *Voice    `json:"voice,omitempty" bson:"voice,omitempty"`
	Video    *Video    `json:"video,omitempty" bson:"video,omitempty"`
	File     *File     `json:"file,omitempty" bson:"file,omitempty"`
	Vote     *Vote     `json:"vote,omitempty" bson:"vote,omitempty"`
	Mixed    *Mixed    `json:"mixed,omitempty" bson:"mixed,omitempty"`
	Emoticon *Emoticon `json:"emoticon,omitempty" bson:"emoticon,omitempty"`
	Card     *Card     `json:"card,omitempty" bson:"card,omitempty"`
	Location *Location `json:"location,omitempty" bson:"location,omitempty"`
}

type SysMessage struct {
	MsgType  string    `json:"msg_type,omitempty" bson:"msg_type,omitempty"`
	Sender   *Sender   `json:"sender,omitempty" bson:"sender,omitempty"`     // 发送者
	Receiver *Receiver `json:"receiver,omitempty" bson:"receiver,omitempty"` // 接收者

	Text *Text `json:"text,omitempty" bson:"text,omitempty"`
}

type NoticeMessage struct {
	MsgType  string    `json:"msg_type,omitempty" bson:"msg_type,omitempty"`
	Sender   *Sender   `json:"sender,omitempty" bson:"sender,omitempty"`     // 发送者
	Receiver *Receiver `json:"receiver,omitempty" bson:"receiver,omitempty"` // 接收者

	Login *Login `json:"login,omitempty" bson:"login,omitempty"`
}

type Sender struct {
	Id   int    `json:"id,omitempty" bson:"id,omitempty"`     // 发送者ID
	Name string `json:"name,omitempty" bson:"name,omitempty"` // 发送者名称
}

type Receiver struct {
	TalkType   int    `json:"talk_type,omitempty" bson:"talk_type,omitempty"`     // 对话类型 1:私聊 2:群聊 todo
	ReceiverId int    `json:"receiver_id,omitempty" bson:"receiver_id,omitempty"` // 接收者ID, 好友ID或群ID
	Id         int    `json:"id,omitempty" bson:"id,omitempty"`                   // 接收者ID, 好友ID或群ID
	Name       string `json:"name,omitempty" bson:"name,omitempty"`               // 接收者名称, 好友名称或群名称
}

type Mention struct {
	Type int   `json:"type,omitempty" bson:"type,omitempty"` // 提及类型, 1: 所有人, 2: 指定人
	Uids []int `json:"uids,omitempty" bson:"uids,omitempty"`
}

// 登录消息
type Login struct {
	Ip       string `json:"ip,omitempty" bson:"ip,omitempty"`
	Address  string `json:"address,omitempty" bson:"address,omitempty"`
	Platform string `json:"platform,omitempty" bson:"platform,omitempty"`
	Agent    string `json:"agent,omitempty" bson:"agent,omitempty"`
	Reason   string `json:"reason,omitempty" bson:"reason,omitempty"`
}

// 表情消息
type Emoticon struct {
	EmoticonId string `json:"emoticon_id,omitempty" bson:"emoticon_id,omitempty" v:"required"`
	Url        string `json:"url,omitempty" bson:"url,omitempty" v:"required"`       // 图片地址
	Width      int    `json:"width,omitempty" bson:"width,omitempty" v:"required"`   // 图片宽度
	Height     int    `json:"height,omitempty" bson:"height,omitempty" v:"required"` // 图片高度
	Size       int    `json:"size,omitempty" bson:"size,omitempty" v:"required"`     // 图片大小
}

// 位置消息
type Card struct {
	UserId int `json:"user_id,omitempty" bson:"user_id,omitempty" v:"required"`
}

// 图文消息
type Mixed struct {
	Items []*MixedItem `json:"items,omitempty" bson:"items,omitempty"`
}

type MixedItem struct {
	MsgType string `json:"msg_type,omitempty" bson:"msg_type,omitempty"`
	Text    *Text  `json:"text,omitempty" bson:"text,omitempty"`
	Image   *Image `json:"image,omitempty" bson:"image,omitempty"`
}

// 代码消息
type Code struct {
	Lang string `json:"lang,omitempty" bson:"lang,omitempty" v:"required"`
	Code string `json:"code,omitempty" bson:"code,omitempty" v:"required"`
}

// 投票消息接口请求参数
type Vote struct {
	Mode      int      `json:"mode,omitempty" bson:"mode,omitempty" v:"in:0,1"`
	Anonymous int      `json:"anonymous,omitempty" bson:"anonymous,omitempty" v:"in:0,1"`
	Title     string   `json:"title,omitempty" bson:"title,omitempty" v:"required"`
	Options   []string `json:"options,omitempty" bson:"options,omitempty"`
}

// 位置消息
type Location struct {
	Longitude   string `json:"longitude,omitempty" bson:"longitude,omitempty" v:"required"`     // 地理位置 经度
	Latitude    string `json:"latitude,omitempty" bson:"latitude,omitempty" v:"required"`       // 地理位置 纬度
	Description string `json:"description,omitempty" bson:"description,omitempty" v:"required"` // 位置描述
}

// 转发消息
type Forward struct {
	MsgType         string `json:"msg_type,omitempty" bson:"msg_type,omitempty"`
	Mode            int    `json:"mode,omitempty" bson:"mode,omitempty" v:"required"`               // 转发模式
	MessageIds      []int  `json:"message_ids,omitempty" bson:"message_ids,omitempty" v:"required"` // 消息ID
	Gids            []int  `json:"gids,omitempty" bson:"gids,omitempty"`                            // 群ID列表
	Uids            []int  `json:"uids,omitempty" bson:"uids,omitempty"`                            // 好友ID列表
	RecordId        int    `json:"record_id,omitempty" bson:"record_id,omitempty" v:"min:0"`        // 上次查询的最小消息ID
	Limit           int    `json:"limit,omitempty" bson:"limit,omitempty" v:"required|max:100"`     // 数据行数
	ForwardMode     int    `json:"forward_mode,omitempty" bson:"forward_mode,omitempty" v:"required|in:1,2"`
	RecordsIds      string `json:"records_ids,omitempty" bson:"records_ids,omitempty" v:"required"`
	ReceiveUserIds  string `json:"receive_user_ids,omitempty" bson:"receive_user_ids,omitempty"`
	ReceiveGroupIds string `json:"receive_group_ids,omitempty" bson:"receive_group_ids,omitempty"`
}

// 文本消息
type Text struct {
	Content string `json:"content,omitempty" bson:"content,omitempty" v:"required"`
}

// 图片消息
type Image struct {
	Url    string `json:"url,omitempty" bson:"url,omitempty" v:"required"`       // 图片地址
	Width  int    `json:"width,omitempty" bson:"width,omitempty" v:"required"`   // 图片宽度
	Height int    `json:"height,omitempty" bson:"height,omitempty" v:"required"` // 图片高度
	Size   int    `json:"size,omitempty" bson:"size,omitempty" v:"required"`     // 图片大小
}

// 语音消息
type Voice struct {
	Url      string `json:"url,omitempty" bson:"url,omitempty" v:"required"`
	Duration int    `json:"duration,omitempty" bson:"duration,omitempty" v:"required"`
	Size     int    `json:"size,omitempty" bson:"size,omitempty" v:"required"` // 语音大小
}

// 视频文件消息
type Video struct {
	Url      string `json:"url,omitempty" bson:"url,omitempty" v:"required"`
	Duration int    `json:"duration,omitempty" bson:"duration,omitempty" v:"required"`
	Size     int    `json:"size,omitempty" bson:"size,omitempty" v:"required"` // 视频大小
	Cover    string `json:"cover,omitempty" bson:"cover,omitempty"`            // 封面图
}

// 文件消息
type File struct {
	UploadId string `json:"upload_id,omitempty" bson:"upload_id,omitempty" v:"required"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`     // 文件名称
	Drive    int    `json:"drive,omitempty" bson:"drive,omitempty"`   // 文件存储方式
	Suffix   string `json:"suffix,omitempty" bson:"suffix,omitempty"` // 文件后缀
	Size     int    `json:"size,omitempty" bson:"size,omitempty"`     // 文件大小
	Path     string `json:"path,omitempty" bson:"path,omitempty"`     // 文件路径
}

type Reply struct {
	MsgId    string `json:"msg_id,omitempty" bson:"msg_id,omitempty"`
	MsgType  string `json:"msg_type,omitempty" bson:"msg_type,omitempty"` // text: 文字, image: 图片
	Content  string `json:"content,omitempty" bson:"content,omitempty"`   // 文字或图片链接
	UserId   int    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty" bson:"nickname,omitempty"`
}
