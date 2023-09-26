package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) ApplyAgree(ctx context.Context, req *v1.ApplyAgreeReq) (res *v1.ApplyAgreeRes, err error) {

	err = service.GroupApply().Agree(ctx, req.ApplyAgreeReq)

	return
}
