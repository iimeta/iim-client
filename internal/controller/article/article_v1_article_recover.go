package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleRecover(ctx context.Context, req *v1.ArticleRecoverReq) (res *v1.ArticleRecoverRes, err error) {

	err = service.Note().Recover(ctx, req.ArticleRecoverReq)

	return
}
