package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) AnnexRecoverList(ctx context.Context, req *v1.AnnexRecoverListReq) (res *v1.AnnexRecoverListRes, err error) {

	annexRecoverListRes, err := service.NoteAnnex().RecoverList(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.AnnexRecoverListRes{
		AnnexRecoverListRes: annexRecoverListRes,
	}

	return
}
