package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupApplyDelete(ctx context.Context, req *v1.GroupApplyDeleteReq) (res *v1.GroupApplyDeleteRes, err error) {

	err = service.GroupApply().Delete(ctx, req.ApplyId, service.Session().GetUid(ctx))

	return
}
