package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupApplyUnread(ctx context.Context, req *v1.GroupApplyUnreadReq) (res *v1.GroupApplyUnreadRes, err error) {

	applyUnreadNumRes, err := service.GroupApply().ApplyUnreadNum(ctx)
	res = &v1.GroupApplyUnreadRes{
		ApplyUnreadNumRes: applyUnreadNumRes,
	}

	return
}
