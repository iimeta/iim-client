// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/iim-client/utility/socket"
)

type (
	IServerEvent interface {
		Call(ctx context.Context, client socket.IClient, event string, data []byte)
		// 连接成功回调事件
		OnOpen(client socket.IClient)
		// 消息回调事件
		OnMessage(client socket.IClient, message []byte)
		// 连接关闭回调事件
		OnClose(client socket.IClient, code int, text string)
	}
)

var (
	localServerEvent IServerEvent
)

func ServerEvent() IServerEvent {
	if localServerEvent == nil {
		panic("implement not found for interface IServerEvent, forgot register?")
	}
	return localServerEvent
}

func RegisterServerEvent(i IServerEvent) {
	localServerEvent = i
}
