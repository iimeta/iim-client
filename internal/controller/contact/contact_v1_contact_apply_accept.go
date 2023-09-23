package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactApplyAccept(ctx context.Context, req *v1.ContactApplyAcceptReq) (res *v1.ContactApplyAcceptRes, err error) {

	err = service.Apply().Accept(ctx, req.ContactApplyAcceptReq)

	return
}
