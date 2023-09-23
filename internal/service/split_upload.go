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
	ISplitUpload interface {
		InitiateMultipartUpload(ctx context.Context, params *model.MultipartInitiateOpt) (string, error)
		MultipartUpload(ctx context.Context, opt *model.MultipartUploadOpt) error
	}
)

var (
	localSplitUpload ISplitUpload
)

func SplitUpload() ISplitUpload {
	if localSplitUpload == nil {
		panic("implement not found for interface ISplitUpload, forgot register?")
	}
	return localSplitUpload
}

func RegisterSplitUpload(i ISplitUpload) {
	localSplitUpload = i
}
