package vip

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
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
	date := tomorrow.Format("20060102")

	logger.Infof(ctx, "InitDailyUsage date: %s start", date)

	now := gtime.Now().UnixMilli()
	defer func() {
		logger.Infof(ctx, "InitDailyUsage date: %s end totalTime: %d ms", date, gtime.Now().UnixMilli()-now)
	}()

	filter := bson.M{
		"created_at": bson.M{
			"$lte": tomorrow.Add(-(config.Cfg.Vip.RegisteredDays * gtime.D)).EndOfDay().Unix(),
		},
	}

	userList, err := dao.User.Find(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	for _, user := range userList {
		_, err = redis.HSet(ctx, s.GenerateUidUsageKey(ctx, user.UserId, date), g.MapStrAny{consts.TOTAL_TOKENS_FIELD: config.Cfg.Vip.Daily.FreeTokens})
		if err != nil {
			logger.Error(ctx, err)
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

	uidUsageKey := s.GenerateUidUsageKey(ctx, user.UserId, gtime.Now().Format("20060102"))

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

	vipInfo := &model.VipInfo{
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
