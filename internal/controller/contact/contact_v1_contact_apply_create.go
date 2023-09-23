package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactApplyCreate(ctx context.Context, req *v1.ContactApplyCreateReq) (res *v1.ContactApplyCreateRes, err error) {

	err = service.Apply().Create(ctx, req.ContactApplyCreateReq)

	return
}
