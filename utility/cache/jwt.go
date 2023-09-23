package cache

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"time"

	"github.com/redis/go-redis/v9"
)

type JwtTokenStorage struct {
	redis *redis.Client
}

func NewTokenSessionStorage(redis *redis.Client) *JwtTokenStorage {
	return &JwtTokenStorage{redis}
}

func (s *JwtTokenStorage) SetBlackList(ctx context.Context, token string, exp time.Duration) error {
	return s.redis.Set(ctx, s.name(token), 1, exp).Err()
}

func (s *JwtTokenStorage) IsBlackList(ctx context.Context, token string) bool {
	return s.redis.Get(ctx, s.name(token)).Val() != ""
}

func (s *JwtTokenStorage) name(token string) string {
	return fmt.Sprintf("jwt:blacklist:%s", gmd5.MustEncryptString(token))
}
