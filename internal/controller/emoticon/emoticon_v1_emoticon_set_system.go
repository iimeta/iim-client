package emoticon

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/emoticon/v1"
)

func (c *ControllerV1) EmoticonSetSystem(ctx context.Context, req *v1.EmoticonSetSystemReq) (res *v1.EmoticonSetSystemRes, err error) {

	emoticonSetSystemRes, err := service.Emoticon().SetSystemEmoticon(ctx, req.EmoticonSetSystemReq)
	if err != nil {
		return nil, err
	}

	res = &v1.EmoticonSetSystemRes{
		EmoticonSetSystemRes: emoticonSetSystemRes,
	}

	return
}
