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
	UserId     int    `bson:"user_id,omitempty"`     // 用户ID
	ClassId    string `bson:"class_id,omitempty"`    // 分类ID
	TagsId     string `bson:"tags_id,omitempty"`     // 笔记关联标签
	Title      string `bson:"title,omitempty"`       // 笔记标题
	Abstract   string `bson:"abstract,omitempty"`    // 笔记摘要
	Image      string `bson:"image,omitempty"`       // 笔记首图
	IsAsterisk int    `bson:"is_asterisk,omitempty"` // 是否星标笔记[0:否;1:是;]
	Status     int    `bson:"status,omitempty"`      // 笔记状态[1:正常;2:已删除;]
	CreatedAt  int64  `bson:"created_at,omitempty"`  // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"`  // 更新时间
	DeletedAt  int64  `bson:"deleted_at,omitempty"`  // 删除时间
}

type NoteDetail struct {
	gmeta.Meta `collection:"note_detail" bson:"-"`
	NoteId     string `bson:"note_id,omitempty"`    // 笔记ID
	MdContent  string `bson:"md_content,omitempty"` // Markdown 内容
	Content    string `bson:"content,omitempty"`    // Markdown 解析HTML内容
}

type NoteClass struct {
	gmeta.Meta `collection:"note_class" bson:"-"`
	UserId     int    `bson:"user_id,omitempty"`    // 用户ID
	ClassName  string `bson:"class_name,omitempty"` // 分类名
	Sort       int    `bson:"sort,omitempty"`       // 排序
	IsDefault  int    `bson:"is_default,omitempty"` // 默认分类[0:否;1:是；]
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}

type NoteTag struct {
	gmeta.Meta `collection:"note_tag" bson:"-"`
	UserId     int    `bson:"user_id,omitempty"`    // 用户ID
	TagName    string `bson:"tag_name,omitempty"`   // 标签名
	Sort       int    `bson:"sort,omitempty"`       // 排序
	CreatedAt  int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64  `bson:"updated_at,omitempty"` // 更新时间
}

type NoteAnnex struct {
	gmeta.Meta   `collection:"note_annex" bson:"-"`
	UserId       int    `bson:"user_id,omitempty"`       // 上传文件的用户ID
	NoteId       string `bson:"note_id,omitempty"`       // 笔记ID
	Drive        int    `bson:"drive,omitempty"`         // 文件驱动[1:local;2:cos;]
	Suffix       string `bson:"suffix,omitempty"`        // 文件后缀名
	Size         int    `bson:"size,omitempty"`          // 文件大小
	Path         string `bson:"path,omitempty"`          // 文件地址(相对地址)
	OriginalName string `bson:"original_name,omitempty"` // 原文件名
	Status       int    `bson:"status,omitempty"`        // 附件状态[1:正常;2:已删除;]
	CreatedAt    int64  `bson:"created_at,omitempty"`    // 创建时间
	UpdatedAt    int64  `bson:"updated_at,omitempty"`    // 更新时间
	DeletedAt    int64  `bson:"deleted_at,omitempty"`    // 删除时间
}

type NoteCreate struct {
	UserId    int
	NoteId    string
	ClassId   string
	Title     string
	Content   string
	MdContent string
}

type NoteEdit struct {
	UserId    int
	NoteId    string
	ClassId   string
	Title     string
	Content   string
	MdContent string
}

type NoteList struct {
	UserId   int
	Keyword  string
	FindType int
	Cid      string
	Page     int
}
