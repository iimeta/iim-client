package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ApplyDecline(ctx context.Context, req *v1.ApplyDeclineReq) (res *v1.ApplyDeclineRes, err error) {

	err = service.ContactApply().Decline(ctx, req.ApplyDeclineReq)

	return
}
