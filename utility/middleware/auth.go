package middleware

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/iim-client/utility/jwt"
	"github.com/iimeta/iim-client/utility/logger"
	"net/http"
	"strconv"
	"strings"
)

const JWTSessionConst = "__JWT_SESSION__"
const UID_KEY = "uid"

var (
	ErrorNoLogin = errors.New("请登录后操作")
)

type IStorage interface {
	// IsBlackList 判断是否是黑名单
	IsBlackList(ctx context.Context, token string) bool
}

type JSession struct {
	Uid       int    `json:"uid"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

// Auth 授权中间件
func Auth(r *ghttp.Request, secret string, guard string, storage IStorage) {

	token := AuthHeaderToken(r)

	claims, err := verify(guard, secret, token)
	if err != nil {
		r.Response.Header().Set("Content-Type", "application/json")
		r.Response.WriteStatus(http.StatusUnauthorized, g.Map{"code": 401, "message": err.Error()})
		r.Exit()
		return
	}

	if storage.IsBlackList(r.Request.Context(), token) {
		r.Response.WriteStatus(http.StatusUnauthorized, g.Map{"code": 401, "message": "请登录再试"})
		r.Exit()
		return
	}

	uid, err := strconv.Atoi(claims.ID)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError, g.Map{"code": 500, "message": "解析 jwt 失败"})
		r.Exit()
		return
	}

	r.SetCtxVar(JWTSessionConst, &JSession{
		Uid:       uid,
		Token:     token,
		ExpiresAt: claims.ExpiresAt.Unix(),
	})

	r.SetCtxVar(UID_KEY, uid)

	if gstr.HasPrefix(r.GetHeader("Content-Type"), "application/json") {
		logger.Debugf(r.GetCtx(), "url: %s, request body: %s", r.GetUrl(), r.GetBodyString())
	} else {
		logger.Debugf(r.GetCtx(), "url: %s, Content-Type: %s", r.GetUrl(), r.GetHeader("Content-Type"))
	}

	r.Middleware.Next()
}

func AuthHeaderToken(r *ghttp.Request) string {

	token := r.GetHeader("Authorization")
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

	// Headers 中没有授权信息则读取 url 中的 token
	if token == "" {
		token = r.Get("token", "").String()
	}

	return token
}

func verify(guard string, secret string, token string) (*jwt.AuthClaims, error) {

	if token == "" {
		return nil, ErrorNoLogin
	}

	claims, err := jwt.ParseToken(token, secret)
	if err != nil {
		return nil, err
	}

	// 判断权限认证守卫是否一致
	if claims.Guard != guard || claims.Valid() != nil {
		return nil, ErrorNoLogin
	}

	return claims, nil
}
