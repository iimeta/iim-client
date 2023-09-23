package v1

// 键盘消息
type KeyboardMessage struct {
	Event string                `json:"event,omitempty"` // 事件名
	Data  *KeyboardMessage_Data `json:"data,omitempty"`  // 数据包
}

type KeyboardMessage_Data struct {
	SenderId   int32 `json:"sender_id,omitempty"`
	ReceiverId int32 `json:"receiver_id,omitempty"`
}
