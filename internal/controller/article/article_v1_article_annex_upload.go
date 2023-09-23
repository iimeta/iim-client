package article

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/article/v1"
)

func (c *ControllerV1) ArticleAnnexUpload(ctx context.Context, req *v1.ArticleAnnexUploadReq) (res *v1.ArticleAnnexUploadRes, err error) {

	articleAnnexUploadRes, err := service.NoteAnnex().Upload(ctx, req.ArticleAnnexUploadReq)
	if err != nil {
		return nil, err
	}

	res = &v1.ArticleAnnexUploadRes{
		ArticleAnnexUploadRes: articleAnnexUploadRes,
	}

	return
}
