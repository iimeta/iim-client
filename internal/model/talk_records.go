package model

type TalkRecordsReq struct {
	TalkType   int `json:"talk_type" v:"required|in:1,2"`  // 对话类型
	MsgType    int `json:"msg_type"`                       // 消息类型
	ReceiverId int `json:"receiver_id" v:"required|min:1"` // 接收者ID
	RecordId   int `json:"record_id"`                      // 上次查询的最小消息ID
	Limit      int `json:"limit" v:"required|max:100"`     // 数据行数
}

type TalkRecordsRes struct {
	Limit    int                `json:"limit"`
	RecordId int                `json:"record_id"`
	Items    []*TalkRecordsItem `json:"items"`
}

type RecordsForwardReq struct {
	RecordId int `json:"record_id"` // 上次查询的最小消息ID
}

type RecordsFileDownloadReq struct {
	RecordId int `json:"cr_id" v:"cr_id@required|min:1"`
}

type TalkRecord struct {
	Id         string `json:"id"`          // ID
	RecordId   int    `json:"record_id"`   // 记录ID
	MsgId      string `json:"msg_id"`      // 消息唯一ID
	Sequence   int64  `json:"sequence"`    // 消息时序ID
	TalkType   int    `json:"talk_type"`   // 对话类型[1:私聊;2:群聊;]
	MsgType    int    `json:"msg_type"`    // 消息类型
	UserId     int    `json:"user_id"`     // 发送者ID[0:系统用户;]
	ReceiverId int    `json:"receiver_id"` // 接收者ID(用户ID 或 群ID)
	IsRevoke   int    `json:"is_revoke"`   // 是否撤回消息[0:否;1:是;]
	IsMark     int    `json:"is_mark"`     // 是否重要消息[0:否;1:是;]
	IsRead     int    `json:"is_read"`     // 是否已读[0:否;1:是;]
	QuoteId    string `json:"quote_id"`    // 引用消息ID
	Content    string `json:"content"`     // 文本消息
	Extra      string `json:"extra"`       // 扩展信信息
	CreatedAt  string `json:"created_at"`  // 创建时间
	UpdatedAt  string `json:"updated_at"`  // 更新时间

	Reply *Reply `json:"reply,omitempty"`

	Text     *Text     `json:"text,omitempty"`
	Code     *Code     `json:"code,omitempty"`
	Image    *Image    `json:"image,omitempty"`
	Voice    *Voice    `json:"voice,omitempty"`
	Video    *Video    `json:"video,omitempty"`
	File     *File     `json:"file,omitempty"`
	Vote     *Vote     `json:"vote,omitempty"`
	Mixed    *Mixed    `json:"mixed,omitempty"`
	Emoticon *Emoticon `json:"emoticon,omitempty"`
	Card     *Card     `json:"card,omitempty"`
	Location *Location `json:"location,omitempty"`

	Login *Login `json:"login,omitempty"`
}

type TalkRecordsItem struct {
	Id         int    `json:"id"`
	Sequence   int    `json:"sequence"`
	MsgId      string `json:"msg_id"`
	TalkType   int    `json:"talk_type"`
	MsgType    int    `json:"msg_type"`
	UserId     int    `json:"user_id"`
	ReceiverId int    `json:"receiver_id"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	IsRevoke   int    `json:"is_revoke"`
	IsMark     int    `json:"is_mark"`
	IsRead     int    `json:"is_read"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	Extra      any    `json:"extra"` // 额外参数
}

type TalkGroupMember struct {
	UserId   int    `json:"user_id"`  // 用户ID
	Nickname string `json:"nickname"` // 用户昵称
}

type TalkRecordReply struct {
	UserId   int    `json:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	MsgType  int    `json:"msg_type,omitempty"` // 1:文字 2:图片
	Content  string `json:"content,omitempty"`  // 文字或图片链接
	MsgId    string `json:"msg_id,omitempty"`
}

type TalkRecordCode struct {
	Lang string `json:"lang"` // 代码语言
	Code string `json:"code"` // 代码内容
}

type TalkRecordLocation struct {
	Longitude   string `json:"longitude"`   // 经度
	Latitude    string `json:"latitude"`    // 纬度
	Description string `json:"description"` // 位置描述
}

type TalkRecordForward struct {
	TalkType   int                 `json:"talk_type"`   // 对话类型
	UserId     int                 `json:"user_id"`     // 发送者ID
	ReceiverId int                 `json:"receiver_id"` // 接收者ID
	RecordsIds []int               `json:"records_ids"` // 消息列表
	Records    []map[string]string `json:"records"`     // 消息快照
}

type TalkRecordLogin struct {
	IP       string `json:"ip"`       // 登录IP
	Address  string `json:"address"`  // 登录地址
	Agent    string `json:"agent"`    // 登录设备
	Platform string `json:"platform"` // 登录平台
	Reason   string `json:"reason"`   // 登录原因
	Datetime string `json:"datetime"` // 登录时间
}

type TalkRecordCard struct {
	UserId int `json:"user_id"` // 名片用户ID
}

type TalkRecordFile struct {
	Name   string `json:"name"`   // 文件名称
	Drive  int    `json:"drive"`  // 文件存储方式
	Suffix string `json:"suffix"` // 文件后缀
	Size   int    `json:"size"`   // 文件大小
	Path   string `json:"path"`   // 文件路径
}

type TalkRecordImage struct {
	Name   string `json:"name"`   // 图片名称
	Suffix string `json:"suffix"` // 图片后缀
	Size   int    `json:"size"`   // 图片大小
	Url    string `json:"url"`    // 图片地址
	Width  int    `json:"width"`  // 图片宽度
	Height int    `json:"height"` // 图片高度
}

type TalkRecordAudio struct {
	Name     string `json:"name"`     // 语音名称
	Suffix   string `json:"suffix"`   // 文件后缀
	Size     int    `json:"size"`     // 语音大小
	Url      string `json:"url"`      // 语音地址
	Duration int    `json:"duration"` // 语音时长
}

type TalkRecordVideo struct {
	Name     string `json:"name"`     // 视频名称
	Cover    string `json:"cover"`    // 视频封面
	Suffix   string `json:"suffix"`   // 文件后缀
	Size     int    `json:"size"`     // 视频大小
	Url      string `json:"url"`      // 视频地址
	Duration int    `json:"duration"` // 视频时长
}

// 创建群消息
type TalkRecordGroupCreate struct {
	OwnerId   int                `json:"owner_id"`   // 操作人ID
	OwnerName string             `json:"owner_name"` // 操作人昵称
	Members   []*TalkGroupMember `json:"members"`    // 成员列表
}

// 群主邀请加入群消息
type TalkRecordGroupJoin struct {
	OwnerId   int                `json:"owner_id"`   // 操作人ID
	OwnerName string             `json:"owner_name"` // 操作人昵称
	Members   []*TalkGroupMember `json:"members"`    // 成员列表
}

// 群主转让群消息
type TalkRecordGroupTransfer struct {
	OldOwnerId   int    `json:"old_owner_id"`   // 老群主ID
	OldOwnerName string `json:"old_owner_name"` // 老群主昵称
	NewOwnerId   int    `json:"new_owner_id"`   // 新群主ID
	NewOwnerName string `json:"new_owner_name"` // 新群主昵称
}

// 管理员设置群禁言消息
type TalkRecordGroupMuted struct {
	OwnerId   int    `json:"owner_id"`   // 操作人ID
	OwnerName string `json:"owner_name"` // 操作人昵称
}

// 管理员解除群禁言消息
type TalkRecordGroupCancelMuted struct {
	OwnerId   int    `json:"owner_id"`   // 操作人ID
	OwnerName string `json:"owner_name"` // 操作人昵称
}

// 管理员设置群成员禁言消息
type TalkRecordGroupMemberMuted struct {
	OwnerId   int                `json:"owner_id"`   // 操作人ID
	OwnerName string             `json:"owner_name"` // 操作人昵称
	Members   []*TalkGroupMember `json:"members"`    // 成员列表
}

// 管理员解除群成员禁言消息
type TalkRecordGroupMemberCancelMuted struct {
	OwnerId   int                `json:"owner_id"`   // 操作人ID
	OwnerName string             `json:"owner_name"` // 操作人昵称
	Members   []*TalkGroupMember `json:"members"`    // 成员列表
}

// 群主解散群消息
type TalkRecordGroupDismissed struct {
	OwnerId   int    `json:"owner_id"`   // 操作人ID
	OwnerName string `json:"owner_name"` // 操作人昵称
}

// 群成员退出群消息
type TalkRecordGroupMemberQuit struct {
	OwnerId   int    `json:"owner_id"`   // 操作人ID
	OwnerName string `json:"owner_name"` // 操作人昵称
}

// 踢出群成员消息
type TalkRecordGroupMemberKicked struct {
	OwnerId   int                `json:"owner_id"`   // 操作人ID
	OwnerName string             `json:"owner_name"` // 操作人昵称
	Members   []*TalkGroupMember `json:"members"`    // 成员列表
}

// 管理员撤回成员消息
type TalkRecordGroupMessageRevoke struct {
	OwnerId         int    `json:"owner_id"`          // 操作人ID
	OwnerName       string `json:"owner_name"`        // 操作人昵称
	RevokeMessageId string `json:"revoke_message_id"` // 被撤回消息ID
}

// 发布群公告
type TalkRecordGroupNotice struct {
	OwnerId   int    `json:"owner_id"`   // 操作人ID
	OwnerName string `json:"owner_name"` // 操作人昵称
	Title     string `json:"title"`      // 标题
	Content   string `json:"content"`    // 内容
}

// 图文混合消息
type TalkRecordMixed struct {
	Items []*MixedMessage `json:"items"` // 消息内容, 可包含图片, 文字, 表情等多种消息
}

type TalkRecordMixedItem struct {
	Type    int    `json:"type"`           // 消息类型, 跟msgtype字段一致
	Content string `json:"content"`        // 消息内容, 可包含图片, 文字, 表情等多种消息
	Link    string `json:"link,omitempty"` // 图片跳转地址
}

type QueryTalkRecordsOpt struct {
	TalkType   int   // 对话类型
	UserId     int   // 获取消息的用户
	ReceiverId int   // 接收者ID
	MsgType    []int // 消息类型
	RecordId   int   // 上次查询的最小消息ID
	Limit      int   // 数据行数
}

type VoteStatistics struct {
	Count   int            `json:"count"`
	Options map[string]int `json:"options"`
}
