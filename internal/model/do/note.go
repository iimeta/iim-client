package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	NOTE_COLLECTION        = "note"
	NOTE_DETAIL_COLLECTION = "note_detail"
	NOTE_CLASS_COLLECTION  = "note_class"
	NOTE_TAG_COLLECTION    = "note_tag"
	NOTE_ANNEX_COLLECTION  = "note_annex"
)

type Note struct {
	gmeta.Meta `collection:"note" bson:"-"`
	UserId     int    `bson:"user_id"`     // 用户ID
	ClassId    string `bson:"class_id"`    // 分类ID
	TagsId     string `bson:"tags_id"`     // 笔记关联标签
	Title      string `bson:"title"`       // 文章标题
	Abstract   string `bson:"abstract"`    // 文章摘要
	Image      string `bson:"image"`       // 文章首图
	IsAsterisk int    `bson:"is_asterisk"` // 是否星标文章[0:否;1:是;]
	Status     int    `bson:"status"`      // 笔记状态[1:正常;2:已删除;]
	CreatedAt  int64  `bson:"created_at"`  // 创建时间
	UpdatedAt  int64  `bson:"updated_at"`  // 更新时间
	DeletedAt  int64  `bson:"deleted_at"`  // 删除时间
}

type NoteDetail struct {
	gmeta.Meta `collection:"note_detail" bson:"-"`
	ArticleId  string `bson:"article_id"` // 文章ID
	MdContent  string `bson:"md_content"` // Markdown 内容
	Content    string `bson:"content"`    // Markdown 解析HTML内容
}

type NoteClass struct {
	gmeta.Meta `collection:"note_class" bson:"-"`
	UserId     int    `bson:"user_id"`    // 用户ID
	ClassName  string `bson:"class_name"` // 分类名
	Sort       int    `bson:"sort"`       // 排序
	IsDefault  int    `bson:"is_default"` // 默认分类[0:否;1:是；]
	CreatedAt  int64  `bson:"created_at"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at"` // 更新时间
}

type NoteTag struct {
	gmeta.Meta `collection:"note_tag" bson:"-"`
	UserId     int    `bson:"user_id"`    // 用户ID
	TagName    string `bson:"tag_name"`   // 标签名
	Sort       int    `bson:"sort"`       // 排序
	CreatedAt  int64  `bson:"created_at"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at"` // 更新时间
}

type NoteAnnex struct {
	gmeta.Meta   `collection:"note_annex" bson:"-"`
	UserId       int    `bson:"user_id"`       // 上传文件的用户ID
	ArticleId    string `bson:"article_id"`    // 笔记ID
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

type NoteCreate struct {
	UserId    int
	ArticleId string
	ClassId   string
	Title     string
	Content   string
	MdContent string
}

type NoteEdit struct {
	UserId    int
	ArticleId string
	ClassId   string
	Title     string
	Content   string
	MdContent string
}

type NoteList struct {
	UserId   int
	Keyword  string
	FindType int
	Cid      int
	Page     int
}
