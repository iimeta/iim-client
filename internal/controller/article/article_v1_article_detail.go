package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleDetail(ctx context.Context, req *v1.ArticleDetailReq) (res *v1.ArticleDetailRes, err error) {

	articleDetailRes, err := service.Note().Detail(ctx, req.ArticleDetailReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ArticleDetailRes{
		ArticleDetailRes: articleDetailRes,
	}

	return
}
