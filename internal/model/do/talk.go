package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	TALK_SESSION_COLLECTION             = "talk_session"
	TALK_RECORDS_COLLECTION             = "talk_records"
	TALK_RECORDS_VOTE_COLLECTION        = "talk_records_vote"
	TALK_RECORDS_VOTE_ANSWER_COLLECTION = "talk_records_vote_answer"
	TALK_RECORDS_DELETE_COLLECTION      = "talk_records_delete"
)

type TalkSession struct {
	gmeta.Meta    `collection:"talk_session" bson:"-"`
	TalkType      int   `bson:"talk_type,omitempty"`   // 聊天类型[1:私信;2:群聊;]
	UserId        int   `bson:"user_id,omitempty"`     // 用户ID
	ReceiverId    int   `bson:"receiver_id,omitempty"` // 接收者ID(用户ID 或 群ID)
	IsTop         int   `bson:"is_top"`                // 是否置顶[0:否;1:是;]
	IsDisturb     int   `bson:"is_disturb"`            // 消息免打扰[0:否;1:是;]
	IsDelete      int   `bson:"is_delete"`             // 是否删除[0:否;1:是;]
	IsRobot       int   `bson:"is_robot"`              // 是否机器人[0:否;1:是;]
	CreatedAt     int64 `bson:"created_at,omitempty"`  // 创建时间
	UpdatedAt     int64 `bson:"updated_at,omitempty"`  // 更新时间
	IsTalk        int   `bson:"is_talk,omitempty"`     // 是否允许对话[0:否;1:是;]
	IsOpenContext int   `bson:"is_open_context"`       // 是否开启上下文[0:是;1:否;]
}

type TalkSessionCreate struct {
	UserId     int `bson:"user_id"`
	TalkType   int `bson:"talk_type"`
	ReceiverId int `bson:"receiver_id"`
	IsRobot    int `bson:"is_robot,omitempty"` // 是否机器人[0:否;1:是;]
	IsTalk     int `bson:"is_talk,omitempty"`  // 是否允许对话[0:否;1:是;]
}

type TalkSessionTop struct {
	UserId int    `bson:"user_id"`
	Id     string `bson:"id"`
	Type   int    `bson:"type"`
}

type TalkSessionDisturb struct {
	UserId     int `bson:"user_id"`
	TalkType   int `bson:"talk_type"`
	ReceiverId int `bson:"receiver_id"`
	IsDisturb  int `bson:"is_disturb"`
}

type TalkSessionOpenContext struct {
	UserId        int `bson:"user_id"`
	TalkType      int `bson:"talk_type"`
	ReceiverId    int `bson:"receiver_id"`
	IsOpenContext int `bson:"is_open_context"`
}

type TalkRecords struct {
	gmeta.Meta `collection:"talk_records" bson:"-"`
	RecordId   int    `bson:"record_id"`   // 记录ID
	MsgId      string `bson:"msg_id"`      // 消息唯一ID
	Sequence   int64  `bson:"sequence"`    // 消息时序ID
	TalkType   int    `bson:"talk_type"`   // 对话类型[1:私信;2:群聊;]
	MsgType    int    `bson:"msg_type"`    // 消息类型
	UserId     int    `bson:"user_id"`     // 发送者ID[0:系统用户;]
	ReceiverId int    `bson:"receiver_id"` // 接收者ID(用户ID 或 群ID)
	IsRevoke   int    `bson:"is_revoke"`   // 是否撤回消息[0:否;1:是;]
	IsMark     int    `bson:"is_mark"`     // 是否重要消息[0:否;1:是;]
	IsRead     int    `bson:"is_read"`     // 是否已读[0:否;1:是;]
	QuoteId    string `bson:"quote_id"`    // 引用消息ID
	Content    string `bson:"content"`     // 文本消息
	Extra      string `bson:"extra"`       // 扩展信信息
	CreatedAt  int64  `bson:"created_at"`  // 创建时间
	UpdatedAt  int64  `bson:"updated_at"`  // 更新时间
}

type TalkRecordsVote struct {
	gmeta.Meta   `collection:"talk_records_vote" bson:"-"`
	RecordId     int    `bson:"record_id"`     // 消息记录ID
	UserId       int    `bson:"user_id"`       // 用户ID
	Title        string `bson:"title"`         // 投票标题
	AnswerMode   int    `bson:"answer_mode"`   // 答题模式[0:单选;1:多选;]
	AnswerOption string `bson:"answer_option"` // 答题选项
	AnswerNum    int    `bson:"answer_num"`    // 应答人数
	AnsweredNum  int    `bson:"answered_num"`  // 已答人数
	IsAnonymous  int    `bson:"is_anonymous"`  // 匿名投票[0:否;1:是;]
	Status       int    `bson:"status"`        // 投票状态[0:投票中;1:已完成;]
	CreatedAt    int64  `bson:"created_at"`    // 创建时间
	UpdatedAt    int64  `bson:"updated_at"`    // 更新时间
}

type TalkRecordsVoteAnswer struct {
	gmeta.Meta `collection:"talk_records_vote_answer" bson:"-"`
	VoteId     string `bson:"vote_id"`    // 投票ID
	UserId     int    `bson:"user_id"`    // 用户ID
	Option     string `bson:"option"`     // 投票选项[A、B、C 、D、E、F]
	CreatedAt  int64  `bson:"created_at"` // 答题时间
}

type RemoveRecord struct {
	UserId     int
	TalkType   int
	ReceiverId int
	RecordIds  string
}

type TalkRecordsDelete struct {
	gmeta.Meta `collection:"talk_records_delete" bson:"-"`
	RecordId   int   `bson:"record_id"`  // 聊天记录ID
	UserId     int   `bson:"user_id"`    // 用户ID
	CreatedAt  int64 `bson:"created_at"` // 创建时间
}

type TalkRecordsQuery struct {
	TalkType   int   // 对话类型
	UserId     int   // 获取消息的用户
	ReceiverId int   // 接收者ID
	MsgType    []int // 消息类型
	RecordId   int   // 上次查询的最小消息ID
	Limit      int   // 数据行数
}

type TalkRecordExtraForward struct {
	TalkType   int              `json:"talk_type"`   // 对话类型
	UserId     int              `json:"user_id"`     // 发送者ID
	ReceiverId int              `json:"receiver_id"` // 接收者ID
	RecordsIds []int            `json:"records_ids"` // 消息列表
	Records    []map[string]any `json:"records"`     // 消息快照
}
