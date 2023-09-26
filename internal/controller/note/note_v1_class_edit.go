package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) ClassEdit(ctx context.Context, req *v1.ClassEditReq) (res *v1.ClassEditRes, err error) {

	classEditRes, err := service.NoteClass().Edit(ctx, req.ClassEditReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ClassEditRes{
		ClassEditRes: classEditRes,
	}

	return
}
