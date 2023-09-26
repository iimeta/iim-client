package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) ApplyList(ctx context.Context, req *v1.ApplyListReq) (res *v1.ApplyListRes, err error) {

	groupApplyListRes, err := service.GroupApply().List(ctx, req.ApplyListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ApplyListRes{
		GroupApplyListRes: groupApplyListRes,
	}

	return
}
