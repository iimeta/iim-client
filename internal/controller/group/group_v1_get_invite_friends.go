package group

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GetInviteFriends(ctx context.Context, req *v1.GetInviteFriendsReq) (res *v1.GetInviteFriendsRes, err error) {

	contactListItem, err := service.Group().GetInviteFriends(ctx, req.GetInviteFriendsReq)
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.WriteJson(model.DefaultResponse{
		Code:    200,
		Message: "success",
		Data:    contactListItem,
	})

	return
}
