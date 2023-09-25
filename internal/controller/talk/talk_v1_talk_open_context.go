package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) TalkOpenContext(ctx context.Context, req *v1.TalkOpenContextReq) (res *v1.TalkOpenContextRes, err error) {

	err = service.Session().OpenContext(ctx, req.TalkOpenContextReq)

	return
}
