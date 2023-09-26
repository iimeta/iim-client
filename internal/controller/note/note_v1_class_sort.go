package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) ClassSort(ctx context.Context, req *v1.ClassSortReq) (res *v1.ClassSortRes, err error) {

	err = service.NoteClass().Sort(ctx, req.ClassSortReq)

	return
}
