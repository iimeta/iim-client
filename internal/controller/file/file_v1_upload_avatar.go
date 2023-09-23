package file

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/file/v1"
)

func (c *ControllerV1) UploadAvatar(ctx context.Context, req *v1.UploadAvatarReq) (res *v1.UploadAvatarRes, err error) {

	uploadAvatarRes, err := service.File().Avatar(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.UploadAvatarRes{
		UploadAvatarRes: uploadAvatarRes,
	}

	return
}
