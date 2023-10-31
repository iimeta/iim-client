package model

type VipInfo struct {
	UserId      int    `bson:"user_id"`      // 用户ID
	Nickname    string `bson:"nickname"`     // 昵称
	Avatar      string `bson:"avatar"`       // 头像
	SecretKey   string `json:"secret_key"`   // 密钥
	RegTime     string `json:"reg_time"`     // 注册时间
	UsageCount  int    `json:"usage_count"`  // 使用次数
	UsedTokens  int    `json:"used_tokens"`  // 已用Tokens
	TotalTokens int    `json:"total_tokens"` // 总Tokens
}
