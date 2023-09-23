package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) TalkSessionList(ctx context.Context, req *v1.TalkSessionListReq) (res *v1.TalkSessionListRes, err error) {

	talkSessionListRes, err := service.Session().List(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.TalkSessionListRes{
		TalkSessionListRes: talkSessionListRes,
	}

	return
}
