package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupRemoveMember(ctx context.Context, req *v1.GroupRemoveMemberReq) (res *v1.GroupRemoveMemberRes, err error) {

	err = service.Group().RemoveMembers(ctx, req.GroupRemoveMemberReq)

	return
}
