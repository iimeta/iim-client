package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleMove(ctx context.Context, req *v1.ArticleMoveReq) (res *v1.ArticleMoveRes, err error) {

	err = service.Note().Move(ctx, req.ArticleMoveReq)

	return
}
