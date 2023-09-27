package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) SessionOpenContext(ctx context.Context, req *v1.SessionOpenContextReq) (res *v1.SessionOpenContextRes, err error) {

	err = service.TalkSession().OpenContext(ctx, req.SessionOpenContextReq)

	return
}
