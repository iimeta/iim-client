package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupApplyCreate(ctx context.Context, req *v1.GroupApplyCreateReq) (res *v1.GroupApplyCreateRes, err error) {

	err = service.GroupApply().Create(ctx, req.GroupApplyCreateReq)

	return
}
