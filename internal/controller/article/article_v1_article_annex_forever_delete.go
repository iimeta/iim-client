package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleAnnexForeverDelete(ctx context.Context, req *v1.ArticleAnnexForeverDeleteReq) (res *v1.ArticleAnnexForeverDeleteRes, err error) {

	err = service.NoteAnnex().ForeverDelete(ctx, req.ArticleAnnexForeverDeleteReq)

	return
}
