package model

// 添加联系人申请接口请求参数
type ContactApplyCreateReq struct {
	FriendId int    `json:"friend_id,omitempty" v:"required"`
	Remark   string `json:"remark,omitempty" v:"required"`
}

// 同意联系人申请接口请求参数
type ContactApplyAcceptReq struct {
	ApplyId string `json:"apply_id,omitempty" v:"required"`
	Remark  string `json:"remark,omitempty" v:"required"`
}

// 拒绝联系人申请接口请求参数
type ContactApplyDeclineReq struct {
	ApplyId string `json:"apply_id,omitempty" v:"required"`
	Remark  string `json:"remark,omitempty" v:"required"`
}

// 联系人申请列表接口响应参数
type ContactApplyListRes struct {
	Items []*ContactApplyListResponse_Item `json:"items"`
}

type ContactApplyListResponse_Item struct {
	Id        string `json:"id,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
	FriendId  int    `json:"friend_id,omitempty"`
	Remark    string `json:"remark,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
