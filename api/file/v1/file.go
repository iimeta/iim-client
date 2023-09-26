package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 头像上传接口请求参数
type UploadAvatarReq struct {
	g.Meta `path:"/avatar" tags:"file" method:"post" summary:"头像上传接口"`
}

// 头像上传接口响应参数
type UploadAvatarRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.UploadAvatarRes
}

// 头像上传接口请求参数
type UploadImageReq struct {
	g.Meta `path:"/image" tags:"file" method:"post" summary:"头像上传接口"`
}

// 头像上传接口响应参数
type UploadImageRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.UploadImageRes
}

// 批量上传文件初始化接口请求参数
type UploadInitiateMultipartReq struct {
	g.Meta `path:"/multipart/initiate" tags:"file" method:"post" summary:"批量上传文件初始化接口"`
	model.UploadInitiateMultipartReq
}

// 批量上传文件初始化接口响应参数
type UploadInitiateMultipartRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.UploadInitiateMultipartRes
}

// 批量上传文件接口请求参数
type UploadMultipartReq struct {
	g.Meta `path:"/multipart" tags:"file" method:"post" summary:"批量上传文件接口"`
	model.UploadMultipartReq
}

// 批量上传文件接口请求参数
type UploadMultipartRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.UploadMultipartRes
}
