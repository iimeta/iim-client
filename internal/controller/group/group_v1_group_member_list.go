package group

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/group/v1"
)

func (c *ControllerV1) GroupMemberList(ctx context.Context, req *v1.GroupMemberListReq) (res *v1.GroupMemberListRes, err error) {

	groupMemberListRes, err := service.Group().Members(ctx, req.GroupMemberListReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GroupMemberListRes{
		GroupMemberListRes: groupMemberListRes,
	}

	return
}
