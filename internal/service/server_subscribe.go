// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"net/http"

	"github.com/iimeta/iim-client/utility/socket"

	"golang.org/x/sync/errgroup"
)

type (
	IServerSubscribe interface {
		// 初始化连接
		Conn(w http.ResponseWriter, r *http.Request) error
		NewClient(uid int, conn socket.IConn) error
		// Start 启动服务
		Start(ctx context.Context, eg *errgroup.Group)
		// 注册健康上报
		SetupHealthSubscribe(ctx context.Context) error
		// 注册消息订阅
		SetupMessageSubscribe(ctx context.Context) error
	}
)

var (
	localServerSubscribe IServerSubscribe
)

func ServerSubscribe() IServerSubscribe {
	if localServerSubscribe == nil {
		panic("implement not found for interface IServerSubscribe, forgot register?")
	}
	return localServerSubscribe
}

func RegisterServerSubscribe(i IServerSubscribe) {
	localServerSubscribe = i
}
