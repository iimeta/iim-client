package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleDelete(ctx context.Context, req *v1.ArticleDeleteReq) (res *v1.ArticleDeleteRes, err error) {

	err = service.Note().Delete(ctx, req.ArticleDeleteReq)

	return
}
