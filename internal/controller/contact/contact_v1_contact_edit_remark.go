package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactEditRemark(ctx context.Context, req *v1.ContactEditRemarkReq) (res *v1.ContactEditRemarkRes, err error) {

	err = service.Contact().Remark(ctx, req.ContactEditRemarkReq)

	return
}
