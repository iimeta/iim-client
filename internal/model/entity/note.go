package entity

type Note struct {
	Id         string `bson:"_id"`         // 笔记ID
	UserId     int    `bson:"user_id"`     // 用户ID
	ClassId    string `bson:"class_id"`    // 分类ID
	TagsId     string `bson:"tags_id"`     // 笔记关联标签
	Title      string `bson:"title"`       // 笔记标题
	Abstract   string `bson:"abstract"`    // 笔记摘要
	Image      string `bson:"image"`       // 笔记首图
	IsAsterisk int    `bson:"is_asterisk"` // 是否星标笔记[0:否;1:是;]
	Status     int    `bson:"status"`      // 笔记状态[1:正常;2:已删除;]
	CreatedAt  int64  `bson:"created_at"`  // 创建时间
	UpdatedAt  int64  `bson:"updated_at"`  // 更新时间
	DeletedAt  int64  `bson:"deleted_at"`  // 删除时间
}

type NoteDetail struct {
	Id        string `bson:"_id"`        // 笔记详情ID
	NoteId    string `bson:"note_id"`    // 笔记ID
	MdContent string `bson:"md_content"` // Markdown 内容
	Content   string `bson:"content"`    // Markdown 解析HTML内容
}

type NoteClass struct {
	Id        string `bson:"_id"`        // 笔记分类ID
	UserId    int    `bson:"user_id"`    // 用户ID
	ClassName string `bson:"class_name"` // 分类名
	Sort      int    `bson:"sort"`       // 排序
	IsDefault int    `bson:"is_default"` // 默认分类[0:否;1:是；]
	CreatedAt int64  `bson:"created_at"` // 创建时间
	UpdatedAt int64  `bson:"updated_at"` // 更新时间
}

type NoteTag struct {
	Id        string `bson:"_id"`        // 笔记标签ID
	UserId    int    `bson:"user_id"`    // 用户ID
	TagName   string `bson:"tag_name"`   // 标签名
	Sort      int    `bson:"sort"`       // 排序
	CreatedAt int64  `bson:"created_at"` // 创建时间
	UpdatedAt int64  `bson:"updated_at"` // 更新时间
}

type NoteAnnex struct {
	Id           string `bson:"_id"`           // 文件ID
	UserId       int    `bson:"user_id"`       // 上传文件的用户ID
	NoteId       string `bson:"note_id"`       // 笔记ID
	Drive        int    `bson:"drive"`         // 文件驱动[1:local;2:cos;]
	Suffix       string `bson:"suffix"`        // 文件后缀名
	Size         int    `bson:"size"`          // 文件大小
	Path         string `bson:"path"`          // 文件地址(相对地址)
	OriginalName string `bson:"original_name"` // 原文件名
	Status       int    `bson:"status"`        // 附件状态[1:正常;2:已删除;]
	CreatedAt    int64  `bson:"created_at"`    // 创建时间
	UpdatedAt    int64  `bson:"updated_at"`    // 更新时间
	DeletedAt    int64  `bson:"deleted_at"`    // 删除时间
}
