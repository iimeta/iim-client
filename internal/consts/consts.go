package consts

const (
	CHAT_MESSAGES_PREFIX_KEY       = "chat_messages:%d_%d"
	CHAT_MESSAGES_SMART_PREFIX_KEY = "chat_messages:smart:%d_%d"
)

const (
	ContactStatusNormal = 1
	ContactStatusDelete = 0
)

const (
	GroupMemberMaxNum = 200 // 最大成员数量
)

const (
	GroupApplyStatusWait   = 1 // 待处理
	GroupApplyStatusPass   = 2 // 通过
	GroupApplyStatusRefuse = 3 // 拒绝
)

const (
	GroupMemberQuitStatusYes = 1
	GroupMemberQuitStatusNo  = 0

	GroupMemberMuteStatusYes = 1
	GroupMemberMuteStatusNo  = 0
)

const (
	RootStatusDeleted = -1
	RootStatusNormal  = 0
	RootStatusDisable = 1
)

const (
	TalkRecordTalkTypePrivate = 1
	TalkRecordTalkTypeGroup   = 2
)

const (
	VoteAnswerModeSingleChoice   = 0
	VoteAnswerModeMultipleChoice = 1
)
