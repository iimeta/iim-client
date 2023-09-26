package model

// 用户表情包列表接口响应参数
type ListRes struct {
	SysEmoticon     []*ListResponse_SysEmoticon `json:"sys_emoticon,omitempty"`
	CollectEmoticon []*ListItem                 `json:"collect_emoticon,omitempty"`
}

type SysListResponse_Item struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Icon   string `json:"icon,omitempty"`
	Status int    `json:"status,omitempty"`
}

type ListResponse_SysEmoticon struct {
	EmoticonId string      `json:"emoticon_id,omitempty"`
	Url        string      `json:"url,omitempty"`
	Name       string      `json:"name,omitempty"`
	List       []*ListItem `json:"list,omitempty"`
}

type ListItem struct {
	MediaId string `json:"media_id,omitempty"`
	Src     string `json:"src,omitempty"`
}

// 系统表情包列表接口响应参数
type SysListRes struct {
	Items []*SysListResponse_Item `json:"items"`
}

// 添加或移出表情包接口响应参数
type SetSystemRes struct {
	EmoticonId string      `json:"emoticon_id,omitempty"`
	Url        string      `json:"url,omitempty"`
	Name       string      `json:"name,omitempty"`
	List       []*ListItem `json:"list,omitempty"`
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

// 添加或移出表情包接口请求参数
type SetSystemReq struct {
	EmoticonId string `json:"emoticon_id,omitempty" v:"required"`
	Type       int    `json:"type,omitempty" v:"required|in:1,2"`
}
