package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleUploadImage(ctx context.Context, req *v1.ArticleUploadImageReq) (res *v1.ArticleUploadImageRes, err error) {

	articleUploadImageRes, err := service.Note().Upload(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ArticleUploadImageRes{
		ArticleUploadImageRes: articleUploadImageRes,
	}

	return
}
