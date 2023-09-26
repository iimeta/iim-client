package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) NoteMove(ctx context.Context, req *v1.NoteMoveReq) (res *v1.NoteMoveRes, err error) {

	err = service.Note().Move(ctx, req.NoteMoveReq)

	return
}
