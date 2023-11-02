package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	VIP_COLLECTION = "vip"
)

type Vip struct {
	gmeta.Meta  `collection:"vip" bson:"-"`
	Level       int      `bson:"level,omitempty"`        // 等级
	Name        string   `bson:"name,omitempty"`         // 名称
	Models      []string `bson:"models,omitempty"`       // 模型权限
	FreeTokens  int      `bson:"free_tokens,omitempty"`  // 免费额度
	MinuteLimit int      `bson:"minute_limit,omitempty"` // 分钟限额
	DailyLimit  int      `bson:"daily_limit,omitempty"`  // 每日限额
	Remark      string   `bson:"remark,omitempty"`       // 备注
	Status      int      `bson:"status,omitempty"`       // 状态[1:正常;2:下线;-1:删除]
	CreatedAt   int64    `bson:"created_at,omitempty"`   // 创建时间
	UpdatedAt   int64    `bson:"updated_at,omitempty"`   // 更新时间
}
