package talk

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) RecordsForward(ctx context.Context, req *v1.RecordsForwardReq) (res *v1.RecordsForwardRes, err error) {

	records, err := service.TalkRecords().GetForwardRecords(ctx, req.RecordsForwardReq)
	if err != nil {
		return nil, err
	}

	res = &v1.RecordsForwardRes{}

	res.Items = records

	return
}
