package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) RecordsSearchHistory(ctx context.Context, req *v1.RecordsSearchHistoryReq) (res *v1.RecordsSearchHistoryRes, err error) {

	getTalkRecordsRes, err := service.TalkRecords().SearchHistoryRecords(ctx, req.TalkRecordsReq)
	if err != nil {
		return nil, err
	}

	res = &v1.RecordsSearchHistoryRes{
		TalkRecordsRes: getTalkRecordsRes,
	}

	return
}
