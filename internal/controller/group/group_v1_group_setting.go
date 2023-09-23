package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupSetting(ctx context.Context, req *v1.GroupSettingReq) (res *v1.GroupSettingRes, err error) {

	err = service.Group().Setting(ctx, req.GroupSettingReq)

	return
}
