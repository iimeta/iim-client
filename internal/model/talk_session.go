package model

// 会话创建接口请求参数
type SessionCreateReq struct {
	TalkType   int `json:"talk_type,omitempty" v:"required|in:1,2"`
	ReceiverId int `json:"receiver_id,omitempty" v:"required"`
}

// 会话创建接口响应参数
type SessionCreateRes struct {
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
type SessionDeleteReq struct {
	ListId string `json:"list_id,omitempty" v:"required"`
}

// 会话置顶接口请求参数
type SessionTopReq struct {
	ListId string `json:"list_id,omitempty" v:"required"`
	Type   int    `json:"type,omitempty" v:"required|in:1,2"`
}

// 会话免打扰接口请求参数
type SessionDisturbReq struct {
	TalkType   int `json:"talk_type,omitempty" v:"required|in:1,2"`
	ReceiverId int `json:"receiver_id,omitempty" v:"required"`
	IsDisturb  int `json:"is_disturb,omitempty" v:"required|in:0,1"`
}

// 会话列表接口响应参数
type SessionListRes struct {
	Items []*TalkSession `json:"items"`
}

// 会话未读数清除接口请求参数
type SessionClearUnreadNumReq struct {
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

type SessionClearContextReq struct {
	ReceiverId int `json:"receiver_id"`
}

type SessionOpenContextReq struct {
	ReceiverId    int `json:"receiver_id"`
	IsOpenContext int `json:"is_open_context"`
	TalkType      int `json:"talk_type,omitempty" v:"required|in:1,2"`
}

type TalkSession struct {
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
