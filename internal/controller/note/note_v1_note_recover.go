package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) NoteRecover(ctx context.Context, req *v1.NoteRecoverReq) (res *v1.NoteRecoverRes, err error) {

	err = service.Note().Recover(ctx, req.NoteRecoverReq)

	return
}
