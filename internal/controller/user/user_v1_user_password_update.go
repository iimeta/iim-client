package user

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/user/v1"
)

func (c *ControllerV1) UserPasswordUpdate(ctx context.Context, req *v1.UserPasswordUpdateReq) (res *v1.UserPasswordUpdateRes, err error) {

	err = service.User().ChangePassword(ctx, req.UserPasswordUpdateReq)

	return
}
