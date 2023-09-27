package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) MessagePublish(ctx context.Context, req *v1.MessagePublishReq) (res *v1.MessagePublishRes, err error) {

	err = service.TalkMessage().Publish(ctx, req.MessagePublishReq)

	return
}
