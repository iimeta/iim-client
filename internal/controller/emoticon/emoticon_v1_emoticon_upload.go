package emoticon

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/emoticon/v1"
)

func (c *ControllerV1) EmoticonUpload(ctx context.Context, req *v1.EmoticonUploadReq) (res *v1.EmoticonUploadRes, err error) {

	emoticonUploadRes, err := service.Emoticon().Upload(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.EmoticonUploadRes{
		EmoticonUploadRes: emoticonUploadRes,
	}

	return
}
