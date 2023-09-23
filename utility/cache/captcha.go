package cache

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"time"

	"github.com/redis/go-redis/v9"
)

type CaptchaStorage struct {
	redis *redis.Client
}

func NewCaptchaStorage(redis *redis.Client) *CaptchaStorage {
	return &CaptchaStorage{redis: redis}
}

func (c *CaptchaStorage) Set(id string, value string) error {
	return c.redis.SetEx(gctx.New(), c.name(id), value, 3*time.Minute).Err()
}

func (c *CaptchaStorage) Get(id string, clear bool) string {
	value := c.redis.Get(gctx.New(), c.name(id)).Val()
	if clear && len(value) > 0 {
		c.redis.Del(gctx.New(), c.name(id))
	}

	return value
}

func (c *CaptchaStorage) Verify(id, answer string, clear bool) bool {
	value := c.redis.Get(gctx.New(), c.name(id)).Val()
	if clear && len(value) > 0 {
		c.redis.Del(gctx.New(), c.name(id))
	}

	return value == answer
}

func (c *CaptchaStorage) name(id string) string {
	return fmt.Sprintf("im:auth:captcha:%s", id)
}
