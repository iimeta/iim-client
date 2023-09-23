package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupAssignAdmin(ctx context.Context, req *v1.GroupAssignAdminReq) (res *v1.GroupAssignAdminRes, err error) {

	err = service.Group().AssignAdmin(ctx, req.GroupAssignAdminReq)

	return
}
