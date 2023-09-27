package model

// 用户表情包列表接口响应参数
type ListRes struct {
	CollectEmoticons []*CollectEmoticon `json:"collect_emoticon,omitempty"`
}

// 删除表情包接口请求参数
type DeleteReq struct {
	Ids string `json:"ids,omitempty" v:"required"`
}

// 表情包上传接口响应参数
type UploadRes struct {
	MediaId string `json:"media_id,omitempty"`
	Src     string `json:"src,omitempty"`
}

type CollectEmoticon struct {
	MediaId string `json:"media_id,omitempty"`
	Src     string `json:"src,omitempty"`
}
