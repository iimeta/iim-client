package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactSearch(ctx context.Context, req *v1.ContactSearchReq) (res *v1.ContactSearchRes, err error) {

	contactSearchRes, err := service.Contact().Search(ctx, req.ContactSearchReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ContactSearchRes{
		ContactSearchRes: contactSearchRes,
	}

	return
}
