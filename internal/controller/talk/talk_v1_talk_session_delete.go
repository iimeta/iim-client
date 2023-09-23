package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) TalkSessionDelete(ctx context.Context, req *v1.TalkSessionDeleteReq) (res *v1.TalkSessionDeleteRes, err error) {

	err = service.Session().Delete(ctx, req.TalkSessionDeleteReq)

	return
}
