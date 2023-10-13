package model

type SysMessage struct {
	MsgId    string    `json:"msg_id,omitempty" bson:"msg_id,omitempty"`
	MsgType  string    `json:"msg_type,omitempty" bson:"msg_type,omitempty"`
	TalkType int       `json:"talk_type,omitempty" bson:"talk_type,omitempty"` // 对话类型 1:私聊 2:群聊
	Sender   *Sender   `json:"sender,omitempty" bson:"sender,omitempty"`       // 发送者
	Receiver *Receiver `json:"receiver,omitempty" bson:"receiver,omitempty"`   // 接收者

	Text                   *Text                   `json:"text,omitempty" bson:"text,omitempty"`
	GroupCreate            *GroupCreate            `json:"group_create,omitempty" bson:"group_create,omitempty"`
	GroupJoin              *GroupJoin              `json:"group_join,omitempty" bson:"group_join,omitempty"`
	GroupTransfer          *GroupTransfer          `json:"group_transfer,omitempty" bson:"group_transfer,omitempty"`
	GroupMuted             *GroupMuted             `json:"group_muted,omitempty" bson:"group_muted,omitempty"`
	GroupCancelMuted       *GroupCancelMuted       `json:"group_cancel_muted,omitempty" bson:"group_cancel_muted,omitempty"`
	GroupMemberMuted       *GroupMemberMuted       `json:"group_member_muted,omitempty" bson:"group_member_muted,omitempty"`
	GroupMemberCancelMuted *GroupMemberCancelMuted `json:"group_member_cancel_muted,omitempty" bson:"group_member_cancel_muted,omitempty"`
	GroupDismissed         *GroupDismissed         `json:"group_dismissed,omitempty" bson:"group_dismissed,omitempty"`
	GroupMemberQuit        *GroupMemberQuit        `json:"group_member_quit,omitempty" bson:"group_member_quit,omitempty"`
	GroupMemberKicked      *GroupMemberKicked      `json:"group_member_kicked,omitempty" bson:"group_member_kicked,omitempty"`
	GroupMessageRevoke     *GroupMessageRevoke     `json:"group_message_revoke,omitempty" bson:"group_message_revoke,omitempty"`
	GroupNotice            *GroupNotice            `json:"group_notice,omitempty" bson:"group_notice,omitempty"`
}

// 创建群消息
type GroupCreate struct {
	OwnerId   int                `json:"owner_id,omitempty" bson:"owner_id,omitempty"`     // 操作人ID
	OwnerName string             `json:"owner_name,omitempty" bson:"owner_name,omitempty"` // 操作人昵称
	Members   []*TalkGroupMember `json:"members,omitempty" bson:"members,omitempty"`       // 成员列表
}

// 群主邀请加入群消息
type GroupJoin struct {
	OwnerId   int                `json:"owner_id,omitempty" bson:"owner_id,omitempty"`     // 操作人ID
	OwnerName string             `json:"owner_name,omitempty" bson:"owner_name,omitempty"` // 操作人昵称
	Members   []*TalkGroupMember `json:"members,omitempty" bson:"members,omitempty"`       // 成员列表
}

// 群主转让群消息
type GroupTransfer struct {
	OldOwnerId   int    `json:"old_owner_id,omitempty" bson:"old_owner_id,omitempty"`     // 老群主ID
	OldOwnerName string `json:"old_owner_name,omitempty" bson:"old_owner_name,omitempty"` // 老群主昵称
	NewOwnerId   int    `json:"new_owner_id,omitempty" bson:"new_owner_id,omitempty"`     // 新群主ID
	NewOwnerName string `json:"new_owner_name,omitempty" bson:"new_owner_name,omitempty"` // 新群主昵称
}

// 管理员设置群禁言消息
type GroupMuted struct {
	OwnerId   int    `json:"owner_id,omitempty" bson:"owner_id,omitempty"`     // 操作人ID
	OwnerName string `json:"owner_name,omitempty" bson:"owner_name,omitempty"` // 操作人昵称
}

// 管理员解除群禁言消息
type GroupCancelMuted struct {
	OwnerId   int    `json:"owner_id,omitempty" bson:"owner_id,omitempty"`     // 操作人ID
	OwnerName string `json:"owner_name,omitempty" bson:"owner_name,omitempty"` // 操作人昵称
}

// 管理员设置群成员禁言消息
type GroupMemberMuted struct {
	OwnerId   int                `json:"owner_id,omitempty" bson:"owner_id,omitempty"`     // 操作人ID
	OwnerName string             `json:"owner_name,omitempty" bson:"owner_name,omitempty"` // 操作人昵称
	Members   []*TalkGroupMember `json:"members,omitempty" bson:"members,omitempty"`       // 成员列表
}

// 管理员解除群成员禁言消息
type GroupMemberCancelMuted struct {
	OwnerId   int                `json:"owner_id,omitempty" bson:"owner_id,omitempty"`     // 操作人ID
	OwnerName string             `json:"owner_name,omitempty" bson:"owner_name,omitempty"` // 操作人昵称
	Members   []*TalkGroupMember `json:"members,omitempty" bson:"members,omitempty"`       // 成员列表
}

// 群主解散群消息
type GroupDismissed struct {
	OwnerId   int    `json:"owner_id,omitempty" bson:"owner_id,omitempty"`     // 操作人ID
	OwnerName string `json:"owner_name,omitempty" bson:"owner_name,omitempty"` // 操作人昵称
}

// 群成员退出群消息
type GroupMemberQuit struct {
	OwnerId   int    `json:"owner_id,omitempty" bson:"owner_id,omitempty"`     // 操作人ID
	OwnerName string `json:"owner_name,omitempty" bson:"owner_name,omitempty"` // 操作人昵称
}

// 踢出群成员消息
type GroupMemberKicked struct {
	OwnerId   int                `json:"owner_id,omitempty" bson:"owner_id,omitempty"`     // 操作人ID
	OwnerName string             `json:"owner_name,omitempty" bson:"owner_name,omitempty"` // 操作人昵称
	Members   []*TalkGroupMember `json:"members,omitempty" bson:"members,omitempty"`       // 成员列表
}

// 管理员撤回成员消息
type GroupMessageRevoke struct {
	OwnerId         int    `json:"owner_id,omitempty" bson:"owner_id,omitempty"`                   // 操作人ID
	OwnerName       string `json:"owner_name,omitempty" bson:"owner_name,omitempty"`               // 操作人昵称
	RevokeMessageId string `json:"revoke_message_id,omitempty" bson:"revoke_message_id,omitempty"` // 被撤回消息ID
}

// 发布群公告
type GroupNotice struct {
	OwnerId   int    `json:"owner_id,omitempty" bson:"owner_id,omitempty"`     // 操作人ID
	OwnerName string `json:"owner_name,omitempty" bson:"owner_name,omitempty"` // 操作人昵称
	Title     string `json:"title,omitempty" bson:"title,omitempty"`           // 标题
	Content   string `json:"content,omitempty" bson:"content,omitempty"`       // 内容
}
