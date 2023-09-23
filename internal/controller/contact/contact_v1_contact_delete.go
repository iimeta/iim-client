package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactDelete(ctx context.Context, req *v1.ContactDeleteReq) (res *v1.ContactDeleteRes, err error) {

	err = service.Contact().Delete(ctx, req.ContactDeleteReq)

	return
}
