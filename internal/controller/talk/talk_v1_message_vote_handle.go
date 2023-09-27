package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) MessageVoteHandle(ctx context.Context, req *v1.MessageVoteHandleReq) (res *v1.MessageVoteHandleRes, err error) {

	voteStatistics, err := service.TalkMessage().HandleVote(ctx, req.MessageVoteHandleReq)
	if err != nil {
		return nil, err
	}

	res = &v1.MessageVoteHandleRes{
		VoteStatistics: voteStatistics,
	}

	return
}
