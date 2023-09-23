package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupNoSpeak(ctx context.Context, req *v1.GroupNoSpeakReq) (res *v1.GroupNoSpeakRes, err error) {

	err = service.Group().NoSpeak(ctx, req.GroupNoSpeakReq)

	return
}
