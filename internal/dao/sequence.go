package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/db"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var Sequence = NewSequenceDao()

type SequenceDao struct {
	*MongoDB[entity.TalkRecords]
	cache *cache.Sequence
}

func NewSequenceDao(database ...string) *SequenceDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &SequenceDao{
		MongoDB: NewMongoDB[entity.TalkRecords](database[0], do.TALK_RECORDS_COLLECTION),
		cache:   cache.NewSequence(redis.Client),
	}
}

func (d *SequenceDao) try(ctx context.Context, userId int, receiverId int) error {

	result := d.cache.Redis().TTL(ctx, d.cache.Name(userId, receiverId)).Val()

	// 当数据不存在时需要从数据库中加载
	if result == time.Duration(-2) {

		lockName := fmt.Sprintf("%s_lock", d.cache.Name(userId, receiverId))

		isTrue := d.cache.Redis().SetNX(ctx, lockName, 1, 10*time.Second).Val()
		if !isTrue {
			return errors.New("请求频繁")
		}

		defer d.cache.Redis().Del(ctx, lockName)

		filter := bson.M{}
		// 检测UserId 是否被设置, 未设置则代表群聊
		if userId == 0 {
			filter["receiver_id"] = receiverId
			filter["talk_type"] = consts.ChatGroupMode
		} else {
			filter["$or"] = bson.A{
				bson.M{"user_id": userId, "receiver_id": receiverId},
				bson.M{"user_id": receiverId, "receiver_id": userId},
			}
		}

		talkRecords, err := d.FindOne(ctx, filter, "-sequence")
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			return err
		}

		if talkRecords != nil {
			if err := d.cache.Set(ctx, userId, receiverId, talkRecords.Sequence); err != nil {
				logger.Error(ctx, "[Sequence Set] 加载异常 err: ", err.Error())
				return err
			}
		}
	} else if result < time.Hour {
		d.cache.Redis().Expire(ctx, d.cache.Name(userId, receiverId), 12*time.Hour)
	}

	return nil
}

// Get 获取会话间的时序ID
func (d *SequenceDao) Get(ctx context.Context, userId int, receiverId int) int64 {

	if err := util.Retry(5, 100*time.Millisecond, func() error {
		return d.try(ctx, userId, receiverId)
	}); err != nil {
		logger.Error(ctx, "Sequence Get Err:", err)
	}

	return d.cache.Get(ctx, userId, receiverId)
}

// BatchGet 批量获取会话间的时序ID
func (d *SequenceDao) BatchGet(ctx context.Context, userId int, receiverId int, num int64) []int64 {

	if err := util.Retry(5, 100*time.Millisecond, func() error {
		return d.try(ctx, userId, receiverId)
	}); err != nil {
		logger.Errorf(ctx, "Sequence BatchGet Err:", err)
	}

	return d.cache.BatchGet(ctx, userId, receiverId, num)
}
