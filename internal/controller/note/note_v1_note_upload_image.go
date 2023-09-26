package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) NoteUploadImage(ctx context.Context, req *v1.NoteUploadImageReq) (res *v1.NoteUploadImageRes, err error) {

	noteUploadImageRes, err := service.Note().Upload(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.NoteUploadImageRes{
		NoteUploadImageRes: noteUploadImageRes,
	}

	return
}
