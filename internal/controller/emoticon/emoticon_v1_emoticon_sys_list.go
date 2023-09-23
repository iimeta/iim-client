package emoticon

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/emoticon/v1"
)

func (c *ControllerV1) EmoticonSysList(ctx context.Context, req *v1.EmoticonSysListReq) (res *v1.EmoticonSysListRes, err error) {

	emoticonSysListRes, err := service.Emoticon().SystemList(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.EmoticonSysListRes{}
	res.Items = emoticonSysListRes

	return
}
