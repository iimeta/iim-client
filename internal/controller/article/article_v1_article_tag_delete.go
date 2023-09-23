package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleTagDelete(ctx context.Context, req *v1.ArticleTagDeleteReq) (res *v1.ArticleTagDeleteRes, err error) {

	err = service.NoteTag().Delete(ctx, req.ArticleTagDeleteReq)

	return
}
