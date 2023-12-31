package auth

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/core"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/crypto"
	"github.com/iimeta/iim-client/utility/jwt"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/middleware"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{}
}

// 注册接口
func (s *sAuth) Register(ctx context.Context, params model.RegisterReq) error {

	// 验证验证码是否正确
	if !service.Email().Verify(ctx, consts.CHANNEL_REGISTER, params.Account, params.Code) {
		return errors.New("验证码填写错误")
	}

	if dao.User.IsAccountExist(ctx, params.Account) {
		return errors.New(params.Account + " 账号已存在")
	}

	salt := grand.Letters(8)

	user := &do.User{
		UserId:    core.IncrUserId(ctx),
		Email:     params.Account,
		Nickname:  params.Nickname,
		CreatedAt: gtime.Timestamp(),
	}

	uid, err := dao.User.Insert(ctx, user)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.User.CreateAccount(ctx, &do.Account{
		Uid:      uid,
		UserId:   user.UserId,
		Account:  params.Account,
		Password: crypto.EncryptPassword(params.Password + salt),
		Salt:     salt,
		Status:   1,
	}); err != nil {
		return err
	}

	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	service.Email().Delete(ctx, consts.CHANNEL_REGISTER, params.Account)

	// 保存邀请记录
	if err = service.Vip().SaveInviteRecord(ctx, user); err != nil {
		logger.Error(ctx, err)
	}

	////////////////////自动申请添加好友和自动通过//////////////////// todo
	value, err := config.Get(ctx, "register.auto_add_uids")
	if err == nil && value != nil && len(value.Ints()) > 0 {

		ctx = context.WithValue(ctx, middleware.UID_KEY, user.UserId)

		for _, uid := range value.Ints() {
			if applyId, err := service.ContactApply().Create(ctx, model.ApplyCreateReq{
				ContactApply: model.ContactApply{
					UserId:   user.UserId,
					Remark:   user.Email,
					FriendId: uid,
				},
			}); err == nil {
				_, _ = service.ContactApply().Accept(ctx, model.ApplyAcceptReq{
					ContactApply: model.ContactApply{
						Remark:  user.Nickname,
						ApplyId: applyId,
						UserId:  uid,
					},
				})
			}
		}
	}

	return nil
}

// 登录接口
func (s *sAuth) Login(ctx context.Context, params model.LoginReq) (res *model.LoginRes, err error) {

	defer func() {
		if err != nil {
			val, _ := redis.Incr(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account))
			if val == 1 {
				_, _ = redis.Expire(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account), 30*60) // 锁定30分钟
			}
		} else {
			_, _ = redis.Del(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account))
		}
	}()

	val, err := redis.GetInt(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account))
	if err == nil && val >= 5 {
		return nil, errors.New("登录失败次数过多, 请稍后再试")
	}

	accountInfo, err := dao.User.FindAccount(ctx, params.Account)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("账号或密码不正确")
		}
		logger.Error(ctx, err)
		return nil, err
	}

	if !crypto.VerifyPassword(accountInfo.Password, params.Password+accountInfo.Salt) {
		return nil, errors.New("账号或密码不正确")
	}

	user, err := dao.User.FindById(ctx, accountInfo.Uid)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("用户不存在或已被禁用") // todo
		}
		logger.Error(ctx, err)
		return nil, err
	}

	loginRobot, err := dao.Robot.GetLoginRobot(ctx)
	if err != nil {
		logger.Error(ctx, err)
	}

	ip := g.RequestFromCtx(ctx).GetClientIp()
	if loginRobot != nil {

		address, err := util.FindAddress(ctx, ip)
		if err != nil {
			logger.Error(ctx, err)
		}

		if _, err = dao.TalkSession.Create(ctx, &do.TalkSessionCreate{
			UserId:     user.UserId,
			TalkType:   consts.ChatPrivateMode,
			ReceiverId: loginRobot.UserId,
			IsRobot:    1,
			IsTalk:     loginRobot.IsTalk,
		}); err != nil {
			logger.Error(ctx, err)
		}

		// 推送登录消息
		if err = service.TalkMessage().SendNoticeMessage(ctx, &model.NoticeMessage{
			TalkType: consts.ChatPrivateMode,
			MsgType:  consts.MsgNoticeLogin,
			Sender: &model.Sender{
				Id: loginRobot.UserId,
			},
			Receiver: &model.Receiver{
				Id:         user.UserId,
				ReceiverId: user.UserId,
			},
			Login: &model.Login{
				IP:       ip,
				Platform: params.Platform,
				Agent:    g.RequestFromCtx(ctx).GetHeader("user-agent"),
				Address:  address,
				Reason:   "常用设备登录",
				Datetime: gtime.Datetime(),
			},
		}); err != nil {
			logger.Error(ctx, err)
		}
	}

	// 记录登录ip和时间
	if err = dao.Account.UpdateById(ctx, accountInfo.Id, bson.M{
		"last_login_ip":   ip,
		"last_login_time": gtime.Timestamp(),
	}); err != nil {
		logger.Error(ctx, err)
	}

	return &model.LoginRes{
		Type:        "Bearer",
		AccessToken: token(user.UserId),
		ExpiresIn:   int(config.Cfg.Jwt.ExpiresTime),
	}, nil
}

// 退出登录接口
func (s *sAuth) Logout(ctx context.Context) error {

	toBlackList(ctx)

	return nil
}

// 账号找回接口
func (s *sAuth) Forget(ctx context.Context, params model.ForgetReq) error {

	// 验证验证码是否正确
	if !service.Email().Verify(ctx, consts.CHANNEL_FORGET_ACCOUNT, params.Account, params.Code) {
		return errors.New("验证码填写错误")
	}

	account, err := dao.User.FindAccount(ctx, params.Account)
	if err != nil || account.Id == "" {
		return errors.New(params.Account + " 账号不存在")
	}

	if err = dao.User.ChangePasswordByUserId(ctx, account.UserId, params.Password); err != nil {
		logger.Error(ctx, err)
		return errors.New("找回密码失败")
	}

	service.Email().Delete(ctx, consts.CHANNEL_FORGET_ACCOUNT, params.Account)

	return nil
}

// Token 刷新接口
func (s *sAuth) Refresh(ctx context.Context) (*model.RefreshRes, error) {

	toBlackList(ctx)

	return &model.RefreshRes{
		Type:        "Bearer",
		AccessToken: token(service.Session().GetUid(ctx)),
		ExpiresIn:   int(config.Cfg.Jwt.ExpiresTime),
	}, nil
}

func token(uid int) string {

	expiresAt := time.Now().Add(time.Second * time.Duration(config.Cfg.Jwt.ExpiresTime))

	// 生成登录凭证
	token := jwt.GenerateToken("api", config.Cfg.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID:        strconv.Itoa(uid),
		Issuer:    "iim.web",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	return token
}

// 设置黑名单
func toBlackList(ctx context.Context) {
	data := ctx.Value(middleware.JWTSessionConst)
	if data != nil {
		session := data.(*middleware.JSession)
		if ex := session.ExpiresAt - time.Now().Unix(); ex > 0 {
			_ = cache.NewTokenSessionStorage(redis.Client).SetBlackList(ctx, session.Token, time.Duration(ex)*time.Second)
		}
	}
}
