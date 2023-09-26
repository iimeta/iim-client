package note

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/note/v1"
)

func (c *ControllerV1) TagList(ctx context.Context, req *v1.TagListReq) (res *v1.TagListRes, err error) {

	tagListRes, err := service.NoteTag().List(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.TagListRes{
		TagListRes: tagListRes,
	}

	return
}
