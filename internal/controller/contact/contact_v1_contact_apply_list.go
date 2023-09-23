package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactApplyList(ctx context.Context, req *v1.ContactApplyListReq) (res *v1.ContactApplyListRes, err error) {

	contactApplyListRes, err := service.Apply().List(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ContactApplyListRes{
		ContactApplyListRes: contactApplyListRes,
	}

	return
}
