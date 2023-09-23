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
	ISession interface {
		// 获取会话中UserId
		GetUid(ctx context.Context) int
		// 获取会话中用户信息
		GetUser(ctx context.Context) *model.User
		// Create 创建会话列表
		Create(ctx context.Context, params model.TalkSessionCreateReq) (*model.TalkSessionCreateRes, error)
		// Delete 删除列表
		Delete(ctx context.Context, params model.TalkSessionDeleteReq) error
		// Top 置顶列表
		Top(ctx context.Context, params model.TalkSessionTopReq) error
		// Disturb 会话免打扰
		Disturb(ctx context.Context, params model.TalkSessionDisturbReq) error
		// List 会话列表
		List(ctx context.Context) (*model.TalkSessionListRes, error)
		ClearUnreadMessage(ctx context.Context, params model.TalkSessionClearUnreadNumReq) error
	}
)

var (
	localSession ISession
)

func Session() ISession {
	if localSession == nil {
		panic("implement not found for interface ISession, forgot register?")
	}
	return localSession
}

func RegisterSession(i ISession) {
	localSession = i
}
