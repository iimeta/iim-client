package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) NoteAsterisk(ctx context.Context, req *v1.NoteAsteriskReq) (res *v1.NoteAsteriskRes, err error) {

	err = service.Note().Asterisk(ctx, req.NoteAsteriskReq)

	return
}
