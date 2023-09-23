package common

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/common/v1"
)

func (c *ControllerV1) CommonSendEmail(ctx context.Context, req *v1.CommonSendEmailReq) (res *v1.CommonSendEmailRes, err error) {

	commonSendEmailRes, err := service.Common().EmailCode(ctx, req.CommonSendEmailReq)
	if err != nil {
		return
	}

	res = &v1.CommonSendEmailRes{
		CommonSendEmailRes: commonSendEmailRes,
	}

	return
}
