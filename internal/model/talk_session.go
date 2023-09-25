package model

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

type TalkOpenContextReq struct {
	ReceiverId    int `json:"receiver_id"`
	IsOpenContext int `json:"is_open_context"`
	TalkType      int `json:"talk_type,omitempty" v:"required|in:1,2"`
}
