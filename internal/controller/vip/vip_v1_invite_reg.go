package vip

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/consts"
	"net/http"

	"github.com/iimeta/iim-client/api/vip/v1"
)

func (c *ControllerV1) InviteReg(ctx context.Context, req *v1.InviteRegReq) (res *v1.InviteRegRes, err error) {

	g.RequestFromCtx(ctx).Cookie.Set(consts.INVITE_CODE_COOKIE, req.Code)

	g.RequestFromCtx(ctx).Response.RedirectTo("/auth/register", http.StatusFound)

	return
}
