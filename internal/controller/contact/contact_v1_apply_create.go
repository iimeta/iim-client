package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ApplyCreate(ctx context.Context, req *v1.ApplyCreateReq) (res *v1.ApplyCreateRes, err error) {

	err = service.ContactApply().ApplyCreate(ctx, req.ApplyCreateReq)

	return
}
