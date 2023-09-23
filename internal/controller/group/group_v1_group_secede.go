package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupSecede(ctx context.Context, req *v1.GroupSecedeReq) (res *v1.GroupSecedeRes, err error) {

	err = service.Group().Secede(ctx, req.GroupId, service.Session().GetUid(ctx))

	return
}
