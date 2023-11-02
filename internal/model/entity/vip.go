package entity

type Vip struct {
	Id          string `bson:"_id,omitempty"`          // ID
	Level       int    `bson:"level,omitempty"`        // 等级
	Name        string `bson:"name,omitempty"`         // 名称
	MinuteLimit int    `bson:"minute_limit,omitempty"` // 每分钟限额
	DailyLimit  int    `bson:"daily_limit,omitempty"`  // 每日限额
	Remark      string `bson:"remark,omitempty"`       // 备注
	Status      int    `bson:"status,omitempty"`       // 状态[1:正常;2:禁用;-1:删除]
	CreatedAt   int64  `bson:"created_at,omitempty"`   // 创建时间
	UpdatedAt   int64  `bson:"updated_at,omitempty"`   // 更新时间
}
