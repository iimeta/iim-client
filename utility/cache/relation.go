package cache

import (
	"context"
	"fmt"
	redis2 "github.com/iimeta/iim-client/utility/redis"
	"time"

	"github.com/redis/go-redis/v9"
)

type Relation struct {
	redis *redis.Client
}

func NewRelation(redis *redis.Client) *Relation {
	return &Relation{redis: redis}
}

func (r *Relation) IsContactRelation(ctx context.Context, uid, uid2 int) error {
	return r.redis.Get(ctx, r.keyContactRelation(uid, uid2)).Err()
}

func (r *Relation) SetContactRelation(ctx context.Context, uid, uid2 int) {
	r.redis.SetEx(ctx, r.keyContactRelation(uid, uid2), "1", time.Hour*1)
}

func (r *Relation) DelContactRelation(ctx context.Context, uid, uid2 int) {
	r.redis.Del(ctx, r.keyContactRelation(uid, uid2))
}

func IsGroupRelation(ctx context.Context, uid, gid int) bool {

	replay, err := redis2.Get(ctx, keyGroupRelation(uid, gid))
	if err != nil {
		return false
	}

	return !replay.IsEmpty()
}

func SetGroupRelation(ctx context.Context, uid, gid int) {
	_ = redis2.SetEX(ctx, keyGroupRelation(uid, gid), "1", int64(time.Hour.Seconds()))
}

func (r *Relation) DelGroupRelation(ctx context.Context, uid, gid int) {
	r.redis.Del(ctx, keyGroupRelation(uid, gid))
}

func (r *Relation) BatchDelGroupRelation(ctx context.Context, uids []int, gid int) {
	for _, uid := range uids {
		r.DelGroupRelation(ctx, uid, gid)
	}
}

func (r *Relation) keyContactRelation(uid, uid2 int) string {
	if uid2 < uid {
		uid, uid2 = uid2, uid
	}

	return fmt.Sprintf("im:contact:relation:%d_%d", uid, uid2)
}

func keyGroupRelation(uid, gid int) string {
	return fmt.Sprintf("im:contact:relation:%d_%d", uid, gid)
}
