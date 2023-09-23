package entity

type SplitUpload struct {
	Id           string `bson:"_id"`           // 临时文件ID
	Type         int    `bson:"type"`          // 文件属性[1:合并文件;2:拆分文件]
	Drive        int    `bson:"drive"`         // 驱动类型[1:local;2:cos;]
	UploadId     string `bson:"upload_id"`     // 临时文件hash名
	UserId       int    `bson:"user_id"`       // 上传的用户ID
	OriginalName string `bson:"original_name"` // 原文件名
	SplitIndex   int    `bson:"split_index"`   // 当前索引块
	SplitNum     int    `bson:"split_num"`     // 总上传索引块
	Path         string `bson:"path"`          // 临时保存路径
	FileExt      string `bson:"file_ext"`      // 文件后缀名
	FileSize     int64  `bson:"file_size"`     // 文件大小
	IsDelete     int    `bson:"is_delete"`     // 文件是否删除[0:否;1:是;]
	Attr         string `bson:"attr"`          // 额外参数json
	CreatedAt    int64  `bson:"created_at"`    // 更新时间
	UpdatedAt    int64  `bson:"updated_at"`    // 创建时间
}
