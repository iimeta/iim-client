package emoticon

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/emoticon/v1"
)

func (c *ControllerV1) SysList(ctx context.Context, req *v1.SysListReq) (res *v1.SysListRes, err error) {

	sysListRes, err := service.Emoticon().SystemList(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.SysListRes{}
	res.Items = sysListRes

	return
}
