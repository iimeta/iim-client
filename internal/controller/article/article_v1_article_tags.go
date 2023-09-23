package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleTags(ctx context.Context, req *v1.ArticleTagsReq) (res *v1.ArticleTagsRes, err error) {

	err = service.Note().Tag(ctx, req.ArticleTagsReq)

	return
}
