package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleForeverDelete(ctx context.Context, req *v1.ArticleForeverDeleteReq) (res *v1.ArticleForeverDeleteRes, err error) {

	err = service.Note().ForeverDelete(ctx, req.ArticleForeverDeleteReq)

	return
}
