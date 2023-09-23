// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/iim-client/internal/model"
)

type (
	IFile interface {
		// 头像上传上传
		Avatar(ctx context.Context) (*model.UploadAvatarRes, error)
		// 图片上传
		Image(ctx context.Context) (*model.UploadImageRes, error)
		// 批量上传初始化
		InitiateMultipart(ctx context.Context, params model.UploadInitiateMultipartReq) (*model.UploadInitiateMultipartRes, error)
		// MultipartUpload 批量分片上传
		MultipartUpload(ctx context.Context, params model.UploadMultipartReq) (*model.UploadMultipartRes, error)
	}
)

var (
	localFile IFile
)

func File() IFile {
	if localFile == nil {
		panic("implement not found for interface IFile, forgot register?")
	}
	return localFile
}

func RegisterFile(i IFile) {
	localFile = i
}
