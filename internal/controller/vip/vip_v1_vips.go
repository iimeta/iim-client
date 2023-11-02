package vip

import (
	"context"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/vip/v1"
)

func (c *ControllerV1) Vips(ctx context.Context, req *v1.VipsReq) (res *v1.VipsRes, err error) {

	vips, err := service.Vip().Vips(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.VipsRes{
		VipsRes: &model.VipsRes{
			Items: vips,
		},
	}

	return
}
