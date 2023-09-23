package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleClassList(ctx context.Context, req *v1.ArticleClassListReq) (res *v1.ArticleClassListRes, err error) {

	articleClassListRes, err := service.NoteClass().List(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ArticleClassListRes{
		ArticleClassListRes: articleClassListRes,
	}

	return
}
