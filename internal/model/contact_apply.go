package model

// 用户添加好友申请
type ContactApply struct {
	ApplyId  string `json:"apply_id"`
	UserId   int    `json:"user_id"`
	Remarks  string `json:"remarks"`
	FriendId int    `json:"friend_id"`
}

type ApplyItem struct {
	Id        string `json:"id"`         // 申请ID
	UserId    int    `json:"user_id"`    // 申请人ID
	FriendId  int    `json:"friend_id"`  // 被申请人
	Remark    string `json:"remark"`     // 申请备注
	Nickname  string `json:"nickname"`   // 申请备注
	Avatar    string `json:"avatar"`     // 申请备注
	CreatedAt int64  `json:"created_at"` // 申请时间
}
