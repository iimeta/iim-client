package model

import (
	"mime/multipart"
)

type MultipartInitiateOpt struct {
	UserId int
	Name   string
	Size   int64
}

type MultipartUploadOpt struct {
	UserId     int
	UploadId   string
	SplitIndex int
	SplitNum   int
	File       *multipart.FileHeader
}
