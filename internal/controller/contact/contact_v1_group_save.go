package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) GroupSave(ctx context.Context, req *v1.GroupSaveReq) (res *v1.GroupSaveRes, err error) {

	err = service.ContactGroup().Save(ctx, req.GroupSaveReq)

	return
}
