package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupDetail(ctx context.Context, req *v1.GroupDetailReq) (res *v1.GroupDetailRes, err error) {

	groupDetailRes, err := service.Group().Detail(ctx, req.GroupDetailReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GroupDetailRes{
		GroupDetailRes: groupDetailRes,
	}

	return
}
