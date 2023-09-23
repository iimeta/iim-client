package message

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/message/v1"
)

func (c *ControllerV1) VoteMessageHandle(ctx context.Context, req *v1.VoteMessageHandleReq) (res *v1.VoteMessageHandleRes, err error) {

	voteStatistics, err := service.Message().HandleVote(ctx, req.VoteMessageHandleReq)
	if err != nil {
		return nil, err
	}

	res = &v1.VoteMessageHandleRes{
		VoteStatistics: voteStatistics,
	}

	return
}
