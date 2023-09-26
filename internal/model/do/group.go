package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	GROUP_COLLECTION        = "group"
	GROUP_MEMBER_COLLECTION = "group_member"
	GROUP_APPLY_COLLECTION  = "group_apply"
	GROUP_NOTICE_COLLECTION = "group_notice"
)

type Group struct {
	gmeta.Meta `collection:"group" bson:"-"`
	GroupId    int    `bson:"group_id,omitempty"`   // 群聊ID
	Type       int    `bson:"type,omitempty"`       // 群类型[1:普通群;2:企业群;]
	CreatorId  int    `bson:"creator_id,omitempty"` // 创建者ID(群主ID)
	Name       string `bson:"group_name,omitempty"` // 群名称
	Profile    string `bson:"profile,omitempty"`    // 群介绍
	IsDismiss  int    `bson:"is_dismiss,omitempty"` // 是否已解散[0:否;1:是;]
	Avatar     string `bson:"avatar,omitempty"`     // 群头像
	MaxNum     int    `bson:"max_num,omitempty"`    // 最大群成员数量
	IsOvert    int    `bson:"is_overt,omitempty"`   // 是否公开可见[0:否;1:是;]
	IsMute     int    `bson:"is_mute,omitempty"`    // 是否全员禁言 [0:否;1:是;], 提示:不包含群主或管理员
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}

type SearchOvert struct {
	Name   string
	UserId int
	Page   int64
	Size   int64
}

type GroupCreate struct {
	UserId    int    // 操作人ID
	Name      string // 群名称
	Avatar    string // 群头像
	Profile   string // 群简介
	MemberIds []int  // 好友ID
}

type GroupUpdate struct {
	GroupId int    // 群ID
	Name    string // 群名称
	Avatar  string // 群头像
	Profile string // 群简介
}

type GroupInvite struct {
	UserId    int   // 操作人ID
	GroupId   int   // 群ID
	MemberIds []int // 群成员ID
}

type GroupMember struct {
	gmeta.Meta  `collection:"group_member" bson:"-"`
	GroupId     int    `bson:"group_id,omitempty"`      // 群聊ID
	UserId      int    `bson:"user_id,omitempty"`       // 用户ID
	Leader      int    `bson:"leader"`                  // 成员属性[0:普通成员;1:管理员;2:群主;]
	UserCard    string `bson:"user_card,omitempty"`     // 群名片
	IsQuit      int    `bson:"is_quit"`                 // 是否退群[0:否;1:是;]
	IsMute      int    `bson:"is_mute"`                 // 是否禁言[0:否;1:是;]
	MinRecordId int    `bson:"min_record_id,omitempty"` // 可查看历史记录最小ID
	JoinTime    int64  `bson:"join_time,omitempty"`     // 入群时间
	CreatedAt   int64  `bson:"created_at,omitempty"`    // 创建时间
	UpdatedAt   int64  `bson:"updated_at,omitempty"`    // 更新时间
}

type GroupMemberRemove struct {
	UserId    int   // 操作人ID
	GroupId   int   // 群ID
	MemberIds []int // 群成员ID
}

type GroupApply struct {
	gmeta.Meta `collection:"group_apply" bson:"-"`
	GroupId    int    `bson:"group_id"`   // 群聊ID
	UserId     int    `bson:"user_id"`    // 用户ID
	Status     int    `bson:"status"`     // 申请状态
	Remark     string `bson:"remark"`     // 备注信息
	Reason     string `bson:"reason"`     // 拒绝原因
	CreatedAt  int64  `bson:"created_at"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at"` // 更新时间
}

type GroupNotice struct {
	gmeta.Meta   `collection:"group_notice" bson:"-"`
	GroupId      int    `bson:"group_id"`      // 群聊ID
	CreatorId    int    `bson:"creator_id"`    // 创建者用户ID
	Title        string `bson:"title"`         // 公告标题
	Content      string `bson:"content"`       // 公告内容
	ConfirmUsers string `bson:"confirm_users"` // 已确认成员
	IsDelete     int    `bson:"is_delete"`     // 是否删除[0:否;1:是;]
	IsTop        int    `bson:"is_top"`        // 是否置顶[0:否;1:是;]
	IsConfirm    int    `bson:"is_confirm"`    // 是否需群成员确认公告[0:否;1:是;]
	CreatedAt    int64  `bson:"created_at"`    // 创建时间
	UpdatedAt    int64  `bson:"updated_at"`    // 更新时间
	DeletedAt    int64  `bson:"deleted_at"`    // 删除时间
}
