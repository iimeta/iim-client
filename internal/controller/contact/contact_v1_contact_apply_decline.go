package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactApplyDecline(ctx context.Context, req *v1.ContactApplyDeclineReq) (res *v1.ContactApplyDeclineRes, err error) {

	err = service.Apply().Decline(ctx, req.ContactApplyDeclineReq)

	return
}
