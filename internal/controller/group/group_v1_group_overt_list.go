package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupOvertList(ctx context.Context, req *v1.GroupOvertListReq) (res *v1.GroupOvertListRes, err error) {

	groupOvertListRes, err := service.Group().OvertList(ctx, req.GroupOvertListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GroupOvertListRes{
		GroupOvertListRes: groupOvertListRes,
	}

	return
}
