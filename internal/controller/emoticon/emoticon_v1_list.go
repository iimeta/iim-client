package emoticon

import (
	"context"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/emoticon/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	collectEmoticons, err := service.Emoticon().CollectList(ctx)

	res = &v1.ListRes{
		ListRes: &model.ListRes{
			CollectEmoticons: collectEmoticons,
		},
	}

	return
}
