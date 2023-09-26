package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) AnnexRecover(ctx context.Context, req *v1.AnnexRecoverReq) (res *v1.AnnexRecoverRes, err error) {

	err = service.NoteAnnex().Recover(ctx, req.AnnexRecoverReq)

	return
}
