package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) TagEdit(ctx context.Context, req *v1.TagEditReq) (res *v1.TagEditRes, err error) {

	tagEditRes, err := service.NoteTag().Edit(ctx, req.TagEditReq)
	if err != nil {
		return nil, err
	}

	res = &v1.TagEditRes{
		TagEditRes: tagEditRes,
	}

	return
}
