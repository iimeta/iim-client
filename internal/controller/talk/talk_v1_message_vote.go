package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) MessageVote(ctx context.Context, req *v1.MessageVoteReq) (res *v1.MessageVoteRes, err error) {

	err = service.TalkMessage().Vote(ctx, req.MessageVoteReq)

	return
}
