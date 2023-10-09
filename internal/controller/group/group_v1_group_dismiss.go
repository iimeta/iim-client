package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupDismiss(ctx context.Context, req *v1.GroupDismissReq) (res *v1.GroupDismissRes, err error) {

	err = service.Group().Dismiss(ctx, req.GroupDismissReq)

	return
}
