package auth

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/auth/v1"
)

func (c *ControllerV1) AuthLogout(ctx context.Context, req *v1.AuthLogoutReq) (res *v1.AuthLogoutRes, err error) {

	err = service.Auth().Logout(ctx)
	if err != nil {
		return
	}

	return
}
