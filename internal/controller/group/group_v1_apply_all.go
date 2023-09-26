package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) ApplyAll(ctx context.Context, req *v1.ApplyAllReq) (res *v1.ApplyAllRes, err error) {

	applyAllRes, err := service.GroupApply().All(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ApplyAllRes{
		ApplyAllRes: applyAllRes,
	}

	return
}
