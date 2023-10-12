package entity

import "github.com/iimeta/iim-client/internal/model"

type TalkSession struct {
	Id            string `bson:"_id,omitempty"`         // 聊天列表ID
	TalkType      int    `bson:"talk_type,omitempty"`   // 聊天类型[1:私聊;2:群聊;]
	UserId        int    `bson:"user_id,omitempty"`     // 用户ID
	ReceiverId    int    `bson:"receiver_id,omitempty"` // 接收者ID(用户ID 或 群ID)
	IsTop         int    `bson:"is_top,omitempty"`      // 是否置顶[0:否;1:是;]
	IsDisturb     int    `bson:"is_disturb,omitempty"`  // 消息免打扰[0:否;1:是;]
	IsDelete      int    `bson:"is_delete,omitempty"`   // 是否删除[0:否;1:是;]
	IsRobot       int    `bson:"is_robot,omitempty"`    // 是否机器人[0:否;1:是;]
	CreatedAt     int64  `bson:"created_at,omitempty"`  // 创建时间
	UpdatedAt     int64  `bson:"updated_at,omitempty"`  // 更新时间
	IsTalk        int    `bson:"is_talk,omitempty"`     // 是否允许对话[0:否;1:是;]
	IsOpenContext int    `bson:"is_open_context"`       // 是否开启上下文[0:是;1:否;]
}

type TalkRecords struct {
	Id         string `bson:"_id"`         // ID
	RecordId   int    `bson:"record_id"`   // 记录ID
	MsgId      string `bson:"msg_id"`      // 消息唯一ID
	Sequence   int    `bson:"sequence"`    // 消息时序ID
	TalkType   int    `bson:"talk_type"`   // 对话类型[1:私聊;2:群聊;]
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

	Reply *model.Reply `json:"reply,omitempty" bson:"reply,omitempty"`

	Text     *model.Text     `json:"text,omitempty" bson:"text,omitempty"`
	Code     *model.Code     `json:"code,omitempty" bson:"code,omitempty"`
	Image    *model.Image    `json:"image,omitempty" bson:"image,omitempty"`
	Voice    *model.Voice    `json:"voice,omitempty" bson:"voice,omitempty"`
	Video    *model.Video    `json:"video,omitempty" bson:"video,omitempty"`
	File     *model.File     `json:"file,omitempty" bson:"file,omitempty"`
	Vote     *model.Vote     `json:"vote,omitempty" bson:"vote,omitempty"`
	Mixed    *model.Mixed    `json:"mixed,omitempty" bson:"mixed,omitempty"`
	Emoticon *model.Emoticon `json:"emoticon,omitempty" bson:"emoticon,omitempty"`
	Card     *model.Card     `json:"card,omitempty" bson:"card,omitempty"`
	Location *model.Location `json:"location,omitempty" bson:"location,omitempty"`
}

type TalkRecordsVote struct {
	Id           string `bson:"_id"`           // 投票ID
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
	Id        string `bson:"_id"`        // 答题ID
	VoteId    string `bson:"vote_id"`    // 投票ID
	UserId    int    `bson:"user_id"`    // 用户ID
	Option    string `bson:"option"`     // 投票选项[A、B、C 、D、E、F]
	CreatedAt int64  `bson:"created_at"` // 答题时间
}

type TalkRecordsDelete struct {
	Id        string `bson:"_id"`        // ID
	RecordId  int    `bson:"record_id"`  // 聊天记录ID
	UserId    int    `bson:"user_id"`    // 用户ID
	CreatedAt int64  `bson:"created_at"` // 创建时间
}
