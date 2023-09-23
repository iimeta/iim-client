package talk

import (
	"context"
	"github.com/iimeta/iim-client/api/talk/v1"
	"github.com/iimeta/iim-client/internal/service"
)

func (c *ControllerV1) GetForwardRecords(ctx context.Context, req *v1.GetForwardRecordsReq) (res *v1.GetForwardRecordsRes, err error) {

	getTalkRecordsRes, err := service.Records().GetForwardRecords(ctx, req.GetForwardTalkRecordReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetForwardRecordsRes{
		GetTalkRecordsRes: getTalkRecordsRes,
	}

	return
}
