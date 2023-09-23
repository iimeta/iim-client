package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleTagEdit(ctx context.Context, req *v1.ArticleTagEditReq) (res *v1.ArticleTagEditRes, err error) {

	articleTagEditRes, err := service.NoteTag().Edit(ctx, req.ArticleTagEditReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ArticleTagEditRes{
		ArticleTagEditRes: articleTagEditRes,
	}

	return
}
