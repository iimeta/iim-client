package model

import (
	"database/sql"
	"time"
)

type NoteListItem struct {
	Id         string `json:"id"`          // 笔记ID
	UserId     int    `json:"user_id"`     // 用户ID
	ClassId    string `json:"class_id"`    // 分类ID
	TagsId     string `json:"tags_id"`     // 笔记关联标签
	Title      string `json:"title"`       // 笔记标题
	Abstract   string `json:"abstract"`    // 笔记摘要
	Image      string `json:"image"`       // 笔记首图
	IsAsterisk int    `json:"is_asterisk"` // 是否星标笔记[0:否;1:是;]
	Status     int    `json:"status"`      // 笔记状态[1:正常;2:已删除;]
	CreatedAt  int64  `json:"created_at"`  // 创建时间
	UpdatedAt  int64  `json:"updated_at"`  // 更新时间
	ClassName  string `json:"class_name"`  // 分类名
}

type NoteDetailInfo struct {
	Id         string `json:"id"`          // 笔记ID
	UserId     int    `json:"user_id"`     // 用户ID
	ClassId    string `json:"class_id"`    // 分类ID
	TagsId     string `json:"tags_id"`     // 笔记关联标签
	Title      string `json:"title"`       // 笔记标题
	Abstract   string `json:"abstract"`    // 笔记摘要
	Image      string `json:"image"`       // 笔记首图
	IsAsterisk int    `json:"is_asterisk"` // 是否星标笔记(0:否  1:是)
	Status     int    `json:"status"`      // 笔记状态 1:正常 2:已删除
	CreatedAt  int64  `json:"created_at"`  // 添加时间
	UpdatedAt  int64  `json:"updated_at"`  // 最后一次更新时间
	MdContent  string `json:"md_content"`  // Markdown 内容
	Content    string `json:"content"`     // Markdown 解析HTML内容
}

// 笔记编辑接口请求参数
type NoteEditReq struct {
	NoteId    string `json:"note_id,omitempty"`
	ClassId   string `json:"class_id,omitempty"`
	Title     string `json:"title,omitempty" v:"required"`
	Content   string `json:"content,omitempty" v:"required"`
	MdContent string `json:"md_content,omitempty" v:"required"`
}

// 笔记编辑接口响应参数
type NoteEditRes struct {
	Id       string `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Abstract string `json:"abstract,omitempty"`
	Image    string `json:"image,omitempty"`
}

// 笔记详情接口请求参数
type NoteDetailReq struct {
	NoteId string `json:"note_id,omitempty" v:"required"`
}

// 笔记详情接口响应参数
type NoteDetailRes struct {
	Id         string                 `json:"id,omitempty"`
	ClassId    string                 `json:"class_id,omitempty"`
	Title      string                 `json:"title,omitempty"`
	Content    string                 `json:"content,omitempty"`
	MdContent  string                 `json:"md_content,omitempty"`
	IsAsterisk int                    `json:"is_asterisk,omitempty"`
	CreatedAt  string                 `json:"created_at,omitempty"`
	UpdatedAt  string                 `json:"updated_at,omitempty"`
	Tags       []*DetailResponse_Tag  `json:"tags,omitempty"`  // 标签列表
	Files      []*DetailResponse_File `json:"files,omitempty"` // 附件列表
}

// 笔记列表接口请求参数
type NoteListReq struct {
	Keyword  string `json:"keyword,omitempty"`
	FindType int    `json:"find_type,omitempty"`
	Cid      string `json:"cid,omitempty"`
	Page     int    `json:"page,omitempty"`
}

// 笔记列表请求接口响应参数
type NoteListRes struct {
	Items    []*ListResponse_Item   `json:"items"`
	Paginate *ListResponse_Paginate `json:"paginate"`
}

// 笔记删除接口请求参数
type NoteDeleteReq struct {
	NoteId string `json:"note_id,omitempty" v:"required"`
}

// 恢复笔记接口请求参数
type NoteRecoverReq struct {
	NoteId string `json:"note_id,omitempty" v:"required"`
}

// 笔记移动分类接口请求参数
type NoteMoveReq struct {
	NoteId  string `json:"note_id,omitempty" v:"required"`
	ClassId string `json:"class_id,omitempty" v:"required"`
}

// 标记笔记接口请求参数
type NoteAsteriskReq struct {
	NoteId string `json:"note_id,omitempty" v:"required"`
	Type   int    `json:"type,omitempty" v:"required|in:1,2"`
}

// 笔记标签接口请求参数
type NoteTagsReq struct {
	NoteId string `json:"note_id,omitempty" v:"required"`
	Tags   []int  `json:"tags,omitempty"`
}

// 永久删除笔记接口请求参数
type NoteForeverDeleteReq struct {
	NoteId string `json:"note_id,omitempty" v:"required"`
}

// 笔记图片上传接口响应参数
type NoteUploadImageRes struct {
	Url string `json:"url"`
}

type DetailResponse_Tag struct {
	Id string `json:"id,omitempty"`
}

type DetailResponse_File struct {
	Id           string `json:"id,omitempty"`
	OriginalName string `json:"original_name,omitempty"`
	Size         int    `json:"size,omitempty"`
	Suffix       string `json:"suffix,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
}

type ListResponse_Item struct {
	Id         string `json:"id,omitempty"`
	ClassId    string `json:"class_id,omitempty"`
	TagsId     string `json:"tags_id,omitempty"`
	Title      string `json:"title,omitempty"`
	ClassName  string `json:"class_name,omitempty"`
	Image      string `json:"image,omitempty"`
	IsAsterisk int    `json:"is_asterisk,omitempty"`
	Status     int    `json:"status,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	Abstract   string `json:"abstract,omitempty"`
}

type ListResponse_Paginate struct {
	Page  int `json:"page,omitempty"`
	Size  int `json:"size,omitempty"`
	Total int `json:"total,omitempty"`
}

type Annex struct {
	Id           string       `json:"id"`            // 文件ID
	UserId       int          `json:"user_id"`       // 上传文件的用户ID
	NoteId       string       `json:"note_id"`       // 笔记ID
	Drive        int          `json:"drive"`         // 文件驱动[1:local;2:cos;]
	Suffix       string       `json:"suffix"`        // 文件后缀名
	Size         int          `json:"size"`          // 文件大小
	Path         string       `json:"path"`          // 文件地址(相对地址)
	OriginalName string       `json:"original_name"` // 原文件名
	Status       int          `json:"status"`        // 附件状态[1:正常;2:已删除;]
	CreatedAt    time.Time    `json:"created_at"`    // 创建时间
	UpdatedAt    time.Time    `json:"updated_at"`    // 更新时间
	DeletedAt    sql.NullTime `json:"deleted_at"`    // 删除时间
}

type RecoverAnnexItem struct {
	Id           string `json:"id"`            // 文件ID
	NoteId       string `json:"note_id"`       // 笔记ID
	Title        string `json:"title"`         // 原文件名
	OriginalName string `json:"original_name"` // 原文件名
	DeletedAt    int64  `json:"deleted_at"`    // 附件删除时间
}

// 笔记附件上传接口请求参数
type AnnexUploadReq struct {
	NoteId string `json:"note_id,omitempty" v:"required"`
}

// 笔记附件上传接口响应参数
type AnnexUploadRes struct {
	Id           string `json:"id,omitempty"`
	Size         int    `json:"size,omitempty"`
	Path         string `json:"path,omitempty"`
	Suffix       string `json:"suffix,omitempty"`
	OriginalName string `json:"original_name,omitempty"`
}

// 笔记附件删除接口请求参数
type AnnexDeleteReq struct {
	AnnexId string `json:"annex_id,omitempty" v:"required"`
}

// 笔记附件恢复删除接口请求参数
type AnnexRecoverReq struct {
	AnnexId string `json:"annex_id,omitempty" v:"required"`
}

// 笔记附件回收站列表接口响应参数
type AnnexRecoverListRes struct {
	Items    []*AnnexRecoverListResponse_Item `json:"items"`
	Paginate *Paginate                        `json:"paginate,omitempty"`
}

type AnnexRecoverListResponse_Item struct {
	Id           string `json:"id,omitempty"`
	NoteId       string `json:"note_id,omitempty"`
	Title        string `json:"title,omitempty"`
	OriginalName string `json:"original_name,omitempty"`
	Day          int    `json:"day,omitempty"`
}

// 笔记附件永久删除接口请求参数
type AnnexForeverDeleteReq struct {
	AnnexId string `json:"annex_id,omitempty" v:"required"`
}

// 笔记附件下载接口请求参数
type AnnexDownloadReq struct {
	AnnexId string `json:"annex_id,omitempty" v:"required"`
}

type ClassItem struct {
	Id        string `json:"id"`         // 笔记分类ID
	ClassName string `json:"class_name"` // 分类名
	IsDefault int    `json:"is_default"` // 默认分类1:是 0:不是
	Count     int    `json:"count"`      // 分类名
}

// 笔记分类列表接口响应参数
type ClassListRes struct {
	Items    []*ClassListResponse_Item `json:"items"`
	Paginate *Paginate                 `json:"paginate"`
}

// 笔记分类编辑接口请求参数
type ClassEditReq struct {
	ClassId   string `json:"class_id,omitempty"`
	ClassName string `json:"class_name,omitempty" v:"required"`
}

// 笔记分类编辑接口响应参数
type ClassEditRes struct {
	Id string `json:"id,omitempty"`
}

// 笔记分类删除接口请求参数
type ClassDeleteReq struct {
	ClassId string `json:"class_id,omitempty" v:"required"`
}

// 笔记分类排序接口请求参数
type ClassSortReq struct {
	ClassId  string `json:"class_id,omitempty" v:"required"`
	SortType int    `json:"sort_type,omitempty" v:"required|in:1,2"`
}

type ClassListResponse_Item struct {
	Id        string `json:"id,omitempty"`
	ClassName string `json:"class_name,omitempty"`
	IsDefault int    `json:"is_default,omitempty"`
	Count     int    `json:"count,omitempty"`
}

type TagItem struct {
	Id      string `json:"id"`       // 笔记分类ID
	TagName string `json:"tag_name"` // 标签名
	Count   int    `json:"count"`    // 排序
}

// 笔记标签列表接口响应参数
type TagListRes struct {
	Tags []*TagListResponse_Item `json:"tags"`
}

// 笔记标签编辑接口请求参数
type TagEditReq struct {
	TagId   string `json:"tag_id,omitempty"`
	TagName string `json:"tag_name,omitempty" v:"required"`
}

// 笔记标签编辑接口请求参数
type TagEditRes struct {
	Id string `json:"id,omitempty"`
}

// 笔记标签删除接口请求参数
type TagDeleteReq struct {
	TagId string `json:"tag_id,omitempty" v:"required"`
}

type TagListResponse_Item struct {
	Id      string `json:"id,omitempty"`
	TagName string `json:"tag_name,omitempty"`
	Count   int    `json:"count,omitempty"`
}
