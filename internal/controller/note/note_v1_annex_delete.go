package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) AnnexDelete(ctx context.Context, req *v1.AnnexDeleteReq) (res *v1.AnnexDeleteRes, err error) {

	err = service.NoteAnnex().Delete(ctx, req.AnnexDeleteReq)

	return
}
