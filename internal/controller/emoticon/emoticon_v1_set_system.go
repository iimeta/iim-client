package emoticon

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/emoticon/v1"
)

func (c *ControllerV1) SetSystem(ctx context.Context, req *v1.SetSystemReq) (res *v1.SetSystemRes, err error) {

	setSystemRes, err := service.Emoticon().SetSystemEmoticon(ctx, req.SetSystemReq)
	if err != nil {
		return nil, err
	}

	res = &v1.SetSystemRes{
		SetSystemRes: setSystemRes,
	}

	return
}
