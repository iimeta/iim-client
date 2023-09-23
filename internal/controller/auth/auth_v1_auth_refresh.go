package auth

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/auth/v1"
)

func (c *ControllerV1) AuthRefresh(ctx context.Context, req *v1.AuthRefreshReq) (res *v1.AuthRefreshRes, err error) {

	authRefreshRes, err := service.Auth().Refresh(ctx)
	if err != nil {
		return
	}

	res = &v1.AuthRefreshRes{
		AuthRefreshRes: authRefreshRes,
	}

	return
}
