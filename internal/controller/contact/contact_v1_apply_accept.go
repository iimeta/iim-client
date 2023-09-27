package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ApplyAccept(ctx context.Context, req *v1.ApplyAcceptReq) (res *v1.ApplyAcceptRes, err error) {

	_, err = service.ContactApply().Accept(ctx, req.ApplyAcceptReq)

	return
}
