package message

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/message/v1"
)

func (c *ControllerV1) FileMessage(ctx context.Context, req *v1.FileMessageReq) (res *v1.FileMessageRes, err error) {

	err = service.Message().File(ctx, req.FileMessageReq)

	return
}
