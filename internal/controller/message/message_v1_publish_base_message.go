package message

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/message/v1"
)

func (c *ControllerV1) PublishBaseMessage(ctx context.Context, req *v1.PublishBaseMessageReq) (res *v1.PublishBaseMessageRes, err error) {

	err = service.Publish().Publish(ctx, req.PublishBaseMessageReq)

	return
}
