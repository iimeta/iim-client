package model

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
	Items []*GroupListResponse_Item `json:"items"`
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
	Mode    int `json:"mode,omitempty" v:"required|in:1,2"` // 操作方式  1:开启全员禁言  2:解除全员禁言
}

// 群公开修改接口请求参数
type GroupOvertReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
	Mode    int `json:"mode,omitempty" v:"required|in:1,2"` // 操作方式  1:开启公开可见  2:关闭公开可见
}

type ApplyList struct {
	Id        string `json:"id"`         // 自增ID
	GroupId   int    `json:"group_id"`   // 群聊ID
	UserId    int    `json:"user_id"`    // 用户ID
	Remark    string `json:"remark"`     // 备注信息
	CreatedAt int64  `json:"created_at"` // 创建时间
	Nickname  string `json:"nickname"`   // 用户昵称
	Avatar    string `json:"avatar"`     // 用户头像地址
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
	Items []*GroupApplyListResponse_Item `json:"items"`
}

// 申请管理-所有入群申请列表接口响应参数
type ApplyAllRes struct {
	Items []*ApplyAllResponse_Item `json:"items"`
}

type GroupApplyListResponse_Item struct {
	Id        string `json:"id,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
	GroupId   int    `json:"group_id,omitempty"`
	Remark    string `json:"remark,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type ApplyAllResponse_Item struct {
	Id        string `json:"id,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
	GroupId   int    `json:"group_id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
	Remark    string `json:"remark,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type GroupApplyUnreadNumRes struct {
	UnreadNum int `json:"unread_num"`
}

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

type NoticeEdit struct {
	UserId    int
	GroupId   int
	NoticeId  string
	Title     string
	Content   string
	IsTop     int
	IsConfirm int
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
	Items []*NoticeListResponse_Item `json:"items"`
}

type NoticeListResponse_Item struct {
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
