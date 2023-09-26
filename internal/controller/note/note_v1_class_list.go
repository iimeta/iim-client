package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) ClassList(ctx context.Context, req *v1.ClassListReq) (res *v1.ClassListRes, err error) {

	classListRes, err := service.NoteClass().List(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ClassListRes{
		ClassListRes: classListRes,
	}

	return
}
