package model

import (
	"time"
)

const (
	ContactStatusNormal = 1
	ContactStatusDelete = 0
)

// Contact 用户好友关系表
type Contact struct {
	Id        string    `json:"id"`         // 关系ID
	UserId    int       `json:"user_id"`    // 用户id
	FriendId  int       `json:"friend_id"`  // 好友id
	Remark    string    `json:"remark"`     // 好友的备注
	Status    int       `json:"status"`     // 好友状态 [0:否;1:是]
	GroupId   int       `json:"group_id"`   // 分组id
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

type ContactListItem struct {
	Id       int    `json:"id"`            // 用户ID
	Nickname string `json:"nickname"`      // 用户昵称
	Gender   int    `json:"gender"`        // 用户性别[0:未知;1:男;2:女;]
	Motto    string `json:"motto"`         // 用户座右铭
	Avatar   string `json:"avatar" `       // 好友头像
	Remark   string `json:"friend_remark"` // 好友的备注
	IsOnline int    `json:"is_online"`     // 是否在线
	GroupId  string `json:"group_id"`      // 联系人分组
}

// 联系人列表接口响应参数
type ContactListRes struct {
	Items []*ContactListResponse_Item `json:"items"`
}

type ContactListResponse_Item struct {

	// 用户ID
	Id int `json:"id"`
	// 昵称
	Nickname string `json:"nickname"`
	// 性别[0:未知;1:男;2:女;]
	Gender int `json:"gender"`
	// 座右铭
	Motto string `json:"motto"`
	// 头像
	Avatar string `json:"avatar"`
	// 备注
	Remark string `json:"remark"`
	// 是否在线
	IsOnline int `json:"is_online"`
	// 联系人分组ID
	GroupId string `json:"group_id"`
}

// 联系人删除接口请求参数
type ContactDeleteReq struct {
	FriendId int `json:"friend_id,omitempty" v:"required"`
}

// 联系人搜索接口请求参数
type ContactSearchReq struct {
	Mobile string `json:"mobile,omitempty" v:"required"`
}

// 联系人搜索接口响应参数
type ContactSearchRes struct {
	Id       int    `json:"id,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Gender   int    `json:"gender,omitempty"`
	Motto    string `json:"motto,omitempty"`
}

// 联系人备注修改接口请求参数
type ContactEditRemarkReq struct {
	FriendId int    `json:"friend_id,omitempty" v:"required"`
	Remark   string `json:"remark,omitempty"`
}

// 联系人详情接口请求参数
type ContactDetailReq struct {
	UserId int `json:"user_id,omitempty" v:"required"`
}

// 联系人详情接口响应参数
type ContactDetailRes struct {
	Id           int    `json:"id"`
	Mobile       string `json:"mobile"`
	Nickname     string `json:"nickname"`
	Remark       string `json:"remark"`
	Avatar       string `json:"avatar"`
	Gender       int    `json:"gender"`
	Motto        string `json:"motto"`
	FriendApply  int    `json:"friend_apply"`
	FriendStatus int    `json:"friend_status"`
	GroupId      string `json:"group_id"`
	Email        string `json:"email"`
}

// 修改联系人分组接口请求参数
type ContactChangeGroupReq struct {
	UserId  int    `json:"user_id,omitempty" v:"required"`
	GroupId string `json:"group_id,omitempty"`
}
