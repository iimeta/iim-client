package auth

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/auth/v1"
)

func (c *ControllerV1) AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {

	authLoginRes, err := service.Auth().Login(ctx, req.AuthLoginReq)
	if err != nil {
		return
	}

	res = &v1.AuthLoginRes{
		AuthLoginRes: authLoginRes,
	}

	return
}
