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
	IPublish interface {
		// Publish 发送消息接口
		Publish(ctx context.Context, params model.PublishBaseMessageReq) error
	}
)

var (
	localPublish IPublish
)

func Publish() IPublish {
	if localPublish == nil {
		panic("implement not found for interface IPublish, forgot register?")
	}
	return localPublish
}

func RegisterPublish(i IPublish) {
	localPublish = i
}
