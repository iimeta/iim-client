package auth

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/auth/v1"
)

func (c *ControllerV1) Forget(ctx context.Context, req *v1.ForgetReq) (res *v1.ForgetRes, err error) {

	err = service.Auth().Forget(ctx, req.ForgetReq)
	if err != nil {
		return
	}

	return
}
