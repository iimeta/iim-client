package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) NoteForeverDelete(ctx context.Context, req *v1.NoteForeverDeleteReq) (res *v1.NoteForeverDeleteRes, err error) {

	err = service.Note().ForeverDelete(ctx, req.NoteForeverDeleteReq)

	return
}
