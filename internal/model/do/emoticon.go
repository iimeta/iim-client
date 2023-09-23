package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	EMOTICON_COLLECTION      = "emoticon"
	EMOTICON_ITEM_COLLECTION = "emoticon_item"
	USER_EMOTICON_COLLECTION = "user_emoticon"
)

type Emoticon struct {
	gmeta.Meta `collection:"emoticon" bson:"-"`
	Name       string `bson:"name"`       // 分组名称
	Icon       string `bson:"icon"`       // 分组图标
	Status     int    `bson:"status"`     // 分组状态[-1:已删除;0:正常;1:已禁用;]
	CreatedAt  int64  `bson:"created_at"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at"` // 更新时间
}

type EmoticonItem struct {
	gmeta.Meta `collection:"emoticon_item" bson:"-"`
	EmoticonId string `bson:"emoticon_id"` // 表情分组ID
	UserId     int    `bson:"user_id"`     // 用户ID(0: 代码系统表情包)
	Describe   string `bson:"describe"`    // 表情描述
	Url        string `bson:"url"`         // 图片链接
	FileSuffix string `bson:"file_suffix"` // 文件后缀名
	FileSize   int    `bson:"file_size"`   // 文件大小(单位字节)
	CreatedAt  int64  `bson:"created_at"`  // 创建时间
	UpdatedAt  int64  `bson:"updated_at"`  // 更新时间
}

type UserEmoticon struct {
	gmeta.Meta  `collection:"user_emoticon" bson:"-"`
	UserId      int    `bson:"user_id"`      // 用户ID
	EmoticonIds string `bson:"emoticon_ids"` // 表情包ID
	CreatedAt   int64  `bson:"created_at"`   // 创建时间
}
