package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) AnnexForeverDelete(ctx context.Context, req *v1.AnnexForeverDeleteReq) (res *v1.AnnexForeverDeleteRes, err error) {

	err = service.NoteAnnex().ForeverDelete(ctx, req.AnnexForeverDeleteReq)

	return
}
