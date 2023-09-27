package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) SessionDisturb(ctx context.Context, req *v1.SessionDisturbReq) (res *v1.SessionDisturbRes, err error) {

	err = service.TalkSession().Disturb(ctx, req.SessionDisturbReq)

	return
}
