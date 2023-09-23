package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleClassSort(ctx context.Context, req *v1.ArticleClassSortReq) (res *v1.ArticleClassSortRes, err error) {

	err = service.NoteClass().Sort(ctx, req.ArticleClassSortReq)

	return
}
