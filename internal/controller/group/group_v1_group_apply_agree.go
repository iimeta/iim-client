package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupApplyAgree(ctx context.Context, req *v1.GroupApplyAgreeReq) (res *v1.GroupApplyAgreeRes, err error) {

	err = service.GroupApply().Agree(ctx, req.GroupApplyAgreeReq)

	return
}
