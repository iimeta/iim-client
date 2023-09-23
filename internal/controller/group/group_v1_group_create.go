package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupCreate(ctx context.Context, req *v1.GroupCreateReq) (res *v1.GroupCreateRes, err error) {

	groupCreateRes, err := service.Group().Create(ctx, req.GroupCreateReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GroupCreateRes{
		GroupCreateRes: groupCreateRes,
	}

	return
}
