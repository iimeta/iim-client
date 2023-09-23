package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupRemarkUpdate(ctx context.Context, req *v1.GroupRemarkUpdateReq) (res *v1.GroupRemarkUpdateRes, err error) {

	err = service.Group().UpdateMemberRemark(ctx, req.GroupRemarkUpdateReq)

	return
}
