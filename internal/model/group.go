package model

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
	Items []*Group `json:"items"`
}

// 群成员列表接口请求参数
type GroupMemberListReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
}

// 群成员列表接口响应参数
type GroupMemberListRes struct {
	Items []*GroupMember `json:"items"`
}

// 公开群聊列表接口响应参数
type GroupOvertListRes struct {
	Items []*GroupOvert `json:"items"`
	Next  bool          `json:"next,omitempty"`
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
	Mode    int `json:"mode,omitempty" v:"required|in:1,2"` // 操作方式  1:开启全员禁言  2:解除全员禁言
}

// 群公开修改接口请求参数
type GroupOvertReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
	Mode    int `json:"mode,omitempty" v:"required|in:1,2"` // 操作方式  1:开启公开可见  2:关闭公开可见
}

// 提交入群申请接口请求参数
type GroupApplyCreateReq struct {
	GroupId int    `json:"group_id,omitempty" v:"required"`
	Remark  string `json:"remark,omitempty" v:"required"`
}

// 拒绝入群申请接口请求参数
type ApplyDeleteReq struct {
	ApplyId string `json:"apply_id,omitempty" v:"required"`
}

// 同意入群申请接口请求参数
type ApplyAgreeReq struct {
	ApplyId string `json:"apply_id,omitempty" v:"required"`
}

// 拒绝入群申请接口请求参数
type GroupApplyDeclineReq struct {
	ApplyId string `json:"apply_id,omitempty" v:"required"`
	Remark  string `json:"remark,omitempty" v:"required"`
}

// 入群申请列表接口请求参数
type ApplyListReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
}

// 入群申请列表接口响应参数
type GroupApplyListRes struct {
	Items []*GroupApply `json:"items"`
}

// 申请管理-所有入群申请列表接口响应参数
type ApplyAllRes struct {
	Items []*GroupApply `json:"items"`
}

type GroupApplyUnreadNumRes struct {
	UnreadNum int `json:"unread_num"`
}

// 添加或编辑群公告接口请求参数
type NoticeEditReq struct {
	GroupId   int    `json:"group_id,omitempty" v:"required"`
	NoticeId  string `json:"notice_id,omitempty"`
	Title     string `json:"title,omitempty" v:"required"`
	Content   string `json:"content,omitempty" v:"required"`
	IsTop     int    `json:"is_top,omitempty" v:"in:0,1"`
	IsConfirm int    `json:"is_confirm,omitempty" v:"in:0,1"`
}

// 删除群公告接口请求参数
type NoticeDeleteReq struct {
	GroupId  int    `json:"group_id,omitempty" v:"required"`
	NoticeId string `json:"notice_id,omitempty" v:"required"`
}

// 群公告列表接口请求参数
type NoticeListReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
}

// 群公告列表接口响应参数
type NoticeListRes struct {
	Items []*Notice `json:"items"`
}

type Group struct {
	Id        int    `json:"id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Profile   string `json:"profile,omitempty"`
	Leader    int    `json:"leader,omitempty"`
	IsDisturb int    `json:"is_disturb,omitempty"`
	CreatorId int    `json:"creator_id,omitempty"`
}

type GroupMember struct {
	UserId   int    `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	Leader   int    `json:"leader"`
	IsMute   int    `json:"is_mute"`
	Remark   string `json:"remark"`
}

type GroupOvert struct {
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

type GroupApply struct {
	Id        string `json:"id,omitempty"`         // ID
	UserId    int    `json:"user_id,omitempty"`    // 用户ID
	GroupId   int    `json:"group_id,omitempty"`   // 群聊ID
	GroupName string `json:"group_name,omitempty"` // 群名称
	Remark    string `json:"remark,omitempty"`     // 备注信息
	Avatar    string `json:"avatar,omitempty"`     // 用户头像地址
	Nickname  string `json:"nickname,omitempty"`   // 用户昵称
	CreatedAt string `json:"created_at,omitempty"` // 创建时间
}

type Notice struct {
	Id           string `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Content      string `json:"content,omitempty"`
	IsTop        int    `json:"is_top,omitempty"`
	IsConfirm    int    `json:"is_confirm,omitempty"`
	ConfirmUsers string `json:"confirm_users,omitempty"`
	CreatorId    int    `json:"creator_id,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Nickname     string `json:"nickname,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

type NoticeEdit struct {
	UserId    int
	GroupId   int
	NoticeId  string
	Title     string
	Content   string
	IsTop     int
	IsConfirm int
}

type GroupAuth struct {
	TalkType          int
	UserId            int
	ReceiverId        int
	IsVerifyGroupMute bool
}
