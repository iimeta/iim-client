package common

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/common/v1"
)

func (c *ControllerV1) SendSms(ctx context.Context, req *v1.SendSmsReq) (res *v1.SendSmsRes, err error) {

	sendSmsRes, err := service.Common().SmsCode(ctx, req.SendSmsReq)
	if err != nil {
		return nil, err
	}

	res = &v1.SendSmsRes{
		SendSmsRes: sendSmsRes,
	}

	return
}
