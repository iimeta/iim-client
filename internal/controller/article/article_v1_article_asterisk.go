package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleAsterisk(ctx context.Context, req *v1.ArticleAsteriskReq) (res *v1.ArticleAsteriskRes, err error) {

	err = service.Note().Asterisk(ctx, req.ArticleAsteriskReq)

	return
}
