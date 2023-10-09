package auth

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/auth/v1"
)

func (c *ControllerV1) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {

	refreshRes, err := service.Auth().Refresh(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.RefreshRes{
		RefreshRes: refreshRes,
	}

	return
}
