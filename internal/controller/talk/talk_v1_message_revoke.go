package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) MessageRevoke(ctx context.Context, req *v1.MessageRevokeReq) (res *v1.MessageRevokeRes, err error) {

	err = service.TalkMessage().Revoke(ctx, req.MessageRevokeReq)

	return
}
