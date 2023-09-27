package session

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
)

type sSession struct{}

func init() {
	service.RegisterSession(New())
}

func New() service.ISession {
	return &sSession{}
}

// 获取会话中UserId
func (s *sSession) GetUid(ctx context.Context) int {
	return ctx.Value("uid").(int)
}

// 获取会话中用户信息
func (s *sSession) GetUser(ctx context.Context) *model.User {

	value := ctx.Value("user")
	if value != nil {
		return value.(*model.User)
	}

	user, err := service.User().GetUserById(ctx, s.GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil
	}

	// todo 应该没有用
	g.RequestFromCtx(ctx).SetCtxVar("user", user)

	return user
}
