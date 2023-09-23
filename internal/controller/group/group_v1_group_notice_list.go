package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupNoticeList(ctx context.Context, req *v1.GroupNoticeListReq) (res *v1.GroupNoticeListRes, err error) {

	groupNoticeListRes, err := service.GroupNotice().List(ctx, req.GroupNoticeListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GroupNoticeListRes{
		GroupNoticeListRes: groupNoticeListRes,
	}

	return
}
