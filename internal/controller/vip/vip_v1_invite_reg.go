package vip

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"

	"github.com/iimeta/iim-client/api/vip/v1"
)

func (c *ControllerV1) InviteReg(ctx context.Context, req *v1.InviteRegReq) (res *v1.InviteRegRes, err error) {

	g.RequestFromCtx(ctx).Cookie.Set("invite_code", req.Code)

	g.RequestFromCtx(ctx).Response.RedirectTo("/auth/register", http.StatusFound)

	return
}
