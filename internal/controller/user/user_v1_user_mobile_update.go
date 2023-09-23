package user

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/user/v1"
)

func (c *ControllerV1) UserMobileUpdate(ctx context.Context, req *v1.UserMobileUpdateReq) (res *v1.UserMobileUpdateRes, err error) {

	err = service.User().ChangeMobile(ctx, req.UserMobileUpdateReq)

	return
}
