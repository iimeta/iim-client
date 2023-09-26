package emoticon

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/emoticon/v1"
)

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {

	uploadRes, err := service.Emoticon().Upload(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.UploadRes{
		UploadRes: uploadRes,
	}

	return
}
