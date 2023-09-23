package user

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/core"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/crypto"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// 注册
func (s *sUser) Register(ctx context.Context, register *model.UserRegister) (*model.User, error) {

	if dao.User.IsAccountExist(ctx, register.Account) {
		return nil, errors.New(register.Account + " 账号已存在")
	}

	salt := grand.Letters(8)

	user := &do.User{
		UserId:    core.IncrUserId(ctx),
		Email:     register.Account,
		Nickname:  register.Nickname,
		CreatedAt: gtime.Timestamp(),
	}

	uid, err := dao.User.Insert(ctx, user)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if _, err = dao.User.CreateAccount(ctx, &do.Account{
		Uid:      uid,
		UserId:   user.UserId,
		Account:  register.Account,
		Password: crypto.EncryptPassword(register.Password + salt),
		Salt:     salt,
		Status:   1,
	}); err != nil {
		return nil, err
	}

	return &model.User{
		UserId:    user.UserId,
		Email:     user.Email,
		Nickname:  user.Nickname,
		CreatedAt: user.CreatedAt,
	}, nil
}

// 登录
func (s *sUser) Login(ctx context.Context, account string, password string) (*model.User, error) {

	accountInfo, err := dao.User.FindAccount(ctx, account)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("账号或密码不正确")
		}
		logger.Error(ctx, err)
		return nil, err
	}

	if !crypto.VerifyPassword(accountInfo.Password, password+accountInfo.Salt) {
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

	return &model.User{
		UserId:    user.UserId,
		Mobile:    user.Mobile,
		Email:     user.Email,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Gender:    user.Gender,
		Motto:     user.Motto,
		Birthday:  user.Birthday,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// 找回密码
func (s *sUser) Forget(ctx context.Context, forget *model.UserForget) (bool, error) {

	account, err := dao.User.FindAccount(ctx, forget.Account)
	if err != nil || account.Id == "" {
		return false, errors.New(forget.Account + " 账号不存在")
	}

	if err = dao.User.ChangePasswordByUserId(ctx, account.UserId, forget.Password); err != nil {
		logger.Error(ctx, err)
		return false, errors.New("找回密码失败")
	}

	return true, nil
}

// 修改密码
func (s *sUser) UpdatePassword(ctx context.Context, uid int, oldPassword string, password string) error {

	user, err := dao.User.FindUserByUserId(ctx, uid)
	if err != nil || user.Id == "" {
		return errors.New("用户不存在")
	}

	account, err := dao.User.FindAccountByUserId(ctx, user.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("账号信息有误")
	}

	if !crypto.VerifyPassword(account.Password, oldPassword+account.Salt) {
		return errors.New("登录密码有误, 请重新输入")
	}

	if err = dao.User.ChangePasswordByUserId(ctx, uid, password); err != nil {
		logger.Error(ctx, err)
		return errors.New("修改密码失败")
	}

	return nil
}

// 用户详情
func (s *sUser) Detail(ctx context.Context) (*model.UserDetailRes, error) {

	user, err := dao.User.FindUserByUserId(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.UserDetailRes{
		Id:       user.UserId,
		Mobile:   user.Mobile,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Gender:   user.Gender,
		Motto:    user.Motto,
		Email:    user.Email,
		Birthday: user.Birthday,
	}, nil
}

// 用户设置
func (s *sUser) Setting(ctx context.Context) (*model.UserSettingRes, error) {

	user, err := dao.User.FindUserByUserId(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.UserSettingRes{
		UserInfo: &model.UserSettingResponse_UserInfo{
			Uid:      user.UserId,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Motto:    user.Motto,
			Gender:   user.Gender,
			Mobile:   user.Mobile,
			Email:    user.Email,
		},
		Setting: &model.UserSettingResponse_ConfigInfo{},
	}, nil
}

// 修改用户信息
func (s *sUser) ChangeDetail(ctx context.Context, params model.UserDetailUpdateReq) error {

	if params.Birthday != "" {
		if !util.IsDateFormat(params.Birthday) {
			return errors.New("birthday 格式错误")
		}
	}

	if err := dao.User.UpdateOne(ctx, bson.M{"user_id": service.Session().GetUid(ctx)}, &do.User{
		Nickname: strings.TrimSpace(strings.Replace(params.Nickname, " ", "", -1)),
		Avatar:   params.Avatar,
		Gender:   params.Gender,
		Motto:    params.Motto,
		Birthday: params.Birthday,
	}); err != nil {
		logger.Error(ctx, err)
		return errors.New("个人信息修改失败")
	}

	return nil
}

// 修改密码接口
func (s *sUser) ChangePassword(ctx context.Context, params model.UserPasswordUpdateReq) error {

	if err := s.UpdatePassword(ctx, service.Session().GetUid(ctx), params.OldPassword, params.NewPassword); err != nil {
		logger.Error(ctx, err)
		return errors.New("密码修改失败")
	}

	return nil
}

// 换绑手机号
func (s *sUser) ChangeMobile(ctx context.Context, params model.UserMobileUpdateReq) error {

	if !service.Sms().Verify(ctx, consts.CHANNEL_CHANGE_MOBILE, params.Mobile, params.Code) {
		return errors.New("短信验证码填写错误")
	}

	user, err := dao.User.FindUserByUserId(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	account, err := dao.User.FindAccountByUserId(ctx, user.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("账号信息有误")
	}

	if !crypto.VerifyPassword(account.Password, params.Password+account.Salt) {
		return errors.New("登录密码有误, 请重新输入")
	}

	if user.Mobile == params.Mobile {
		return errors.New("手机号与原手机号一致无需修改")
	}

	if dao.User.IsAccountExist(ctx, params.Mobile) {
		return errors.New(params.Mobile + " 手机号已被其它账号使用")
	}

	if err = dao.User.UpdateById(ctx, user.Id, bson.M{
		"mobile": params.Mobile,
	}); err != nil {
		logger.Error(ctx, err)
		return errors.New("手机号修改失败")
	}

	if account.Account == user.Mobile {
		if err = dao.User.ChangeAccountById(ctx, account.Id, params.Mobile); err != nil {
			logger.Error(ctx, err)
			return err
		}
	} else {

		accountInfo, err := dao.User.FindAccount(ctx, user.Mobile)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
			return err
		}

		if accountInfo != nil {
			if err = dao.User.ChangeAccountById(ctx, accountInfo.Id, params.Mobile); err != nil {
				logger.Error(ctx, err)
				return err
			}
		} else {
			if _, err := dao.User.CreateAccount(ctx, &do.Account{
				Uid:      account.Uid,
				UserId:   account.UserId,
				Account:  params.Mobile,
				Password: account.Password,
				Salt:     account.Salt,
				Status:   1,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 换绑邮箱
func (s *sUser) ChangeEmail(ctx context.Context, params model.UserEmailUpdateReq) error {

	if !service.Email().Verify(ctx, consts.CHANNEL_CHANGE_EMAIL, params.Email, params.Code) {
		return errors.New("邮件验证码填写错误")
	}

	user, err := dao.User.FindUserByUserId(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	account, err := dao.User.FindAccountByUserId(ctx, user.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("账号信息有误")
	}

	if !crypto.VerifyPassword(account.Password, params.Password+account.Salt) {
		return errors.New("登录密码有误, 请重新输入")
	}

	if user.Email == params.Email {
		return errors.New("邮箱与原邮箱一致无需修改")
	}

	if dao.User.IsAccountExist(ctx, params.Email) {
		return errors.New(params.Email + " 邮箱已被其它账号使用")
	}

	if err = dao.User.UpdateById(ctx, user.Id, bson.M{
		"email": params.Email,
	}); err != nil {
		logger.Error(ctx, err)
		return errors.New("邮箱修改失败")
	}

	if account.Account == user.Email {
		if err = dao.User.ChangeAccountById(ctx, account.Id, params.Email); err != nil {
			logger.Error(ctx, err)
			return err
		}
	} else {

		accountInfo, err := dao.User.FindAccount(ctx, user.Email)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
			return err
		}

		if accountInfo != nil {
			if err = dao.User.ChangeAccountById(ctx, accountInfo.Id, params.Email); err != nil {
				logger.Error(ctx, err)
				return err
			}
		} else {
			if _, err := dao.User.CreateAccount(ctx, &do.Account{
				Uid:      account.Uid,
				UserId:   account.UserId,
				Account:  params.Email,
				Password: account.Password,
				Salt:     account.Salt,
				Status:   1,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 根据userId获取用户信息
func (s *sUser) GetUserById(ctx context.Context, userId int) (*model.User, error) {

	user, err := dao.User.FindUserByUserId(ctx, userId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.User{
		Id:        user.Id,
		UserId:    user.UserId,
		Mobile:    user.Mobile,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Gender:    user.Gender,
		Motto:     user.Motto,
		Email:     user.Email,
		Birthday:  user.Birthday,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
