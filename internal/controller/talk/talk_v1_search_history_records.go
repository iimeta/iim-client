package talk

import (
	"context"
	"github.com/iimeta/iim-client/api/talk/v1"
	"github.com/iimeta/iim-client/internal/service"
)

func (c *ControllerV1) SearchHistoryRecords(ctx context.Context, req *v1.SearchHistoryRecordsReq) (res *v1.SearchHistoryRecordsRes, err error) {

	getTalkRecordsRes, err := service.Records().SearchHistoryRecords(ctx, req.GetTalkRecordsReq)
	if err != nil {
		return nil, err
	}

	res = &v1.SearchHistoryRecordsRes{
		GetTalkRecordsRes: getTalkRecordsRes,
	}

	return
}
