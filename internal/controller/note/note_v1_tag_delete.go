package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) TagDelete(ctx context.Context, req *v1.TagDeleteReq) (res *v1.TagDeleteRes, err error) {

	err = service.NoteTag().Delete(ctx, req.TagDeleteReq)

	return
}
