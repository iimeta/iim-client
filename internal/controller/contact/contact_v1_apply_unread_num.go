package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) ApplyUnreadNum(ctx context.Context, req *v1.ApplyUnreadNumReq) (res *v1.ApplyUnreadNumRes, err error) {

	applyUnreadNum, err := service.ContactApply().ApplyUnreadNum(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ApplyUnreadNumRes{}
	res.UnreadNum = applyUnreadNum

	return
}
