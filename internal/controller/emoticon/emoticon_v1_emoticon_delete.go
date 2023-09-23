package emoticon

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/emoticon/v1"
)

func (c *ControllerV1) EmoticonDelete(ctx context.Context, req *v1.EmoticonDeleteReq) (res *v1.EmoticonDeleteRes, err error) {

	err = service.Emoticon().DeleteCollect(ctx, req.EmoticonDeleteReq)

	return
}
