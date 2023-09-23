package talk

import (
	"context"
	"github.com/iimeta/iim-client/api/talk/v1"
	"github.com/iimeta/iim-client/internal/service"
)

func (c *ControllerV1) GetRecords(ctx context.Context, req *v1.GetRecordsReq) (res *v1.GetRecordsRes, err error) {

	getTalkRecordsRes, err := service.Records().GetRecords(ctx, req.GetTalkRecordsReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetRecordsRes{
		GetTalkRecordsRes: getTalkRecordsRes,
	}

	return
}
