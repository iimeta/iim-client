package dao

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/db"
	"github.com/iimeta/iim-client/utility/redis"
	"go.mongodb.org/mongo-driver/bson"
)

var TalkRecordsVote = NewTalkRecordsVoteDao()

type TalkRecordsVoteDao struct {
	*MongoDB[entity.TalkRecordsVote]
	cache *cache.Vote
}

func NewTalkRecordsVoteDao(database ...string) *TalkRecordsVoteDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &TalkRecordsVoteDao{
		MongoDB: NewMongoDB[entity.TalkRecordsVote](database[0], do.TALK_RECORDS_VOTE_COLLECTION),
		cache:   cache.NewVote(redis.Client),
	}
}

func (d *TalkRecordsVoteDao) GetVoteAnswerUser(ctx context.Context, vid string) ([]int, error) {
	// 读取缓存
	if uids, err := d.cache.GetVoteAnswerUser(ctx, vid); err == nil {
		return uids, nil
	}

	uids, err := d.SetVoteAnswerUser(ctx, vid)
	if err != nil {
		return nil, err
	}

	return uids, nil
}

func (d *TalkRecordsVoteDao) SetVoteAnswerUser(ctx context.Context, vid string) ([]int, error) {

	talkRecordsVoteAnswerList := make([]*entity.TalkRecordsVoteAnswer, 0)
	if err := Find(ctx, d.Database, do.TALK_RECORDS_VOTE_ANSWER_COLLECTION, bson.M{"vote_id": vid}, &talkRecordsVoteAnswerList); err != nil {
		return nil, err
	}

	uids := make([]int, 0)
	for _, answer := range talkRecordsVoteAnswerList {
		uids = append(uids, answer.UserId)
	}

	_ = d.cache.SetVoteAnswerUser(ctx, vid, uids)

	return uids, nil
}

type VoteStatistics struct {
	Count   int            `json:"count"`
	Options map[string]int `json:"options"`
}

func (d *TalkRecordsVoteDao) GetVoteStatistics(ctx context.Context, vid string) (*VoteStatistics, error) {

	value, err := d.cache.GetVoteStatistics(ctx, vid)
	if err != nil {
		return d.SetVoteStatistics(ctx, vid)
	}

	statistic := &VoteStatistics{}

	_ = gjson.Unmarshal([]byte(value), statistic)

	return statistic, nil
}

func (d *TalkRecordsVoteDao) SetVoteStatistics(ctx context.Context, vid string) (*VoteStatistics, error) {

	vote, err := d.FindById(ctx, vid)
	if err != nil {
		return nil, err
	}

	answerOption := make(map[string]any)
	if err := gjson.Unmarshal([]byte(vote.AnswerOption), &answerOption); err != nil {
		return nil, err
	}

	talkRecordsVoteAnswerList := make([]*entity.TalkRecordsVoteAnswer, 0)
	if err := Find(ctx, d.Database, do.TALK_RECORDS_VOTE_ANSWER_COLLECTION, bson.M{"vote_id": vid}, &talkRecordsVoteAnswerList); err != nil {
		return nil, err
	}

	options := make([]string, 0)
	for _, answer := range talkRecordsVoteAnswerList {
		options = append(options, answer.Option)
	}

	opts := make(map[string]int)
	for option := range answerOption {
		opts[option] = 0
	}

	for _, option := range options {
		opts[option] += 1
	}

	statistic := &VoteStatistics{
		Options: opts,
		Count:   len(options),
	}

	_ = d.cache.SetVoteStatistics(ctx, vid, gjson.MustEncodeString(statistic))

	return statistic, nil
}
