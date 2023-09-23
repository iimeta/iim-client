package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleClassDelete(ctx context.Context, req *v1.ArticleClassDeleteReq) (res *v1.ArticleClassDeleteRes, err error) {

	err = service.NoteClass().Delete(ctx, req.ArticleClassDeleteReq)

	return
}
