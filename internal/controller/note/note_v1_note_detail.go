package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) NoteDetail(ctx context.Context, req *v1.NoteDetailReq) (res *v1.NoteDetailRes, err error) {

	noteDetailRes, err := service.Note().Detail(ctx, req.NoteDetailReq)
	if err != nil {
		return nil, err
	}

	res = &v1.NoteDetailRes{
		NoteDetailRes: noteDetailRes,
	}

	return
}
