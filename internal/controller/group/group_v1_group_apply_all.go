package group

import (
	"context"
	"github.com/iimeta/iim-client/api/group/v1"
	"github.com/iimeta/iim-client/internal/service"
)

func (c *ControllerV1) GroupApplyAll(ctx context.Context, req *v1.GroupApplyAllReq) (res *v1.GroupApplyAllRes, err error) {

	groupApplyAllRes, err := service.GroupApply().All(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.GroupApplyAllRes{
		GroupApplyAllRes: groupApplyAllRes,
	}

	return
}
