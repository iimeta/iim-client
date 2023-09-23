package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactGroupList(ctx context.Context, req *v1.ContactGroupListReq) (res *v1.ContactGroupListRes, err error) {

	contactGroupListRes, err := service.ContactGroup().List(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ContactGroupListRes{
		ContactGroupListRes: contactGroupListRes,
	}

	return
}
