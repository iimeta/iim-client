package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleAnnexDownload(ctx context.Context, req *v1.ArticleAnnexDownloadReq) (res *v1.ArticleAnnexDownloadRes, err error) {

	err = service.NoteAnnex().Download(ctx, req.ArticleAnnexDownloadReq)

	return
}
