package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupList(ctx context.Context, req *v1.GroupListReq) (res *v1.GroupListRes, err error) {

	groupListRes, err := service.Group().List(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.GroupListRes{
		GroupListRes: groupListRes,
	}

	return
}
