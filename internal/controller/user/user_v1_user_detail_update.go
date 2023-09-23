package user

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/user/v1"
)

func (c *ControllerV1) UserDetailUpdate(ctx context.Context, req *v1.UserDetailUpdateReq) (res *v1.UserDetailUpdateRes, err error) {

	err = service.User().ChangeDetail(ctx, req.UserDetailUpdateReq)

	return
}
