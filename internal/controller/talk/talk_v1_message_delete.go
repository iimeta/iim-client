package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) MessageDelete(ctx context.Context, req *v1.MessageDeleteReq) (res *v1.MessageDeleteRes, err error) {

	err = service.TalkMessage().Delete(ctx, req.MessageDeleteReq)

	return
}
