package vip

import (
	"context"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/vip/v1"
)

func (c *ControllerV1) InviteFriends(ctx context.Context, req *v1.InviteFriendsReq) (res *v1.InviteFriendsRes, err error) {

	inviteUrl, records, err := service.Vip().InviteFriends(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.InviteFriendsRes{
		InviteFriendsRes: &model.InviteFriendsRes{
			InviteUrl: inviteUrl,
			Items:     records,
		},
	}

	return
}
