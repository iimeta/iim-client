package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) SessionClearUnreadNum(ctx context.Context, req *v1.SessionClearUnreadNumReq) (res *v1.SessionClearUnreadNumRes, err error) {

	err = service.TalkSession().ClearUnreadMessage(ctx, req.SessionClearUnreadNumReq)

	return
}
