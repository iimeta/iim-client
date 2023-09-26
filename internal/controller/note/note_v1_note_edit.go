package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) NoteEdit(ctx context.Context, req *v1.NoteEditReq) (res *v1.NoteEditRes, err error) {

	noteEditRes, err := service.Note().Edit(ctx, req.NoteEditReq)
	if err != nil {
		return nil, err
	}

	res = &v1.NoteEditRes{
		NoteEditRes: noteEditRes,
	}

	return
}
