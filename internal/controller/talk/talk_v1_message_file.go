package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) MessageFile(ctx context.Context, req *v1.MessageFileReq) (res *v1.MessageFileRes, err error) {

	err = service.TalkMessage().File(ctx, req.MessageFileReq)

	return
}
