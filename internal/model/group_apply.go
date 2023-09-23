package model

const (
	GroupApplyStatusWait   = 1 // 待处理
	GroupApplyStatusPass   = 2 // 通过
	GroupApplyStatusRefuse = 3 // 拒绝
)

type GroupApplyList struct {
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

// 同意入群申请接口请求参数
type GroupApplyAgreeReq struct {
	ApplyId string `json:"apply_id,omitempty" v:"required"`
}

// 拒绝入群申请接口请求参数
type GroupApplyDeclineReq struct {
	ApplyId string `json:"apply_id,omitempty" v:"required"`
	Remark  string `json:"remark,omitempty" v:"required"`
}

// 入群申请列表接口请求参数
type GroupApplyListReq struct {
	GroupId int `json:"group_id,omitempty" v:"required"`
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

// 入群申请列表接口响应参数
type GroupApplyListRes struct {
	Items []*GroupApplyListResponse_Item `json:"items"`
}

// 申请管理-所有入群申请列表接口响应参数
type GroupApplyAllRes struct {
	Items []*GroupApplyAllResponse_Item `json:"items"`
}

type GroupApplyAllResponse_Item struct {
	Id        string `json:"id,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
	GroupId   int    `json:"group_id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
	Remark    string `json:"remark,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type ApplyUnreadNumRes struct {
	UnreadNum int `json:"unread_num"`
}
