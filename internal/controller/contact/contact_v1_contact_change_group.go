package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactChangeGroup(ctx context.Context, req *v1.ContactChangeGroupReq) (res *v1.ContactChangeGroupRes, err error) {

	err = service.Contact().MoveGroup(ctx, req.ContactChangeGroupReq)

	return
}
