package cache

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	VoteUsersCache     = "talk:vote:answer-users:%s"
	VoteStatisticCache = "talk:vote:statistic:%s"
)

type Vote struct {
	redis *redis.Client
}

func NewVote(rds *redis.Client) *Vote {
	return &Vote{redis: rds}
}

func (t *Vote) GetVoteAnswerUser(ctx context.Context, voteId string) ([]int, error) {
	val, err := t.redis.Get(ctx, fmt.Sprintf(VoteUsersCache, voteId)).Result()

	if err != nil {
		return nil, err
	}

	var ids []int
	if err := gjson.Unmarshal([]byte(val), &ids); err != nil {
		return nil, err
	}

	return ids, nil
}

func (t *Vote) SetVoteAnswerUser(ctx context.Context, vid string, uids []int) error {
	return t.redis.Set(ctx, fmt.Sprintf(VoteUsersCache, vid), gjson.MustEncodeString(uids), time.Hour*24).Err()
}

func (t *Vote) GetVoteStatistics(ctx context.Context, vid string) (string, error) {
	return t.redis.Get(ctx, fmt.Sprintf(VoteStatisticCache, vid)).Result()
}

func (t *Vote) SetVoteStatistics(ctx context.Context, vid, value string) error {
	return t.redis.Set(ctx, fmt.Sprintf(VoteStatisticCache, vid), value, time.Hour*24).Err()
}
