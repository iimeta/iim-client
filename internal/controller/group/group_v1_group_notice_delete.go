package group

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupNoticeDelete(ctx context.Context, req *v1.GroupNoticeDeleteReq) (res *v1.GroupNoticeDeleteRes, err error) {

	msg, err := service.GroupNotice().Delete(ctx, req.GroupNoticeDeleteReq)
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.WriteJson(model.DefaultResponse{
		Code:    200,
		Message: msg,
	})

	return
}
