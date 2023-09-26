package group

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) NoticeDelete(ctx context.Context, req *v1.NoticeDeleteReq) (res *v1.NoticeDeleteRes, err error) {

	msg, err := service.GroupNotice().Delete(ctx, req.NoticeDeleteReq)
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.WriteJson(model.DefaultResponse{
		Code:    200,
		Message: msg,
	})

	return
}
