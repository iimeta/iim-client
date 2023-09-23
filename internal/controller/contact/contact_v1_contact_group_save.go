package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactGroupSave(ctx context.Context, req *v1.ContactGroupSaveReq) (res *v1.ContactGroupSaveRes, err error) {

	err = service.ContactGroup().Save(ctx, req.ContactGroupSaveReq)

	return
}
