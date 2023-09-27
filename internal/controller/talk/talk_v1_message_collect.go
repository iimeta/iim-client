package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) MessageCollect(ctx context.Context, req *v1.MessageCollectReq) (res *v1.MessageCollectRes, err error) {

	err = service.TalkMessage().Collect(ctx, req.MessageCollectReq)

	return
}
