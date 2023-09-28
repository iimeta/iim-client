package model

// 好友列表接口响应参数
type ContactListRes struct {
	Items []*Contact `json:"items"`
}

// 好友删除接口请求参数
type ContactDeleteReq struct {
	FriendId int `json:"friend_id,omitempty" v:"required"`
}

// 好友搜索接口请求参数
type ContactSearchReq struct {
	Mobile string `json:"mobile,omitempty" v:"required"`
}

// 好友搜索接口响应参数
type ContactSearchRes struct {
	Id       int    `json:"id,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Gender   int    `json:"gender,omitempty"`
	Motto    string `json:"motto,omitempty"`
}

// 好友备注修改接口请求参数
type ContactEditRemarkReq struct {
	FriendId int    `json:"friend_id,omitempty" v:"required"`
	Remark   string `json:"remark,omitempty"`
}

// 好友详情接口请求参数
type ContactDetailReq struct {
	UserId int `json:"user_id,omitempty" v:"required"`
}

// 好友详情接口响应参数
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

// 修改好友分组接口请求参数
type ContactChangeGroupReq struct {
	UserId  int    `json:"user_id,omitempty" v:"required"`
	GroupId string `json:"group_id,omitempty"`
}

// 添加好友申请接口请求参数
type ApplyCreateReq struct {
	ContactApply
}

// 同意好友申请接口请求参数
type ApplyAcceptReq struct {
	ContactApply
}

// 拒绝好友申请接口请求参数
type ApplyDeclineReq struct {
	ApplyId string `json:"apply_id,omitempty" v:"required"`
	Remark  string `json:"remark,omitempty" v:"required"`
}

// 好友申请列表接口响应参数
type ApplyListRes struct {
	Items []*Apply `json:"items"`
}

// 好友申请未读数接口响应参数
type ApplyUnreadNumRes struct {
	UnreadNum int `json:"unread_num"`
}

// 好友分组列表接口响应参数
type ContactGroupListRes struct {
	// 分组列表
	Items []*ContactGroup `json:"items"`
}

// 保存好友分组列表接口请求参数
type GroupSaveReq struct {
	Items []*ContactGroup `json:"items" v:"required"`
}

// 添加好友分组接口请求参数
type ContactGroupCreateReq struct {
	Name string `json:"name,omitempty" v:"required"`
	Sort int32  `json:"sort,omitempty" v:"required"`
}

// 添加好友分组接口响应参数
type ContactGroupCreateRes struct {
	Id int32 `json:"id,omitempty"`
}

// 更新好友分组接口请求参数
type GroupUpdateReq struct {
	Id   int32  `json:"id,omitempty" v:"required"`
	Name string `json:"name,omitempty" v:"required"`
	Sort int32  `json:"sort,omitempty" v:"required"`
}

// 更新好友分组接口响应参数
type GroupUpdateRes struct {
	Id int32 `json:"id,omitempty"`
}

// 删除好友分组接口请求参数
type GroupDeleteReq struct {
	Id int32 `json:"id,omitempty" v:"required"`
}

// 删除好友分组接口响应参数
type GroupDeleteRes struct {
	Id int32 `json:"id,omitempty"`
}

// 排序好友分组接口请求参数
type GroupSortReq struct {
	Items []*ContactGroup `json:"items" v:"required"`
}

type Contact struct {
	Id       int    `json:"id"`        // 用户ID
	Nickname string `json:"nickname"`  // 昵称
	Gender   int    `json:"gender"`    // 性别[0:未知;1:男;2:女;]
	Motto    string `json:"motto"`     // 座右铭
	Avatar   string `json:"avatar"`    // 头像
	Remark   string `json:"remark"`    // 备注
	IsOnline int    `json:"is_online"` // 是否在线
	GroupId  string `json:"group_id"`  // 好友分组ID
}

type ContactListItem struct {
	Id       int    `json:"id"`            // 用户ID
	Nickname string `json:"nickname"`      // 用户昵称
	Gender   int    `json:"gender"`        // 用户性别[0:未知;1:男;2:女;]
	Motto    string `json:"motto"`         // 用户座右铭
	Avatar   string `json:"avatar" `       // 好友头像
	Remark   string `json:"friend_remark"` // 好友的备注
	IsOnline int    `json:"is_online"`     // 是否在线
	GroupId  string `json:"group_id"`      // 好友分组ID
}

type Apply struct {
	Id        string `json:"id,omitempty"`         // 申请ID
	UserId    int    `json:"user_id,omitempty"`    // 申请人ID
	FriendId  int    `json:"friend_id,omitempty"`  // 被申请人
	Remark    string `json:"remark,omitempty"`     // 申请备注
	Nickname  string `json:"nickname,omitempty"`   // 申请备注
	Avatar    string `json:"avatar,omitempty"`     // 申请备注
	CreatedAt string `json:"created_at,omitempty"` // 申请时间
}

// 用户添加好友申请
type ContactApply struct {
	FriendId int    `json:"friend_id,omitempty"`
	Remark   string `json:"remark,omitempty"`
	ApplyId  string `json:"apply_id,omitempty"`
	UserId   int    `json:"user_id,omitempty"`
}

// 好友分组
type ContactGroup struct {
	Id    string `json:"id"`    // 主键ID
	Name  string `json:"name"`  // 分组名称
	Count int    `json:"count"` // 成员总数
	Sort  int    `json:"sort"`  // 分组排序
}
