package vip

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/vip/v1"
)

func (c *ControllerV1) VipInfo(ctx context.Context, req *v1.VipInfoReq) (res *v1.VipInfoRes, err error) {

	vipInfo, err := service.Vip().VipInfo(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.VipInfoRes{
		VipInfo: vipInfo,
	}

	return
}
