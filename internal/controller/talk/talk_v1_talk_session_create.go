package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) TalkSessionCreate(ctx context.Context, req *v1.TalkSessionCreateReq) (res *v1.TalkSessionCreateRes, err error) {

	talkSessionCreateRes, err := service.Session().Create(ctx, req.TalkSessionCreateReq)
	if err != nil {
		return nil, err
	}

	res = &v1.TalkSessionCreateRes{
		TalkSessionCreateRes: talkSessionCreateRes,
	}

	return
}
