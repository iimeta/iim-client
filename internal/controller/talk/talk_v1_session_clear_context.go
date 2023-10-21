package talk

import (
	"context"
	"github.com/iimeta/iim-client/api/talk/v1"
	"github.com/iimeta/iim-client/internal/service"
)

func (c *ControllerV1) SessionClearContext(ctx context.Context, req *v1.SessionClearContextReq) (res *v1.SessionClearContextRes, err error) {

	err = service.TalkSession().ClearContext(ctx, req.SessionClearContextReq)

	return
}
