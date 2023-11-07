package vip

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type sVip struct{}

func init() {

	service.RegisterVip(New())

	ctx := gctx.New()
	_, err := gcron.AddSingleton(ctx, config.Cfg.Vip.InitDailyCron, service.Vip().InitDailyUsage, "InitDailyUsage")
	if err != nil {
		logger.Error(ctx, err)
	}
}

func New() service.IVip {
	return &sVip{}
}

func (s *sVip) InitDailyUsage(ctx context.Context) {

	tomorrow := gtime.Now().Add(gtime.D)
	date := tomorrow.Time.Format(util.DateDayFormat)

	logger.Infof(ctx, "InitDailyUsage date: %s start", date)

	now := gtime.Now().UnixMilli()
	defer func() {
		logger.Infof(ctx, "InitDailyUsage date: %s end totalTime: %d ms", date, gtime.Now().UnixMilli()-now)
	}()

	vips, err := dao.Vip.Find(ctx, bson.M{}, "level")
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	for _, vip := range vips {

		filter := bson.M{
			"vip_level": bson.M{
				"$lte": vip.Level,
			},
		}

		if vip.Rule.RegDays > 0 {
			filter["created_at"] = bson.M{
				"$lte": tomorrow.Add(-(time.Duration(vip.Rule.RegDays) * gtime.D)).EndOfDay().Unix(),
			}
		}

		userList, err := dao.User.Find(ctx, filter)
		if err != nil {
			logger.Error(ctx, err)
			return
		}

		for _, user := range userList {

			if user.VipLevel < vip.Level {

				filter := bson.M{
					"inviter": user.UserId,
				}

				if vip.Rule.InviteRegDays > 0 {
					filter["created_at"] = bson.M{
						"$lte": tomorrow.Add(-(time.Duration(vip.Rule.InviteRegDays) * gtime.D)).EndOfDay().Unix(),
					}
				}

				count, err := dao.Invite.CountDocuments(ctx, filter)
				if err != nil {
					logger.Error(ctx, err)
					continue
				}

				if int(count) >= vip.Rule.InviteNum {
					if err := dao.User.UpdateOne(ctx, bson.M{"user_id": user.UserId}, bson.M{
						"vip_level": vip.Level,
					}); err != nil {
						logger.Error(ctx, err)
					}
					user.VipLevel = vip.Level
				}
			}

			if user.VipLevel >= vip.Level {

				if vip.Rule.OnlineTime > 0 {

					firstTime, err := redis.HGetInt(ctx, fmt.Sprintf(consts.USER_TIME_LOGIN_KEY, util.DateNumber(), user.UserId), consts.FIRST_TIME_FIELD)
					if err != nil {
						logger.Error(ctx, err)
						continue
					}

					lastTime, err := redis.HGetInt(ctx, fmt.Sprintf(consts.USER_TIME_LOGIN_KEY, util.DateNumber(), user.UserId), consts.LAST_TIME_FIELD)
					if err != nil {
						logger.Error(ctx, err)
						continue
					}

					if lastTime-firstTime < int((time.Duration(vip.Rule.OnlineTime) * time.Minute).Seconds()) {
						continue
					}
				}

				if _, err = redis.HSet(ctx, s.GenerateUidUsageKey(ctx, user.UserId, date), g.MapStrAny{consts.TOTAL_TOKENS_FIELD: vip.FreeTokens}); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}
}

func (s *sVip) GenerateUidUsageKey(ctx context.Context, uid int, date string) string {
	return fmt.Sprintf(consts.UID_USAGE_KEY, uid, date)
}

func (s *sVip) GenerateSecretKey(ctx context.Context) (string, error) {

	uid := service.Session().GetUid(ctx)

	secretKey := util.NewSecretKey(uid, 48, "IIM")

	if err := dao.User.UpdateOne(ctx, bson.M{"user_id": uid}, bson.M{
		"secret_key": secretKey,
	}); err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	return secretKey, nil
}

func (s *sVip) VipInfo(ctx context.Context) (*model.VipInfo, error) {

	user, err := dao.User.FindUserByUserId(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	uidUsageKey := s.GenerateUidUsageKey(ctx, user.UserId, gtime.Now().Time.Format(util.DateDayFormat))

	usageCount, err := redis.HGetInt(ctx, uidUsageKey, consts.USAGE_COUNT_FIELD)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	usedTokens, err := redis.HGetInt(ctx, uidUsageKey, consts.USED_TOKENS_FIELD)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	totalTokens, err := redis.HGetInt(ctx, uidUsageKey, consts.TOTAL_TOKENS_FIELD)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	vipName := ""
	vip, err := dao.Vip.FindOne(ctx, bson.M{"level": user.VipLevel})
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(ctx, err)
		return nil, err
	}

	if vip != nil {
		vipName = vip.Name
	}

	vipInfo := &model.VipInfo{
		VipName:     vipName,
		UserId:      user.UserId,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
		SecretKey:   user.SecretKey,
		RegTime:     util.FormatDatetime(user.CreatedAt),
		UsageCount:  usageCount,
		UsedTokens:  usedTokens,
		TotalTokens: totalTokens,
	}

	return vipInfo, nil
}

func (s *sVip) Vips(ctx context.Context) ([]*model.Vip, error) {

	vips, err := dao.Vip.Find(ctx, bson.M{}, "level")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	vipList := make([]*model.Vip, 0)
	for _, vip := range vips {
		vipList = append(vipList, &model.Vip{
			Level:       vip.Level,
			Name:        vip.Name,
			Models:      vip.Models,
			FreeTokens:  vip.FreeTokens,
			MinuteLimit: vip.MinuteLimit,
			DailyLimit:  vip.DailyLimit,
			Remark:      vip.Remark,
			Status:      vip.Status,
			CreatedAt:   vip.CreatedAt,
			UpdatedAt:   vip.UpdatedAt,
		})
	}

	return vipList, err
}

func (s *sVip) InviteFriends(ctx context.Context) (string, []*model.InviteRecord, error) {

	inviteUrl := "/invite/" + s.InviteCode(ctx)

	inviteRecords, err := dao.Invite.Find(ctx, bson.M{"inviter": service.Session().GetUid(ctx)}, "-created_at")
	if err != nil {
		return inviteUrl, nil, err
	}

	items := make([]*model.InviteRecord, 0)
	for _, inviteRecord := range inviteRecords {
		items = append(items, &model.InviteRecord{
			Nickname:  inviteRecord.Nickname,
			Email:     inviteRecord.Email,
			CreatedAt: util.FormatDatetime(inviteRecord.CreatedAt),
		})
	}

	return inviteUrl, items, nil
}

func (s *sVip) SaveInviteRecord(ctx context.Context, user *do.User) error {

	inviteCode := g.RequestFromCtx(ctx).Cookie.Get(consts.INVITE_CODE_COOKIE)
	if inviteCode == nil || inviteCode.String() == "" {
		return nil
	}

	inviter := s.InviteCodeToUid(ctx, inviteCode.String())

	if _, err := dao.Invite.Insert(ctx, do.InviteRecord{
		Nickname:  user.Nickname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Inviter:   inviter,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	g.RequestFromCtx(ctx).Cookie.SetHttpCookie(&http.Cookie{
		Domain: "*",
		Path:   "/",
		Name:   consts.INVITE_CODE_COOKIE,
		MaxAge: -1,
	})

	return nil
}

func (s *sVip) InviteCode(ctx context.Context) string {

	inviteCode := ""
	for i, v := range gconv.String(service.Session().GetUid(ctx)) {
		inviteCode += fmt.Sprintf("%c", int32(i)+48+v)
	}

	return inviteCode
}

func (s *sVip) InviteCodeToUid(ctx context.Context, inviteCode string) int {

	uid := ""
	for i, v := range inviteCode {
		uid += fmt.Sprintf("%c", v-48-int32(i))
	}

	return gconv.Int(uid)
}
