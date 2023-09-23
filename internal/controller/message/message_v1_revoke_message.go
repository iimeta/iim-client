package message

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/message/v1"
)

func (c *ControllerV1) RevokeMessage(ctx context.Context, req *v1.RevokeMessageReq) (res *v1.RevokeMessageRes, err error) {

	err = service.Message().Revoke(ctx, req.RevokeMessageReq)

	return
}
