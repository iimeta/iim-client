package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) NoticeDelete(ctx context.Context, req *v1.NoticeDeleteReq) (res *v1.NoticeDeleteRes, err error) {

	err = service.GroupNotice().Delete(ctx, req.NoticeDeleteReq)

	return
}
