// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package common

import (
	"context"

	"github.com/iimeta/iim-client/api/common/v1"
)

type ICommonV1 interface {
	CommonSendSms(ctx context.Context, req *v1.CommonSendSmsReq) (res *v1.CommonSendSmsRes, err error)
	CommonSendEmail(ctx context.Context, req *v1.CommonSendEmailReq) (res *v1.CommonSendEmailRes, err error)
}
