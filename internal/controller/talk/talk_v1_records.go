package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) Records(ctx context.Context, req *v1.RecordsReq) (res *v1.RecordsRes, err error) {

	talkRecordsRes, err := service.TalkRecords().GetRecords(ctx, req.TalkRecordsReq)
	if err != nil {
		return nil, err
	}

	res = &v1.RecordsRes{
		TalkRecordsRes: talkRecordsRes,
	}

	return
}
