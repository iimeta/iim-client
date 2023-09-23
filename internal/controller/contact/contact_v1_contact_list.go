package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactList(ctx context.Context, req *v1.ContactListReq) (res *v1.ContactListRes, err error) {

	contactListRes, err := service.Contact().List(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ContactListRes{
		ContactListRes: contactListRes,
	}

	return
}
