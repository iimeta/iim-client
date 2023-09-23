package group

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupNoticeEdit(ctx context.Context, req *v1.GroupNoticeEditReq) (res *v1.GroupNoticeEditRes, err error) {

	msg, err := service.GroupNotice().CreateAndUpdate(ctx, req.GroupNoticeEditReq)
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.WriteJson(model.DefaultResponse{
		Code:    200,
		Message: msg,
	})

	return
}
