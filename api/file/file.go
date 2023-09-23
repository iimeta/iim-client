// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package file

import (
	"context"

	"github.com/iimeta/iim-client/api/file/v1"
)

type IFileV1 interface {
	UploadAvatar(ctx context.Context, req *v1.UploadAvatarReq) (res *v1.UploadAvatarRes, err error)
	UploadImage(ctx context.Context, req *v1.UploadImageReq) (res *v1.UploadImageRes, err error)
	UploadInitiateMultipart(ctx context.Context, req *v1.UploadInitiateMultipartReq) (res *v1.UploadInitiateMultipartRes, err error)
	UploadMultipart(ctx context.Context, req *v1.UploadMultipartReq) (res *v1.UploadMultipartRes, err error)
}
