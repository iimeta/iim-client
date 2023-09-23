package model

type Paginate struct {
	Page  int `json:"page,omitempty"`
	Size  int `json:"size,omitempty"`
	Total int `json:"total,omitempty"`
}
