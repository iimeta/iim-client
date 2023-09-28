package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) NoticeEdit(ctx context.Context, req *v1.NoticeEditReq) (res *v1.NoticeEditRes, err error) {

	err = service.GroupNotice().CreateAndUpdate(ctx, req.NoticeEditReq)

	return
}
