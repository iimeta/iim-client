package model

type ConsumeContactStatus struct {
	Status int `json:"status"`
	UserId int `json:"user_id"`
}

type ConsumeContactApply struct {
	ApplyId string `json:"apply_id"`
	Type    int    `json:"type"`
}

type ConsumeGroupJoin struct {
	Gid  int   `json:"group_id"`
	Type int   `json:"type"`
	Uids []int `json:"uids"`
}

type ConsumeGroupApply struct {
	GroupId int `json:"group_id"`
	UserId  int `json:"user_id"`
}

type ConsumeTalkKeyboard struct {
	SenderID   int `json:"sender_id"`
	ReceiverID int `json:"receiver_id"`
}

type ConsumeTalk struct {
	TalkType   int `json:"talk_type"`
	SenderID   int `json:"sender_id"`
	ReceiverID int `json:"receiver_id"`
	RecordID   int `json:"record_id"`
}

type ConsumeTalkRead struct {
	SenderId   int   `json:"sender_id"`
	ReceiverId int   `json:"receiver_id"`
	Ids        []int `json:"ids"`
}

type ConsumeTalkRevoke struct {
	RecordId int `json:"record_id"`
}
