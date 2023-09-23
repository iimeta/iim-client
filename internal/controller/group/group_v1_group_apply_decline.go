package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupApplyDecline(ctx context.Context, req *v1.GroupApplyDeclineReq) (res *v1.GroupApplyDeclineRes, err error) {

	err = service.GroupApply().Decline(ctx, req.GroupApplyDeclineReq)

	return
}
