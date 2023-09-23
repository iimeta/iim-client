package message

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/message/v1"
)

func (c *ControllerV1) VoteMessage(ctx context.Context, req *v1.VoteMessageReq) (res *v1.VoteMessageRes, err error) {

	err = service.Message().Vote(ctx, req.VoteMessageReq)

	return
}
