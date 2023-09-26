package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) NoteList(ctx context.Context, req *v1.NoteListReq) (res *v1.NoteListRes, err error) {

	noteListRes, err := service.Note().List(ctx, req.NoteListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.NoteListRes{
		NoteListRes: noteListRes,
	}

	return
}
