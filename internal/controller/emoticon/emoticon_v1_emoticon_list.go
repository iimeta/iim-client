package emoticon

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/emoticon/v1"
)

func (c *ControllerV1) EmoticonList(ctx context.Context, req *v1.EmoticonListReq) (res *v1.EmoticonListRes, err error) {

	emoticonListRes, err := service.Emoticon().CollectList(ctx)

	res = &v1.EmoticonListRes{
		EmoticonListRes: emoticonListRes,
	}

	return
}
