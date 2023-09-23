// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"github.com/iimeta/iim-client/api/auth/v1"
)

type IAuthV1 interface {
	AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error)
	AuthLogout(ctx context.Context, req *v1.AuthLogoutReq) (res *v1.AuthLogoutRes, err error)
	AuthRegister(ctx context.Context, req *v1.AuthRegisterReq) (res *v1.AuthRegisterRes, err error)
	AuthRefresh(ctx context.Context, req *v1.AuthRefreshReq) (res *v1.AuthRefreshRes, err error)
	AuthForget(ctx context.Context, req *v1.AuthForgetReq) (res *v1.AuthForgetRes, err error)
}
