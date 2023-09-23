package auth

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/auth/v1"
)

func (c *ControllerV1) AuthForget(ctx context.Context, req *v1.AuthForgetReq) (res *v1.AuthForgetRes, err error) {

	err = service.Auth().Forget(ctx, req.AuthForgetReq)
	if err != nil {
		return
	}

	return
}
