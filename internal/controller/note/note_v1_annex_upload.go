package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) AnnexUpload(ctx context.Context, req *v1.AnnexUploadReq) (res *v1.AnnexUploadRes, err error) {

	annexUploadRes, err := service.NoteAnnex().Upload(ctx, req.AnnexUploadReq)
	if err != nil {
		return nil, err
	}

	res = &v1.AnnexUploadRes{
		AnnexUploadRes: annexUploadRes,
	}

	return
}
