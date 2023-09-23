package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleAnnexRecover(ctx context.Context, req *v1.ArticleAnnexRecoverReq) (res *v1.ArticleAnnexRecoverRes, err error) {

	err = service.NoteAnnex().Recover(ctx, req.ArticleAnnexRecoverReq)

	return
}
