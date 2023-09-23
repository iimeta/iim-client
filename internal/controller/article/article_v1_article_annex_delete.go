package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleAnnexDelete(ctx context.Context, req *v1.ArticleAnnexDeleteReq) (res *v1.ArticleAnnexDeleteRes, err error) {

	err = service.NoteAnnex().Delete(ctx, req.ArticleAnnexDeleteReq)

	return
}
