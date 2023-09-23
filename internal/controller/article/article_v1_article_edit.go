package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleEdit(ctx context.Context, req *v1.ArticleEditReq) (res *v1.ArticleEditRes, err error) {

	articleEditRes, err := service.Note().Edit(ctx, req.ArticleEditReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ArticleEditRes{
		ArticleEditRes: articleEditRes,
	}

	return
}
