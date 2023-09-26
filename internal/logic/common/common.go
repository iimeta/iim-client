package common

import (
	"context"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
)

type sCommon struct{}

func init() {
	service.RegisterCommon(New())
}

func New() service.ICommon {
	return &sCommon{}
}

// 发送短信验证码
func (s *sCommon) SmsCode(ctx context.Context, params model.SendSmsReq) (*model.SendSmsRes, error) {

	switch params.Channel {
	// 需要判断账号是否存在
	case consts.CHANNEL_LOGIN, consts.CHANNEL_FORGET_ACCOUNT:
		if !dao.User.IsAccountExist(ctx, params.Mobile) {
			return nil, errors.New("账号不存在或密码错误")
		}

	// 需要判断账号是否存在
	case consts.CHANNEL_REGISTER, consts.CHANNEL_CHANGE_MOBILE:
		if dao.User.IsAccountExist(ctx, params.Mobile) {
			return nil, errors.New("手机号已被他人使用")
		}

	default:
		return nil, errors.New("发送异常")
	}

	// 发送短信验证码
	code, err := service.Sms().Send(ctx, params.Channel, params.Mobile)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if config.Cfg.App.Debug {
		return &model.SendSmsRes{
			IsDebug: true,
			SmsCode: code,
		}, nil
	}

	return nil, nil
}

// 发送邮件验证码
func (s *sCommon) EmailCode(ctx context.Context, params model.SendEmailReq) (*model.SendEmailRes, error) {

	switch params.Channel {
	// 需要判断账号是否存在
	case consts.CHANNEL_LOGIN:
		if !dao.User.IsAccountExist(ctx, params.Email) {
			return nil, errors.New("账号不存在或密码错误")
		}

	// 需要判断账号是否存在
	case consts.CHANNEL_FORGET_ACCOUNT:
		if !dao.User.IsAccountExist(ctx, params.Email) {
			return nil, errors.New("账号不存在")
		}

	// 需要判断账号是否存在
	case consts.CHANNEL_REGISTER, consts.CHANNEL_CHANGE_EMAIL:
		if dao.User.IsAccountExist(ctx, params.Email) {
			return nil, errors.New("邮箱已被他人使用")
		}

	default:
		return nil, errors.New("发送异常")
	}

	// 发送邮件验证码
	code, err := service.Email().Send(ctx, params.Channel, params.Email)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if config.Cfg.App.Debug {
		return &model.SendEmailRes{
			IsDebug: true,
			Code:    code,
		}, nil
	}

	return nil, nil
}

// 公共设置
func (s *sCommon) Setting(ctx context.Context) error {
	return nil
}
