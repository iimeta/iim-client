package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) TalkSessionTop(ctx context.Context, req *v1.TalkSessionTopReq) (res *v1.TalkSessionTopRes, err error) {

	err = service.Session().Top(ctx, req.TalkSessionTopReq)

	return
}
