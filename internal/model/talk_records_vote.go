package model

const (
	VoteAnswerModeSingleChoice   = 0
	VoteAnswerModeMultipleChoice = 1
)

type QueryVoteModel struct {
	RecordId     int    `json:"record_id"`
	ReceiverId   int    `json:"receiver_id"`
	TalkType     int    `json:"talk_type"`
	MsgType      int    `json:"msg_type"`
	VoteId       string `json:"vote_id"`
	AnswerMode   int    `json:"answer_mode"`
	AnswerOption string `json:"answer_option"`
	AnswerNum    int    `json:"answer_num"`
	VoteStatus   int    `json:"vote_status"`
}
