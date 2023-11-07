package cmd

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/controller/auth"
	"github.com/iimeta/iim-client/internal/controller/common"
	"github.com/iimeta/iim-client/internal/controller/contact"
	"github.com/iimeta/iim-client/internal/controller/emoticon"
	"github.com/iimeta/iim-client/internal/controller/file"
	"github.com/iimeta/iim-client/internal/controller/group"
	"github.com/iimeta/iim-client/internal/controller/note"
	"github.com/iimeta/iim-client/internal/controller/talk"
	"github.com/iimeta/iim-client/internal/controller/user"
	"github.com/iimeta/iim-client/internal/controller/vip"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/middleware"
	"github.com/iimeta/iim-client/utility/redis"
	"net/http"

	_ "github.com/iimeta/iim-server/server"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			s := g.Server()

			s.BindHookHandler("/*", ghttp.HookBeforeServe, beforeServeHook)

			s.SetServerRoot("./resource/iim-web/")

			s.AddStaticPath("/auth", "./resource/iim-web")
			s.AddStaticPath("/auth/login", "./resource/iim-web")
			s.AddStaticPath("/auth/forget", "./resource/iim-web")
			s.AddStaticPath("/auth/register", "./resource/iim-web")
			s.AddStaticPath("/message", "./resource/iim-web")
			s.AddStaticPath("/contact", "./resource/iim-web")
			s.AddStaticPath("/contact/apply", "./resource/iim-web")
			s.AddStaticPath("/contact/friend", "./resource/iim-web")
			s.AddStaticPath("/contact/group", "./resource/iim-web")
			s.AddStaticPath("/contact/group/open", "./resource/iim-web")
			s.AddStaticPath("/settings", "./resource/iim-web")
			s.AddStaticPath("/settings/detail", "./resource/iim-web")
			s.AddStaticPath("/settings/security", "./resource/iim-web")
			s.AddStaticPath("/settings/personalize", "./resource/iim-web")
			s.AddStaticPath("/settings/notification", "./resource/iim-web")
			s.AddStaticPath("/settings/binding", "./resource/iim-web")
			s.AddStaticPath("/settings/apply", "./resource/iim-web")
			s.AddStaticPath("/note", "./resource/iim-web")
			s.AddStaticPath("/vip/info", "./resource/iim-web")
			s.AddStaticPath("/vip/vip", "./resource/iim-web")
			s.AddStaticPath("/vip/invite", "./resource/iim-web")

			s.AddStaticPath("/public", "./resource/public")

			s.Group("/", func(g *ghttp.RouterGroup) {
				g.Middleware(MiddlewareAuth)
				g.Middleware(MiddlewareHandlerResponse)
				g.Bind()
			})

			s.Group("/invite/:code", func(g *ghttp.RouterGroup) {
				g.Middleware(MiddlewareHandlerResponse)
				g.Bind(
					vip.NewV1(),
				)
			})

			s.Group("/api/v1", func(v1 *ghttp.RouterGroup) {

				v1.Middleware(MiddlewareHandlerResponse)

				v1.Group("/common", func(g *ghttp.RouterGroup) {
					g.Bind(
						common.NewV1(),
					)
				})

				v1.Group("/auth", func(g *ghttp.RouterGroup) {
					g.Bind(
						auth.NewV1(),
					)
				})

				v1.Group("/users", func(g *ghttp.RouterGroup) {
					g.Middleware(MiddlewareAuth)
					g.Bind(
						user.NewV1(),
					)
				})

				v1.Group("/contact", func(g *ghttp.RouterGroup) {
					g.Middleware(MiddlewareAuth)
					g.Bind(
						contact.NewV1(),
					)
				})

				v1.Group("/group", func(g *ghttp.RouterGroup) {
					g.Middleware(MiddlewareAuth)
					g.Bind(
						group.NewV1(),
					)
				})

				v1.Group("/talk", func(g *ghttp.RouterGroup) {
					g.Middleware(MiddlewareAuth)
					g.Bind(
						talk.NewV1(),
					)
				})

				v1.Group("/emoticon", func(g *ghttp.RouterGroup) {
					g.Middleware(MiddlewareAuth)
					g.Bind(
						emoticon.NewV1(),
					)
				})

				v1.Group("/upload", func(g *ghttp.RouterGroup) {
					g.Middleware(MiddlewareAuth)
					g.Bind(
						file.NewV1(),
					)
				})

				v1.Group("/note", func(g *ghttp.RouterGroup) {
					g.Middleware(MiddlewareAuth)
					g.Bind(
						note.NewV1(),
					)
				})

				v1.Group("/vip", func(g *ghttp.RouterGroup) {
					g.Middleware(MiddlewareAuth)
					g.Bind(
						vip.NewV1(),
					)
				})
			})

			s.Run()
			return nil
		},
	}
)

func beforeServeHook(r *ghttp.Request) {
	logger.Debugf(r.GetCtx(), "beforeServeHook [isFile: %v] URI: %s", r.IsFileRequest(), r.RequestURI)
	r.Response.CORSDefault()
}

func MiddlewareAuth(r *ghttp.Request) {
	middleware.Auth(r, config.Cfg.Jwt.Secret, "api", cache.NewTokenSessionStorage(redis.Client))
}

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func MiddlewareHandlerResponse(r *ghttp.Request) {

	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.New(200, "success", "success")
			msg = code.Message()
		}
	}

	data := DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	}

	logger.Debugf(r.GetCtx(), "url: %s, response body: %s", r.GetUrl(), gjson.MustEncodeString(data))

	r.Response.WriteJson(data)
}
