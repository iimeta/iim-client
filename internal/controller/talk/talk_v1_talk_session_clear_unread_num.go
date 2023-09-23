package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) TalkSessionClearUnreadNum(ctx context.Context, req *v1.TalkSessionClearUnreadNumReq) (res *v1.TalkSessionClearUnreadNumRes, err error) {

	err = service.Session().ClearUnreadMessage(ctx, req.TalkSessionClearUnreadNumReq)

	return
}
