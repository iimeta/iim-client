package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactDetail(ctx context.Context, req *v1.ContactDetailReq) (res *v1.ContactDetailRes, err error) {

	detail, err := service.Contact().Detail(ctx, req.ContactDetailReq)

	if err != nil {
		return nil, err
	}

	res = &v1.ContactDetailRes{
		ContactDetailRes: detail,
	}

	return
}
