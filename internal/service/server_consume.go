// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IServerConsume interface {
		Call(ctx context.Context, event string, data []byte)
	}
)

var (
	localServerConsume IServerConsume
)

func ServerConsume() IServerConsume {
	if localServerConsume == nil {
		panic("implement not found for interface IServerConsume, forgot register?")
	}
	return localServerConsume
}

func RegisterServerConsume(i IServerConsume) {
	localServerConsume = i
}
