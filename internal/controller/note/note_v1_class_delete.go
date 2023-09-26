package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) ClassDelete(ctx context.Context, req *v1.ClassDeleteReq) (res *v1.ClassDeleteRes, err error) {

	err = service.NoteClass().Delete(ctx, req.ClassDeleteReq)

	return
}
