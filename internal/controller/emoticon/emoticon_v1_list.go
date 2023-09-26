package emoticon

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/emoticon/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	listRes, err := service.Emoticon().CollectList(ctx)

	res = &v1.ListRes{
		ListRes: listRes,
	}

	return
}
