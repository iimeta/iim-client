package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	CONTACT_COLLECTION       = "contact"
	CONTACT_APPLY_COLLECTION = "contact_apply"
	CONTACT_GROUP_COLLECTION = "contact_group"
)

type Contact struct {
	gmeta.Meta `collection:"contact" bson:"-"`
	UserId     int    `bson:"user_id,omitempty"`    // 用户id
	FriendId   int    `bson:"friend_id,omitempty"`  // 好友id
	Remark     string `bson:"remark,omitempty"`     // 好友的备注
	Status     int    `bson:"status,omitempty"`     // 好友状态 [0:否;1:是]
	GroupId    string `bson:"group_id,omitempty"`   // 分组id
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}

type ContactApply struct {
	gmeta.Meta `collection:"contact_apply" bson:"-"`
	UserId     int    `bson:"user_id,omitempty"`    // 申请人ID
	Nickname   string `bson:"nickname,omitempty"`   // 申请人昵称
	Avatar     string `bson:"avatar,omitempty"`     // 申请人头像地址
	FriendId   int    `bson:"friend_id,omitempty"`  // 被申请人
	Remark     string `bson:"remark,omitempty"`     // 申请备注
	CreatedAt  int64  `bson:"created_at,omitempty"` // 申请时间
}

// 好友分组
type ContactGroup struct {
	gmeta.Meta `collection:"contact_group" bson:"-"`
	Id         string `bson:"_id"`        // ID todo
	UserId     int    `bson:"user_id"`    // 用户ID
	Name       string `bson:"remark"`     // 分组名称
	Num        int    `bson:"num"`        // 成员总数
	Sort       int    `bson:"sort"`       // 分组名称
	CreatedAt  int64  `bson:"created_at"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at"` // 更新时间
}
