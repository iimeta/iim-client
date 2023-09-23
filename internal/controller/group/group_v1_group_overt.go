package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupOvert(ctx context.Context, req *v1.GroupOvertReq) (res *v1.GroupOvertRes, err error) {

	err = service.Group().Overt(ctx, req.GroupOvertReq)

	return
}
