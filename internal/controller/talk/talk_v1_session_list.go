package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) SessionList(ctx context.Context, req *v1.SessionListReq) (res *v1.SessionListRes, err error) {

	talkSessionListRes, err := service.TalkSession().List(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.SessionListRes{
		SessionListRes: talkSessionListRes,
	}

	return
}
