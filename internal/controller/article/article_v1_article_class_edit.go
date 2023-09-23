package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleClassEdit(ctx context.Context, req *v1.ArticleClassEditReq) (res *v1.ArticleClassEditRes, err error) {

	articleClassEditRes, err := service.NoteClass().Edit(ctx, req.ArticleClassEditReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ArticleClassEditRes{
		ArticleClassEditRes: articleClassEditRes,
	}

	return
}
