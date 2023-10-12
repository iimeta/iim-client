package consts

const (
	MsgTypeText     = "text"     // 文本消息
	MsgTypeCode     = "code"     // 代码消息
	MsgTypeImage    = "image"    // 图片文件
	MsgTypeVoice    = "voice"    // 语音文件
	MsgTypeVideo    = "video"    // 视频文件
	MsgTypeFile     = "file"     // 其它文件
	MsgTypeForward  = "forward"  // 转发消息
	MsgTypeVote     = "vote"     // 投票消息
	MsgTypeMixed    = "mixed"    // 图文消息
	MsgTypeEmoticon = "emoticon" // 表情消息
	MsgTypeCard     = "card"     // 名片消息
	MsgTypeLocation = "location" // 位置消息
	MsgTypeLogin    = "login"    // 登录消息
)

const (
	MsgSysText                   = "sys_text"                      // 系统文本消息
	MsgSysGroupCreate            = "sys_group_create"              // 创建群聊消息
	MsgSysGroupMemberJoin        = "sys_group_member_join"         // 加入群聊消息
	MsgSysGroupMemberQuit        = "sys_group_member_quit"         // 群成员退出群消息
	MsgSysGroupMemberKicked      = "sys_group_member_kicked"       // 踢出群成员消息
	MsgSysGroupMessageRevoke     = "sys_group_message_revoke"      // 管理员撤回成员消息
	MsgSysGroupDismissed         = "sys_group_dismissed"           // 群解散
	MsgSysGroupMuted             = "sys_group_muted"               // 群禁言
	MsgSysGroupCancelMuted       = "sys_group_cancel_muted"        // 群解除禁言
	MsgSysGroupMemberMuted       = "sys_group_member_muted"        // 群成员禁言
	MsgSysGroupMemberCancelMuted = "sys_group_member_cancel_muted" // 群成员解除禁言
	MsgSysGroupNotice            = "sys_group_notice"              // 编辑群公告
	MsgSysGroupTransfer          = "sys_group_transfer"            // 变更群主
)
