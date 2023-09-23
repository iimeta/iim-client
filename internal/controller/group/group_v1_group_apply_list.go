package group

import (
	"context"
	"github.com/iimeta/iim-client/api/group/v1"
	"github.com/iimeta/iim-client/internal/service"
)

func (c *ControllerV1) GroupApplyList(ctx context.Context, req *v1.GroupApplyListReq) (res *v1.GroupApplyListRes, err error) {

	groupApplyListRes, err := service.GroupApply().List(ctx, req.GroupApplyListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GroupApplyListRes{
		GroupApplyListRes: groupApplyListRes,
	}

	return
}
