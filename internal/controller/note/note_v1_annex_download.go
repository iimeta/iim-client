package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) AnnexDownload(ctx context.Context, req *v1.AnnexDownloadReq) (res *v1.AnnexDownloadRes, err error) {

	err = service.NoteAnnex().Download(ctx, req.AnnexDownloadReq)

	return
}
