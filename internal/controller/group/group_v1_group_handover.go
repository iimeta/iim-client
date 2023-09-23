package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupHandover(ctx context.Context, req *v1.GroupHandoverReq) (res *v1.GroupHandoverRes, err error) {

	err = service.Group().Handover(ctx, req.GroupHandoverReq)

	return
}
