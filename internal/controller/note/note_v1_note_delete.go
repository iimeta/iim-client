package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) NoteDelete(ctx context.Context, req *v1.NoteDeleteReq) (res *v1.NoteDeleteRes, err error) {

	err = service.Note().Delete(ctx, req.NoteDeleteReq)

	return
}
