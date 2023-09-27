package file

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/file/v1"
)

func (c *ControllerV1) UploadMultipart(ctx context.Context, req *v1.UploadMultipartReq) (res *v1.UploadMultipartRes, err error) {

	uploadMultipartRes, err := service.File().MultipartSplitUpload(ctx, req.UploadMultipartReq)
	if err != nil {
		return nil, err
	}

	res = &v1.UploadMultipartRes{
		UploadMultipartRes: uploadMultipartRes,
	}

	return
}
