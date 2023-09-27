package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) SessionTop(ctx context.Context, req *v1.SessionTopReq) (res *v1.SessionTopRes, err error) {

	err = service.TalkSession().Top(ctx, req.SessionTopReq)

	return
}
