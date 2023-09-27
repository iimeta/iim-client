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
	ITalkSession interface {
		// 创建会话列表
		Create(ctx context.Context, params model.SessionCreateReq) (*model.SessionCreateRes, error)
		// 删除列表
		Delete(ctx context.Context, params model.SessionDeleteReq) error
		// 置顶列表
		Top(ctx context.Context, params model.SessionTopReq) error
		// 会话免打扰
		Disturb(ctx context.Context, params model.SessionDisturbReq) error
		// 会话列表
		List(ctx context.Context) (*model.SessionListRes, error)
		// 清除消息未读数
		ClearUnreadMessage(ctx context.Context, params model.SessionClearUnreadNumReq) error
		// 会话免打扰
		OpenContext(ctx context.Context, params model.SessionOpenContextReq) error
		// 获取会话
		FindBySession(ctx context.Context, uid int, receiverId int, talkType int) (*model.SessionItem, error)
	}
)

var (
	localTalkSession ITalkSession
)

func TalkSession() ITalkSession {
	if localTalkSession == nil {
		panic("implement not found for interface ITalkSession, forgot register?")
	}
	return localTalkSession
}

func RegisterTalkSession(i ITalkSession) {
	localTalkSession = i
}
