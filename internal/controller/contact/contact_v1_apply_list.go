package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ApplyList(ctx context.Context, req *v1.ApplyListReq) (res *v1.ApplyListRes, err error) {

	applyItems, err := service.ContactApply().List(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ApplyListRes{}
	res.Items = applyItems

	return
}
