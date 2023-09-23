package message

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/message/v1"
)

func (c *ControllerV1) DeleteMessage(ctx context.Context, req *v1.DeleteMessageReq) (res *v1.DeleteMessageRes, err error) {

	err = service.Message().Delete(ctx, req.DeleteMessageReq)

	return
}
