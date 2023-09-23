package model

// 联系人分组
type ContactGroup struct {
	Id        string `json:"id"`         // 主键ID
	UserId    int    `json:"user_id"`    // 用户ID
	Name      string `json:"remark"`     // 分组名称
	Num       int    `json:"num"`        // 成员总数
	Sort      int    `json:"sort"`       // 分组名称
	CreatedAt int64  `json:"created_at"` // 创建时间
	UpdatedAt int64  `json:"updated_at"` // 更新时间
}

type ContactGroupListResponse_Item struct {
	// 分组ID
	Id string `json:"id"`
	// 分组名称
	Name string `json:"name"`
	// 联系人数
	Count int `json:"count"`
	// 分组排序
	Sort int `json:"sort"`
}

// 联系人分组列表接口响应参数
type ContactGroupListRes struct {
	// 分组列表
	Items []*ContactGroupListResponse_Item `json:"items"`
}

// 保存联系人分组列表接口请求参数
type ContactGroupSaveReq struct {
	Items []*ContactGroupSaveRequest_Item `json:"items" v:"required"`
}

type ContactGroupSaveRequest_Item struct {
	Id string `json:"id,omitempty" v:"required"`
	//Sort int    `json:"sort,omitempty" v:"required"`
	Name string `json:"name,omitempty" v:"required"`
}
