package model

type GetTalkRecordsReq struct {
	TalkType   int `json:"talk_type" v:"required|in:1,2"`  // 对话类型
	MsgType    int `json:"msg_type"`                       // 消息类型
	ReceiverId int `json:"receiver_id" v:"required|min:1"` // 接收者ID
	RecordId   int `json:"record_id"`                      // 上次查询的最小消息ID
	Limit      int `json:"limit" v:"required|max:100"`     // 数据行数
}

type GetTalkRecordsRes struct {
	Limit    int                `json:"limit"`
	RecordId int                `json:"record_id"`
	Items    []*TalkRecordsItem `json:"items"`
}

type GetForwardTalkRecordReq struct {
	RecordId int `json:"record_id"` // 上次查询的最小消息ID
}

type DownloadChatFileReq struct {
	RecordId int `json:"cr_id" v:"cr_id@required|min:1"`
}
