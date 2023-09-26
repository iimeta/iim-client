package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) ApplyUnread(ctx context.Context, req *v1.ApplyUnreadReq) (res *v1.ApplyUnreadRes, err error) {

	applyUnreadNumRes, err := service.GroupApply().ApplyUnreadNum(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ApplyUnreadRes{
		GroupApplyUnreadNumRes: applyUnreadNumRes,
	}

	return
}
