package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

const (
	GroupMemberMaxNum = 200 // 最大成员数量
)

type GroupItem struct {
	Id        int    `json:"id"`
	GroupName string `json:"group_name"`
	Avatar    string `json:"avatar"`
	Profile   string `json:"profile"`
	Leader    int    `json:"leader"`
	IsDisturb int    `json:"is_disturb"`
	CreatorId int    `json:"creator_id"`
}

type AuthOption struct {
	TalkType          int
	UserId            int
	ReceiverId        int
	IsVerifyGroupMute bool
}

// 创建群聊接口请求参数
type GroupCreateReq struct {
	Name   string `json:"name,omitempty" v:"required"`
	Ids    string `json:"ids,omitempty" v:"required"`
	Avatar string `json:"avatar,omitempty"`
}

// 创建群聊接口响应参数
type GroupCreateRes struct {
	GroupId int `json:"group_id,omitempty"`
}

// 解散群聊接口请求参数
type GroupDismissReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
}

// 邀请加入群聊接口请求参数
type GroupInviteReq struct {
	GroupId int    `json:"group_id,omitempty" v:"required"`
	Ids     string `json:"ids,omitempty" v:"required"`
}

// 退出群聊接口请求参数
type GroupSecedeReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
}

// 设置群聊接口请求参数
type GroupSettingReq struct {
	GroupId   int    `json:"group_id,omitempty" v:"required"`
	GroupName string `json:"group_name,omitempty" v:"required"`
	Avatar    string `json:"avatar,omitempty"`
	Profile   string `json:"profile,omitempty" v:"max-length:255"`
}

// 移出群成员接口请求参数
type GroupRemoveMemberReq struct {
	GroupId    int    `json:"group_id,omitempty" v:"required"`
	MembersIds string `json:"members_ids,omitempty" v:"required"`
}

// 群聊详情接口请求参数
type GroupDetailReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
}

// 群聊详情接口响应参数
type GroupDetailRes struct {
	GroupId   int    `json:"group_id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
	Profile   string `json:"profile,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	IsManager bool   `json:"is_manager,omitempty"`
	IsDisturb int    `json:"is_disturb,omitempty"`
	VisitCard string `json:"visit_card,omitempty"`
	IsMute    int    `json:"is_mute,omitempty"`
	IsOvert   int    `json:"is_overt,omitempty"`
}

// 群聊名片更新接口请求参数
type GroupRemarkUpdateReq struct {
	GroupId   int    `json:"group_id,omitempty" v:"required"`
	VisitCard string `json:"visit_card,omitempty"`
}

// 获取待审批入群申请列表接口请求参数
type GetInviteFriendsReq struct {
	GroupId int `json:"group_id,omitempty"`
}

// 群列表接口响应参数
type GroupListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	Items  []*GroupListResponse_Item `json:"items"`
}

type GroupListResponse_Item struct {
	Id        int    `json:"id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Profile   string `json:"profile,omitempty"`
	Leader    int    `json:"leader,omitempty"`
	IsDisturb int    `json:"is_disturb,omitempty"`
	CreatorId int    `json:"creator_id,omitempty"`
}

// 群成员列表接口请求参数
type GroupMemberListReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
}

// 群成员列表接口响应参数
type GroupMemberListRes struct {
	Items []*GroupMemberListResponse_Item `json:"items"`
}

type GroupMemberListResponse_Item struct {
	UserId   int    `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	Leader   int    `json:"leader"`
	IsMute   int    `json:"is_mute"`
	Remark   string `json:"remark"`
}

// 公开群聊列表接口响应参数
type GroupOvertListRes struct {
	Items []*GroupOvertListResponse_Item `json:"items"`
	Next  bool                           `json:"next,omitempty"`
}

type GroupOvertListResponse_Item struct {
	Id        int    `json:"id,omitempty"`
	Type      int    `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Profile   string `json:"profile,omitempty"`
	Count     int    `json:"count,omitempty"`
	MaxNum    int    `json:"max_num,omitempty"`
	IsMember  bool   `json:"is_member,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

// 公开群聊列表接口请求参数
type GroupOvertListReq struct {
	Page int    `json:"page,omitempty" v:"required"`
	Name string `json:"name,omitempty" v:"max-length:50"`
}

// 群主更换接口请求参数
type GroupHandoverReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
	UserId  int `json:"user_id,omitempty" v:"required"`
}

// 分配管理员接口请求参数
type GroupAssignAdminReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
	UserId  int `json:"user_id,omitempty" v:"required"`
	Mode    int `json:"mode,omitempty" v:"required|in:1,2"`
}

// 群成员禁言接口请求参数
type GroupNoSpeakReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
	UserId  int `json:"user_id,omitempty" v:"required"`
	Mode    int `json:"mode,omitempty" v:"required|in:1,2"`
}

// 全员禁言接口请求参数
type GroupMuteReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
	// 操作方式  1:开启全员禁言  2:解除全员禁言
	Mode int `json:"mode,omitempty" v:"required|in:1,2"`
}

// 群公开修改接口请求参数
type GroupOvertReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
	// 操作方式  1:开启公开可见  2:关闭公开可见
	Mode int `json:"mode,omitempty" v:"required|in:1,2"`
}
