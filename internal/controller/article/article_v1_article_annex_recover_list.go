package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleAnnexRecoverList(ctx context.Context, req *v1.ArticleAnnexRecoverListReq) (res *v1.ArticleAnnexRecoverListRes, err error) {

	articleAnnexRecoverListRes, err := service.NoteAnnex().RecoverList(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ArticleAnnexRecoverListRes{
		ArticleAnnexRecoverListRes: articleAnnexRecoverListRes,
	}

	return
}
