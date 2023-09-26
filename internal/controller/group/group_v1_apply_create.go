package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) ApplyCreate(ctx context.Context, req *v1.ApplyCreateReq) (res *v1.ApplyCreateRes, err error) {

	err = service.GroupApply().Create(ctx, req.GroupApplyCreateReq)

	return
}
