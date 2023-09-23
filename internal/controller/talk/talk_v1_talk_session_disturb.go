package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) TalkSessionDisturb(ctx context.Context, req *v1.TalkSessionDisturbReq) (res *v1.TalkSessionDisturbRes, err error) {

	err = service.Session().Disturb(ctx, req.TalkSessionDisturbReq)

	return
}
