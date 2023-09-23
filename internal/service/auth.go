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
	IAuth interface {
		// Login 登录接口
		Login(ctx context.Context, params model.AuthLoginReq) (*model.AuthLoginRes, error)
		// Register 注册接口
		Register(ctx context.Context, params model.AuthRegisterReq) error
		// Logout 退出登录接口
		Logout(ctx context.Context) error
		// Refresh Token 刷新接口
		Refresh(ctx context.Context) (*model.AuthRefreshRes, error)
		// Forget 账号找回接口
		Forget(ctx context.Context, params model.AuthForgetReq) error
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
