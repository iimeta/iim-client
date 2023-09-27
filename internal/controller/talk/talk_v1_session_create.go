package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) SessionCreate(ctx context.Context, req *v1.SessionCreateReq) (res *v1.SessionCreateRes, err error) {

	talkSessionCreateRes, err := service.TalkSession().Create(ctx, req.SessionCreateReq)
	if err != nil {
		return nil, err
	}

	res = &v1.SessionCreateRes{
		SessionCreateRes: talkSessionCreateRes,
	}

	return
}
