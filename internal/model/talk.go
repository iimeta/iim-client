package model

type TalkRecords struct {
	Id         string `json:"id"`          // ID
	RecordId   int    `json:"record_id"`   // 记录ID
	MsgId      string `json:"msg_id"`      // 消息唯一ID
	Sequence   int64  `json:"sequence"`    // 消息时序ID
	TalkType   int    `json:"talk_type"`   // 对话类型[1:私信;2:群聊;]
	MsgType    int    `json:"msg_type"`    // 消息类型
	UserId     int    `json:"user_id"`     // 发送者ID[0:系统用户;]
	ReceiverId int    `json:"receiver_id"` // 接收者ID(用户ID 或 群ID)
	IsRevoke   int    `json:"is_revoke"`   // 是否撤回消息[0:否;1:是;]
	IsMark     int    `json:"is_mark"`     // 是否重要消息[0:否;1:是;]
	IsRead     int    `json:"is_read"`     // 是否已读[0:否;1:是;]
	QuoteId    string `json:"quote_id"`    // 引用消息ID
	Content    string `json:"content"`     // 文本消息
	Extra      string `json:"extra"`       // 扩展信信息
	CreatedAt  int64  `json:"created_at"`  // 创建时间
	UpdatedAt  int64  `json:"updated_at"`  // 更新时间
}

type TalkRecordExtraGroupMembers struct {
	UserId   int    `json:"user_id"`  // 用户ID
	Nickname string `json:"nickname"` // 用户昵称
}

type Reply struct {
	UserId   int    `json:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	MsgType  int    `json:"msg_type,omitempty"` // 1:文字 2:图片
	Content  string `json:"content,omitempty"`  // 文字或图片连接
	MsgId    string `json:"msg_id,omitempty"`
}

type TalkRecordExtraCode struct {
	Lang string `json:"lang"` // 代码语言
	Code string `json:"code"` // 代码内容
}

type TalkRecordExtraLocation struct {
	Longitude   string `json:"longitude"`   // 经度
	Latitude    string `json:"latitude"`    // 纬度
	Description string `json:"description"` // 位置描述
}

type TalkRecordExtraForward struct {
	TalkType   int              `json:"talk_type"`   // 对话类型
	UserId     int              `json:"user_id"`     // 发送者ID
	ReceiverId int              `json:"receiver_id"` // 接收者ID
	MsgIds     []int            `json:"msg_ids"`     // 消息列表
	Records    []map[string]any `json:"records"`     // 消息快照
}

type TalkRecordExtraLogin struct {
	IP       string `json:"ip"`       // 登录IP
	Address  string `json:"address"`  // 登录地址
	Agent    string `json:"agent"`    // 登录设备
	Platform string `json:"platform"` // 登录平台
	Reason   string `json:"reason"`   // 登录原因
	Datetime string `json:"datetime"` // 登录时间
}

type TalkRecordExtraCard struct {
	UserId int `json:"user_id"` // 名片用户ID
}

type TalkRecordExtraFile struct {
	Name   string `json:"name"`   // 文件名称
	Drive  int    `json:"drive"`  // 文件存储方式
	Suffix string `json:"suffix"` // 文件后缀
	Size   int    `json:"size"`   // 文件大小
	Path   string `json:"path"`   // 文件路径
}

type TalkRecordExtraImage struct {
	Name   string `json:"name"`   // 图片名称
	Suffix string `json:"suffix"` // 图片后缀
	Size   int    `json:"size"`   // 图片大小
	Url    string `json:"url"`    // 图片地址
	Width  int    `json:"width"`  // 图片宽度
	Height int    `json:"height"` // 图片高度
}

type TalkRecordExtraAudio struct {
	Name     string `json:"name"`     // 语音名称
	Suffix   string `json:"suffix"`   // 文件后缀
	Size     int    `json:"size"`     // 语音大小
	Url      string `json:"url"`      // 语音地址
	Duration int    `json:"duration"` // 语音时长
}

type TalkRecordExtraVideo struct {
	Name     string `json:"name"`     // 视频名称
	Cover    string `json:"cover"`    // 视频封面
	Suffix   string `json:"suffix"`   // 文件后缀
	Size     int    `json:"size"`     // 视频大小
	Url      string `json:"url"`      // 视频地址
	Duration int    `json:"duration"` // 视频时长
}

// 创建群消息
type TalkRecordExtraGroupCreate struct {
	OwnerId   int                            `json:"owner_id"`   // 操作人ID
	OwnerName string                         `json:"owner_name"` // 操作人昵称
	Members   []*TalkRecordExtraGroupMembers `json:"members"`    // 成员列表
}

// 群主邀请加入群消息
type TalkRecordExtraGroupJoin struct {
	OwnerId   int                            `json:"owner_id"`   // 操作人ID
	OwnerName string                         `json:"owner_name"` // 操作人昵称
	Members   []*TalkRecordExtraGroupMembers `json:"members"`    // 成员列表
}

// 群主转让群消息
type TalkRecordExtraGroupTransfer struct {
	OldOwnerId   int    `json:"old_owner_id"`   // 老群主ID
	OldOwnerName string `json:"old_owner_name"` // 老群主昵称
	NewOwnerId   int    `json:"new_owner_id"`   // 新群主ID
	NewOwnerName string `json:"new_owner_name"` // 新群主昵称
}

// 管理员设置群禁言消息
type TalkRecordExtraGroupMuted struct {
	OwnerId   int    `json:"owner_id"`   // 操作人ID
	OwnerName string `json:"owner_name"` // 操作人昵称
}

// 管理员解除群禁言消息
type TalkRecordExtraGroupCancelMuted struct {
	OwnerId   int    `json:"owner_id"`   // 操作人ID
	OwnerName string `json:"owner_name"` // 操作人昵称
}

// 管理员设置群成员禁言消息
type TalkRecordExtraGroupMemberMuted struct {
	OwnerId   int                            `json:"owner_id"`   // 操作人ID
	OwnerName string                         `json:"owner_name"` // 操作人昵称
	Members   []*TalkRecordExtraGroupMembers `json:"members"`    // 成员列表
}

// 管理员解除群成员禁言消息
type TalkRecordExtraGroupMemberCancelMuted struct {
	OwnerId   int                            `json:"owner_id"`   // 操作人ID
	OwnerName string                         `json:"owner_name"` // 操作人昵称
	Members   []*TalkRecordExtraGroupMembers `json:"members"`    // 成员列表
}

// 群主解散群消息
type TalkRecordExtraGroupDismissed struct {
	OwnerId   int    `json:"owner_id"`   // 操作人ID
	OwnerName string `json:"owner_name"` // 操作人昵称
}

// 群成员退出群消息
type TalkRecordExtraGroupMemberQuit struct {
	OwnerId   int    `json:"owner_id"`   // 操作人ID
	OwnerName string `json:"owner_name"` // 操作人昵称
}

// 踢出群成员消息
type TalkRecordExtraGroupMemberKicked struct {
	OwnerId   int                            `json:"owner_id"`   // 操作人ID
	OwnerName string                         `json:"owner_name"` // 操作人昵称
	Members   []*TalkRecordExtraGroupMembers `json:"members"`    // 成员列表
}

// 管理员撤回成员消息
type TalkRecordExtraGroupMessageRevoke struct {
	OwnerId         int    `json:"owner_id"`          // 操作人ID
	OwnerName       string `json:"owner_name"`        // 操作人昵称
	RevokeMessageId string `json:"revoke_message_id"` // 被撤回消息ID
}

// 发布群公告
type TalkRecordExtraGroupNotice struct {
	OwnerId   int    `json:"owner_id"`   // 操作人ID
	OwnerName string `json:"owner_name"` // 操作人昵称
	Title     string `json:"title"`      // 标题
	Content   string `json:"content"`    // 内容
}

type TalkRecordExtraMixedItem struct {
	Type    int    `json:"type"`           // 消息类型, 跟msgtype字段一致
	Content string `json:"content"`        // 消息内容, 可包含图片, 文字, 表情等多种消息
	Link    string `json:"link,omitempty"` // 图片跳转地址
}

// 图文混合消息
type TalkRecordExtraMixed struct {
	Items []*TalkRecordExtraMixedItem `json:"items"` // 消息内容, 可包含图片, 文字, 表情等多种消息
}

type RemoveRecordListOpt struct {
	UserId     int
	TalkType   int
	ReceiverId int
	RecordIds  string
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

type QueryTalkRecordsOpt struct {
	TalkType   int   // 对话类型
	UserId     int   // 获取消息的用户
	ReceiverId int   // 接收者ID
	MsgType    []int // 消息类型
	RecordId   int   // 上次查询的最小消息ID
	Limit      int   // 数据行数
}

type GetTalkRecordsReq struct {
	TalkType   int `json:"talk_type" v:"required|in:1,2"`  // 对话类型
	MsgType    int `json:"msg_type"`                       // 消息类型
	ReceiverId int `json:"receiver_id" v:"required|min:1"` // 接收者ID
	RecordId   int `json:"record_id"`                      // 上次查询的最小消息ID
	Limit      int `json:"limit" v:"required|max:100"`     // 数据行数
}

type GetTalkRecordsRes struct {
	Limit    int                `json:"limit"`
	RecordId int                `json:"record_id"`
	Items    []*TalkRecordsItem `json:"items"`
}

type GetForwardTalkRecordReq struct {
	RecordId int `json:"record_id"` // 上次查询的最小消息ID
}

type DownloadChatFileReq struct {
	RecordId int `json:"cr_id" v:"cr_id@required|min:1"`
}

type QueryVoteModel struct {
	RecordId     int    `json:"record_id"`
	ReceiverId   int    `json:"receiver_id"`
	TalkType     int    `json:"talk_type"`
	MsgType      int    `json:"msg_type"`
	VoteId       string `json:"vote_id"`
	AnswerMode   int    `json:"answer_mode"`
	AnswerOption string `json:"answer_option"`
	AnswerNum    int    `json:"answer_num"`
	VoteStatus   int    `json:"vote_status"`
}

type VoteStatistics struct {
	Count   int            `json:"count"`
	Options map[string]int `json:"options"`
}

// 会话创建接口请求参数
type TalkSessionCreateReq struct {
	TalkType   int `json:"talk_type,omitempty" v:"required|in:1,2"`
	ReceiverId int `json:"receiver_id,omitempty" v:"required"`
}

// 会话创建接口响应参数
type TalkSessionCreateRes struct {
	Id         string `json:"id,omitempty"`
	TalkType   int    `json:"talk_type,omitempty"`
	ReceiverId int    `json:"receiver_id,omitempty"`
	IsTop      int    `json:"is_top,omitempty"`
	IsDisturb  int    `json:"is_disturb,omitempty"`
	IsOnline   int    `json:"is_online,omitempty"`
	IsRobot    int    `json:"is_robot,omitempty"`
	Name       string `json:"name,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	Remark     string `json:"remark,omitempty"`
	UnreadNum  int    `json:"unread_num,omitempty"`
	MsgText    string `json:"msg_text,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	IsTalk     int    `json:"is_talk,omitempty"`
}

// 会话删除接口请求参数
type TalkSessionDeleteReq struct {
	ListId string `json:"list_id,omitempty" v:"required"`
}

// 会话置顶接口请求参数
type TalkSessionTopReq struct {
	ListId string `json:"list_id,omitempty" v:"required"`
	Type   int    `json:"type,omitempty" v:"required|in:1,2"`
}

// 会话免打扰接口请求参数
type TalkSessionDisturbReq struct {
	TalkType   int `json:"talk_type,omitempty" v:"required|in:1,2"`
	ReceiverId int `json:"receiver_id,omitempty" v:"required"`
	IsDisturb  int `json:"is_disturb,omitempty" v:"required|in:0,1"`
}

// 会话列表接口响应参数
type TalkSessionListRes struct {
	Items []*TalkSessionItem `json:"items"`
}

// 会话列表
type TalkSessionItem struct {
	Id            string `json:"id,omitempty"`
	TalkType      int    `json:"talk_type,omitempty"`
	ReceiverId    int    `json:"receiver_id,omitempty"`
	IsTop         int    `json:"is_top,omitempty"`
	IsDisturb     int    `json:"is_disturb,omitempty"`
	IsOnline      int    `json:"is_online,omitempty"`
	IsRobot       int    `json:"is_robot,omitempty"`
	Name          string `json:"name,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
	Remark        string `json:"remark,omitempty"`
	UnreadNum     int    `json:"unread_num,omitempty"`
	MsgText       string `json:"msg_text,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	IsTalk        int    `json:"is_talk,omitempty"`
	IsOpenContext int    `json:"is_open_context"`
}

// 会话未读数清除接口请求参数
type TalkSessionClearUnreadNumReq struct {
	TalkType   int `json:"talk_type,omitempty" v:"required|in:1,2"`
	ReceiverId int `json:"receiver_id,omitempty" v:"required"`
}

type SearchTalkSession struct {
	Id            string `json:"id" `
	TalkType      int    `json:"talk_type" `
	ReceiverId    int    `json:"receiver_id" `
	IsDelete      int    `json:"is_delete"`
	IsTop         int    `json:"is_top"`
	IsRobot       int    `json:"is_robot"`
	IsDisturb     int    `json:"is_disturb"`
	UserAvatar    string `json:"user_avatar"`
	Nickname      string `json:"nickname"`
	GroupName     string `json:"group_name"`
	GroupAvatar   string `json:"group_avatar"`
	UpdatedAt     int64  `json:"updated_at"`
	IsTalk        int    `json:"is_talk"`
	IsOpenContext int    `json:"is_open_context"`
}

type TalkClearContextReq struct {
	ReceiverId int `json:"receiver_id"`
}

type TalkOpenContextReq struct {
	ReceiverId    int `json:"receiver_id"`
	IsOpenContext int `json:"is_open_context"`
	TalkType      int `json:"talk_type,omitempty" v:"required|in:1,2"`
}
