package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) NoticeList(ctx context.Context, req *v1.NoticeListReq) (res *v1.NoticeListRes, err error) {

	noticeListRes, err := service.GroupNotice().List(ctx, req.NoticeListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.NoticeListRes{
		NoticeListRes: noticeListRes,
	}

	return
}
