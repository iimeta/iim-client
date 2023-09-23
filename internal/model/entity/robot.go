package entity

type Robot struct {
	Id        string `bson:"_id,omitempty"`        // 机器人ID
	UserId    int    `bson:"user_id,omitempty"`    // 关联用户ID
	RobotName string `bson:"robot_name,omitempty"` // 机器人名称
	Describe  string `bson:"describe,omitempty"`   // 描述信息
	Logo      string `bson:"logo,omitempty"`       // 机器人logo
	IsTalk    int    `bson:"is_talk,omitempty"`    // 可发送消息[0:否;1:是;]
	Status    int    `bson:"status,omitempty"`     // 状态[-1:已删除;0:正常;1:已禁用;]
	Type      int    `bson:"type,omitempty"`       // 机器人类型
	Company   string `bson:"company,omitempty"`    // 公司
	Model     string `bson:"model,omitempty"`      // 模型
	ModelType string `bson:"model_type,omitempty"` // 模型类型, 聊天: chat, 画图: image
	Role      string `bson:"role,omitempty"`       // 角色
	Prompt    string `bson:"prompt,omitempty"`     // 提示
	MsgType   int    `bson:"msg_type,omitempty"`   // 消息类型[1:文本消息;2:文件消息;3:会话消息;4:代码消息;5:投票消息;6:群公告;7:好友申请;8:登录通知;9:入群消息/退群消息;]
	Proxy     string `bson:"proxy,omitempty"`      // 代理
	CreatedAt int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt int64  `bson:"updated_at,omitempty"` // 更新时间
}
