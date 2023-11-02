package model

// 登录用户详情接口响应参数
type VipsRes struct {
	Items []*Vip `json:"items"`
}

type Vip struct {
	Level       int      `json:"level,omitempty"`        // 等级
	Name        string   `json:"name,omitempty"`         // 名称
	Models      []string `json:"models,omitempty"`       // 模型权限
	FreeTokens  int      `json:"free_tokens,omitempty"`  // 免费额度
	MinuteLimit int      `json:"minute_limit,omitempty"` // 分钟限额
	DailyLimit  int      `json:"daily_limit,omitempty"`  // 每日限额
	Remark      string   `json:"remark,omitempty"`       // 备注
	Status      int      `json:"status,omitempty"`       // 状态[1:正常;2:下线;-1:删除]
	CreatedAt   int64    `json:"created_at,omitempty"`   // 创建时间
	UpdatedAt   int64    `json:"updated_at,omitempty"`   // 更新时间
}

type VipInfo struct {
	UserId      int    `json:"user_id"`      // 用户ID
	Nickname    string `json:"nickname"`     // 昵称
	Avatar      string `json:"avatar"`       // 头像
	SecretKey   string `json:"secret_key"`   // 密钥
	RegTime     string `json:"reg_time"`     // 注册时间
	UsageCount  int    `json:"usage_count"`  // 使用次数
	UsedTokens  int    `json:"used_tokens"`  // 已用Tokens
	TotalTokens int    `json:"total_tokens"` // 总Tokens
}
