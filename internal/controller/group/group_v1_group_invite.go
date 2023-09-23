package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupInvite(ctx context.Context, req *v1.GroupInviteReq) (res *v1.GroupInviteRes, err error) {

	err = service.Group().Invite(ctx, req.GroupInviteReq)

	return
}
