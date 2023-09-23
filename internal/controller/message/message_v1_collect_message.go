package message

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/message/v1"
)

func (c *ControllerV1) CollectMessage(ctx context.Context, req *v1.CollectMessageReq) (res *v1.CollectMessageRes, err error) {

	err = service.Message().Collect(ctx, req.CollectMessageReq)

	return
}
