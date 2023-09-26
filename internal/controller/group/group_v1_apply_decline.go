package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) ApplyDecline(ctx context.Context, req *v1.ApplyDeclineReq) (res *v1.ApplyDeclineRes, err error) {

	err = service.GroupApply().Decline(ctx, req.GroupApplyDeclineReq)

	return
}
