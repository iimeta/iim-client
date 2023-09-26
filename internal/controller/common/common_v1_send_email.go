package common

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/common/v1"
)

func (c *ControllerV1) SendEmail(ctx context.Context, req *v1.SendEmailReq) (res *v1.SendEmailRes, err error) {

	sendEmailRes, err := service.Common().EmailCode(ctx, req.SendEmailReq)
	if err != nil {
		return
	}

	res = &v1.SendEmailRes{
		SendEmailRes: sendEmailRes,
	}

	return
}
