package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) ApplyDelete(ctx context.Context, req *v1.ApplyDeleteReq) (res *v1.ApplyDeleteRes, err error) {

	err = service.GroupApply().Delete(ctx, req.ApplyId, service.Session().GetUid(ctx))

	return
}
