package file

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/file/v1"
)

func (c *ControllerV1) UploadImage(ctx context.Context, req *v1.UploadImageReq) (res *v1.UploadImageRes, err error) {

	uploadImageRes, err := service.File().Image(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.UploadImageRes{
		UploadImageRes: uploadImageRes,
	}

	return
}
