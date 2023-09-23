package model

import (
	"database/sql"
	"time"
)

type ArticleListItem struct {
	Id         string `json:"id"`          // 文章ID
	UserId     int    `json:"user_id"`     // 用户ID
	ClassId    string `json:"class_id"`    // 分类ID
	TagsId     string `json:"tags_id"`     // 笔记关联标签
	Title      string `json:"title"`       // 文章标题
	Abstract   string `json:"abstract"`    // 文章摘要
	Image      string `json:"image"`       // 文章首图
	IsAsterisk int    `json:"is_asterisk"` // 是否星标文章[0:否;1:是;]
	Status     int    `json:"status"`      // 笔记状态[1:正常;2:已删除;]
	CreatedAt  int64  `json:"created_at"`  // 创建时间
	UpdatedAt  int64  `json:"updated_at"`  // 更新时间
	ClassName  string `json:"class_name"`  // 分类名
}

type ArticleDetailInfo struct {
	Id         string `json:"id"`          // 文章ID
	UserId     int    `json:"user_id"`     // 用户ID
	ClassId    string `json:"class_id"`    // 分类ID
	TagsId     string `json:"tags_id"`     // 笔记关联标签
	Title      string `json:"title"`       // 文章标题
	Abstract   string `json:"abstract"`    // 文章摘要
	Image      string `json:"image"`       // 文章首图
	IsAsterisk int    `json:"is_asterisk"` // 是否星标文章(0:否  1:是)
	Status     int    `json:"status"`      // 笔记状态 1:正常 2:已删除
	CreatedAt  int64  `json:"created_at"`  // 添加时间
	UpdatedAt  int64  `json:"updated_at"`  // 最后一次更新时间
	MdContent  string `json:"md_content"`  // Markdown 内容
	Content    string `json:"content"`     // Markdown 解析HTML内容
}

type ArticleEditOpt struct {
	UserId    int
	ArticleId string
	ClassId   string
	Title     string
	Content   string
	MdContent string
}

type ArticleListOpt struct {
	UserId   int
	Keyword  string
	FindType int
	Cid      int
	Page     int
}

// 文章编辑接口请求参数
type ArticleEditReq struct {
	ArticleId string `json:"article_id,omitempty"`
	ClassId   string `json:"class_id,omitempty"`
	Title     string `json:"title,omitempty" v:"required"`
	Content   string `json:"content,omitempty" v:"required"`
	MdContent string `json:"md_content,omitempty" v:"required"`
}

// 文章编辑接口响应参数
type ArticleEditRes struct {
	Id       string `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Abstract string `json:"abstract,omitempty"`
	Image    string `json:"image,omitempty"`
}

// 文章详情接口请求参数
type ArticleDetailReq struct {
	ArticleId string `json:"article_id,omitempty" v:"required"`
}

// 文章详情接口响应参数
type ArticleDetailRes struct {
	Id         string `json:"id,omitempty"`
	ClassId    string `json:"class_id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	MdContent  string `json:"md_content,omitempty"`
	IsAsterisk int    `json:"is_asterisk,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	// 标签列表
	Tags []*ArticleDetailResponse_Tag `json:"tags,omitempty"`
	// 附件列表
	Files []*ArticleDetailResponse_File `json:"files,omitempty"`
}

// 文章列表接口请求参数
type ArticleListReq struct {
	Keyword  string `json:"keyword,omitempty"`
	FindType int    `json:"find_type,omitempty"`
	Cid      int    `json:"cid,omitempty"`
	Page     int    `json:"page,omitempty"`
}

// 文章列表请求接口响应参数
type ArticleListRes struct {
	Items    []*ArticleListResponse_Item   `json:"items"`
	Paginate *ArticleListResponse_Paginate `json:"paginate"`
}

// 文章删除接口请求参数
type ArticleDeleteReq struct {
	ArticleId string `json:"article_id,omitempty" v:"required"`
}

// 恢复文章接口请求参数
type ArticleRecoverReq struct {
	ArticleId string `json:"article_id,omitempty" v:"required"`
}

// 文章移动分类接口请求参数
type ArticleMoveReq struct {
	ArticleId string `json:"article_id,omitempty" v:"required,gt=0"`
	ClassId   string `json:"class_id,omitempty" v:"required,gt=0"`
}

// 标记文章接口请求参数
type ArticleAsteriskReq struct {
	ArticleId string `json:"article_id,omitempty" v:"required,gt=0"`
	Type      int    `json:"type,omitempty" v:"required,oneof=1 2"`
}

// 文章标签接口请求参数
type ArticleTagsReq struct {
	ArticleId string `json:"article_id,omitempty" v:"required,gt=0"`
	Tags      []int  `json:"tags,omitempty"`
}

// 永久删除文章接口请求参数
type ArticleForeverDeleteReq struct {
	ArticleId string `json:"article_id,omitempty" v:"required,gt=0"`
}

// 文章图片上传接口请求参数
type ArticleUploadImageReq struct {
	ArticleId string `json:"article_id,omitempty" v:"required,gt=0"`
}

// 文章图片上传接口响应参数
type ArticleUploadImageRes struct {
	Url string `json:"url"`
}

type ArticleDetailResponse_Tag struct {
	Id string `json:"id,omitempty"`
}

type ArticleDetailResponse_File struct {
	Id           string `json:"id,omitempty"`
	OriginalName string `json:"original_name,omitempty"`
	Size         int    `json:"size,omitempty"`
	Suffix       string `json:"suffix,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
}

type ArticleListResponse_Item struct {
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

type ArticleListResponse_Paginate struct {
	Page  int `json:"page,omitempty"`
	Size  int `json:"size,omitempty"`
	Total int `json:"total,omitempty"`
}

type ArticleAnnex struct {
	Id           string       `json:"id"`            // 文件ID
	UserId       int          `json:"user_id"`       // 上传文件的用户ID
	ArticleId    string       `json:"article_id"`    // 笔记ID
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
	ArticleId    string `json:"article_id"`    // 笔记ID
	Title        string `json:"title"`         // 原文件名
	OriginalName string `json:"original_name"` // 原文件名
	DeletedAt    int64  `json:"deleted_at"`    // 附件删除时间
}

// 文章附件上传接口请求参数
type ArticleAnnexUploadReq struct {
	ArticleId string `json:"article_id,omitempty" v:"required"`
}

// 文章附件上传接口响应参数
type ArticleAnnexUploadRes struct {
	Id           string `json:"id,omitempty"`
	Size         int    `json:"size,omitempty"`
	Path         string `json:"path,omitempty"`
	Suffix       string `json:"suffix,omitempty"`
	OriginalName string `json:"original_name,omitempty"`
}

// 文章附件删除接口请求参数
type ArticleAnnexDeleteReq struct {
	AnnexId string `json:"annex_id,omitempty" v:"required"`
}

// 文章附件恢复删除接口请求参数
type ArticleAnnexRecoverReq struct {
	AnnexId string `json:"annex_id,omitempty" v:"required"`
}

// 文章附件回收站列表接口响应参数
type ArticleAnnexRecoverListRes struct {
	Items    []*ArticleAnnexRecoverListResponse_Item `json:"items"`
	Paginate *Paginate                               `json:"paginate,omitempty"`
}

type ArticleAnnexRecoverListResponse_Item struct {
	Id           string `json:"id,omitempty"`
	ArticleId    string `json:"article_id,omitempty"`
	Title        string `json:"title,omitempty"`
	OriginalName string `json:"original_name,omitempty"`
	Day          int    `json:"day,omitempty"`
}

// 文章附件永久删除接口请求参数
type ArticleAnnexForeverDeleteReq struct {
	AnnexId string `json:"annex_id,omitempty" v:"required"`
}

// 文章附件下载接口请求参数
type ArticleAnnexDownloadReq struct {
	AnnexId string `json:"annex_id,omitempty" v:"required"`
}

type ArticleClassItem struct {
	Id        string `json:"id"`         // 文章分类ID
	ClassName string `json:"class_name"` // 分类名
	IsDefault int    `json:"is_default"` // 默认分类1:是 0:不是
	Count     int    `json:"count"`      // 分类名
}

// 文章分类列表接口响应参数
type ArticleClassListRes struct {
	Items    []*ArticleClassListResponse_Item `json:"items"`
	Paginate *Paginate                        `json:"paginate"`
}

// 文章分类编辑接口请求参数
type ArticleClassEditReq struct {
	ClassId   string `json:"class_id,omitempty"`
	ClassName string `json:"class_name,omitempty" v:"required"`
}

// 文章分类编辑接口响应参数
type ArticleClassEditRes struct {
	Id string `json:"id,omitempty"`
}

// 文章分类删除接口请求参数
type ArticleClassDeleteReq struct {
	ClassId string `json:"class_id,omitempty" v:"required"`
}

// 文章分类排序接口请求参数
type ArticleClassSortReq struct {
	ClassId  string `json:"class_id,omitempty" v:"required"`
	SortType int    `json:"sort_type,omitempty" v:"required|in:1,2"`
}

type ArticleClassListResponse_Item struct {
	Id        string `json:"id,omitempty"`
	ClassName string `json:"class_name,omitempty"`
	IsDefault int    `json:"is_default,omitempty"`
	Count     int    `json:"count,omitempty"`
}

type TagItem struct {
	Id      string `json:"id"`       // 文章分类ID
	TagName string `json:"tag_name"` // 标签名
	Count   int    `json:"count"`    // 排序
}

// 文章标签列表接口响应参数
type ArticleTagListRes struct {
	Tags []*ArticleTagListResponse_Item `json:"tags"`
}

// 文章标签编辑接口请求参数
type ArticleTagEditReq struct {
	TagId   string `json:"tag_id,omitempty"`
	TagName string `json:"tag_name,omitempty" v:"required"`
}

// 文章标签编辑接口请求参数
type ArticleTagEditRes struct {
	Id string `json:"id,omitempty"`
}

// 文章标签删除接口请求参数
type ArticleTagDeleteReq struct {
	TagId string `json:"tag_id,omitempty" v:"required"`
}

type ArticleTagListResponse_Item struct {
	Id      string `json:"id,omitempty"`
	TagName string `json:"tag_name,omitempty"`
	Count   int    `json:"count,omitempty"`
}
