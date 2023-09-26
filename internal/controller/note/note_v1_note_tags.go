package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) NoteTags(ctx context.Context, req *v1.NoteTagsReq) (res *v1.NoteTagsRes, err error) {

	err = service.Note().Tag(ctx, req.NoteTagsReq)

	return
}
