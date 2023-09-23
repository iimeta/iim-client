package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleTagList(ctx context.Context, req *v1.ArticleTagListReq) (res *v1.ArticleTagListRes, err error) {

	articleTagListRes, err := service.NoteTag().List(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ArticleTagListRes{
		ArticleTagListRes: articleTagListRes,
	}

	return
}
