package model

type SearchNoticeItem struct {
	Id           string `json:"id"`
	CreatorId    int    `json:"creator_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	IsTop        int    `json:"is_top"`
	IsConfirm    int    `json:"is_confirm"`
	ConfirmUsers string `json:"confirm_users"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
}

type GroupNoticeEdit struct {
	UserId    int
	GroupId   int
	NoticeId  string
	Title     string
	Content   string
	IsTop     int
	IsConfirm int
}

// 添加或编辑群公告接口请求参数
type GroupNoticeEditReq struct {
	GroupId   int    `json:"group_id,omitempty" v:"required"`
	NoticeId  string `json:"notice_id,omitempty"`
	Title     string `json:"title,omitempty" v:"required"`
	Content   string `json:"content,omitempty" v:"required"`
	IsTop     int    `json:"is_top,omitempty" v:"in:0,1"`
	IsConfirm int    `json:"is_confirm,omitempty" v:"in:0,1"`
}

// 删除群公告接口请求参数
type GroupNoticeDeleteReq struct {
	GroupId  int    `json:"group_id,omitempty" v:"required"`
	NoticeId string `json:"notice_id,omitempty" v:"required"`
}

// 群公告列表接口请求参数
type GroupNoticeListReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
}

// 群公告列表接口响应参数
type GroupNoticeListRes struct {
	Items []*GroupNoticeListResponse_Item `json:"items"`
}

type GroupNoticeListResponse_Item struct {
	Id           string `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Content      string `json:"content,omitempty"`
	IsTop        int    `json:"is_top,omitempty"`
	IsConfirm    int    `json:"is_confirm,omitempty"`
	ConfirmUsers string `json:"confirm_users,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	CreatorId    int    `json:"creator_id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}
