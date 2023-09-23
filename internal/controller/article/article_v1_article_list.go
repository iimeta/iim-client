package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleList(ctx context.Context, req *v1.ArticleListReq) (res *v1.ArticleListRes, err error) {

	articleListRes, err := service.Note().List(ctx, req.ArticleListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ArticleListRes{
		ArticleListRes: articleListRes,
	}

	return
}
