package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ContactApplyUnreadNum(ctx context.Context, req *v1.ContactApplyUnreadNumReq) (res *v1.ContactApplyUnreadNumRes, err error) {

	applyUnreadNum, err := service.Apply().ApplyUnreadNum(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ContactApplyUnreadNumRes{
		UnreadNum: applyUnreadNum,
	}

	return
}
