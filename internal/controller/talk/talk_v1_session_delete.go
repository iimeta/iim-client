package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) SessionDelete(ctx context.Context, req *v1.SessionDeleteReq) (res *v1.SessionDeleteRes, err error) {

	err = service.TalkSession().Delete(ctx, req.SessionDeleteReq)

	return
}
