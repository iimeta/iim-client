package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupMute(ctx context.Context, req *v1.GroupMuteReq) (res *v1.GroupMuteRes, err error) {

	err = service.Group().Mute(ctx, req.GroupMuteReq)

	return
}
