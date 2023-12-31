package contact

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/iimeta/iim-client/api/contact/v1"
)

func (c *ControllerV1) GroupCreate(ctx context.Context, req *v1.GroupCreateReq) (res *v1.GroupCreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
