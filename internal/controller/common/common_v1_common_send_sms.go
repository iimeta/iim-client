package common

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/common/v1"
)

func (c *ControllerV1) CommonSendSms(ctx context.Context, req *v1.CommonSendSmsReq) (res *v1.CommonSendSmsRes, err error) {

	commonSendSmsRes, err := service.Common().SmsCode(ctx, req.CommonSendSmsReq)
	if err != nil {
		return
	}

	res = &v1.CommonSendSmsRes{
		CommonSendSmsRes: commonSendSmsRes,
	}

	return
}
