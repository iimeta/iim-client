package file

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/file/v1"
)

func (c *ControllerV1) UploadInitiateMultipart(ctx context.Context, req *v1.UploadInitiateMultipartReq) (res *v1.UploadInitiateMultipartRes, err error) {

	uploadInitiateMultipartRes, err := service.File().InitiateMultipart(ctx, req.UploadInitiateMultipartReq)
	if err != nil {
		return nil, err
	}

	res = &v1.UploadInitiateMultipartRes{
		UploadInitiateMultipartRes: uploadInitiateMultipartRes,
	}

	return
}
