package model

// 头像上传接口响应参数
type UploadAvatarRes struct {
	Avatar string `json:"avatar,omitempty"`
}

// 头像上传接口响应参数
type UploadImageRes struct {
	Src string `json:"src,omitempty"`
}

// 批量上传文件初始化接口请求参数
type UploadInitiateMultipartReq struct {
	FileName string `json:"file_name,omitempty" v:"required"`
	FileSize int64  `json:"file_size,omitempty" v:"required"`
}

// 批量上传文件初始化接口响应参数
type UploadInitiateMultipartRes struct {
	UploadId    string `json:"upload_id,omitempty"`
	SplitSize   int    `json:"split_size,omitempty"`
	UploadIdMd5 string `json:"upload_id_md5,omitempty"`
}

// 批量上传文件接口请求参数
type UploadMultipartReq struct {
	UploadId   string `json:"upload_id,omitempty" v:"required"`
	SplitIndex int    `json:"split_index,omitempty" v:"min:0"`
	SplitNum   int    `json:"split_num,omitempty" v:"required|min:1"`
}

// 批量上传文件接口请求参数
type UploadMultipartRes struct {
	UploadId string `json:"upload_id,omitempty"`
	IsMerge  bool   `json:"is_merge,omitempty"`
}
