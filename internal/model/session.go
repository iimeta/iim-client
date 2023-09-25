package model

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

// 会话列表
type TalkSessionItem struct {
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

// 会话未读数清除接口请求参数
type TalkSessionClearUnreadNumReq struct {
	TalkType   int `json:"talk_type,omitempty" v:"required|in:1,2"`
	ReceiverId int `json:"receiver_id,omitempty" v:"required"`
}
