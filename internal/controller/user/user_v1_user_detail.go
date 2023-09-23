package user

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/user/v1"
)

func (c *ControllerV1) UserDetail(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error) {

	userDetailRes, err := service.User().Detail(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.UserDetailRes{
		UserDetailRes: userDetailRes,
	}

	return
}
