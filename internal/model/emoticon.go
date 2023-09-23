package model

// 用户表情包列表接口响应参数
type EmoticonListRes struct {
	SysEmoticon     []*EmoticonListResponse_SysEmoticon `json:"sys_emoticon,omitempty"`
	CollectEmoticon []*EmoticonListItem                 `json:"collect_emoticon,omitempty"`
}

type EmoticonSysListResponse_Item struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Icon   string `json:"icon,omitempty"`
	Status int    `json:"status,omitempty"`
}

type EmoticonListResponse_SysEmoticon struct {
	EmoticonId string              `json:"emoticon_id,omitempty"`
	Url        string              `json:"url,omitempty"`
	Name       string              `json:"name,omitempty"`
	List       []*EmoticonListItem `json:"list,omitempty"`
}

type EmoticonListItem struct {
	MediaId string `json:"media_id,omitempty"`
	Src     string `json:"src,omitempty"`
}

// 系统表情包列表接口响应参数
type EmoticonSysListRes struct {
	Items []*EmoticonSysListResponse_Item `json:"items"`
}

// 添加或移出表情包接口响应参数
type EmoticonSetSystemRes struct {
	EmoticonId string              `json:"emoticon_id,omitempty"`
	Url        string              `json:"url,omitempty"`
	Name       string              `json:"name,omitempty"`
	List       []*EmoticonListItem `json:"list,omitempty"`
}

// 删除表情包接口请求参数
type EmoticonDeleteReq struct {
	Ids string `json:"ids,omitempty" v:"required"`
}

// 表情包上传接口响应参数
type EmoticonUploadRes struct {
	MediaId string `json:"media_id,omitempty"`
	Src     string `json:"src,omitempty"`
}

// 添加或移出表情包接口请求参数
type EmoticonSetSystemReq struct {
	EmoticonId string `json:"emoticon_id,omitempty" v:"required"`
	Type       int    `json:"type,omitempty" v:"required|in:1,2"`
}
