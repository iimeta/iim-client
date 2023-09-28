package dao

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"sort"
)

var TalkRecords = NewTalkRecordsDao()

type TalkRecordsDao struct {
	*MongoDB[entity.TalkRecords]
}

func NewTalkRecordsDao(database ...string) *TalkRecordsDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &TalkRecordsDao{
		MongoDB: NewMongoDB[entity.TalkRecords](database[0], do.TALK_RECORDS_COLLECTION),
	}
}

// 删除消息记录
func (d *TalkRecordsDao) DeleteRecord(ctx context.Context, remove *do.RemoveRecord) error {

	var err error
	ids := util.Unique(util.ParseIds(remove.RecordIds))
	talkRecordsList := make([]*entity.TalkRecords, 0)

	if remove.TalkType == consts.ChatPrivateMode {

		filter := bson.M{
			"record_id": bson.M{"$in": ids},
			"talk_type": consts.ChatPrivateMode,
			"$or": bson.A{
				bson.M{"user_id": remove.UserId, "receiver_id": remove.ReceiverId},
				bson.M{"user_id": remove.ReceiverId, "receiver_id": remove.UserId},
			},
		}

		if talkRecordsList, err = d.Find(ctx, filter); err != nil {
			return err
		}

	} else {

		if !GroupMember.IsMember(ctx, remove.ReceiverId, remove.UserId, false) {
			return consts.ErrPermissionDenied
		}

		if talkRecordsList, err = d.Find(ctx, bson.M{"record_id": bson.M{"$in": ids}, "talk_type": consts.ChatGroupMode}); err != nil {
			return err
		}

	}

	findIds := make([]int, 0)
	for _, records := range talkRecordsList {
		findIds = append(findIds, records.RecordId)
	}

	if len(ids) != len(findIds) {
		return errors.New("删除异常")
	}

	items := make([]interface{}, 0, len(ids))
	for _, val := range ids {
		items = append(items, &do.TalkRecordsDelete{
			RecordId:  val,
			UserId:    remove.UserId,
			CreatedAt: gtime.Timestamp(),
		})
	}

	if _, err := Inserts(ctx, d.Database, items); err != nil {
		return err
	}

	return nil
}

// 获取对话消息
func (d *TalkRecordsDao) GetTalkRecords(ctx context.Context, query *do.TalkRecordsQuery) ([]*entity.TalkRecords, []*entity.User, []*entity.TalkRecordsDelete, error) {

	filter := bson.M{
		"talk_type": query.TalkType,
	}

	if query.RecordId > 0 {
		filter["sequence"] = bson.M{
			"$lt": query.RecordId,
		}
	}

	if query.TalkType == consts.ChatPrivateMode {
		filter["$or"] = bson.A{
			bson.M{"user_id": query.UserId, "receiver_id": query.ReceiverId},
			bson.M{"user_id": query.ReceiverId, "receiver_id": query.UserId},
		}
	} else {
		filter["receiver_id"] = query.ReceiverId
	}

	if query.MsgType != nil && len(query.MsgType) > 0 {
		filter["msg_type"] = bson.M{
			"$in": query.MsgType,
		}
	}

	talkRecordsList, err := d.FindByPage(ctx, &db.Paging{PageSize: int64(query.Limit)}, filter, "-sequence")
	if err != nil {
		return nil, nil, nil, err
	}

	userIds := make([]int, 0)
	recordsIds := make([]int, 0)
	for _, records := range talkRecordsList {
		userIds = append(userIds, records.UserId)
		recordsIds = append(recordsIds, records.RecordId)
	}

	userList, err := User.FindUserListByUserIds(ctx, userIds)
	if err != nil {
		return nil, nil, nil, err
	}

	talkRecordsDeleteList := make([]*entity.TalkRecordsDelete, 0)
	if err := Find(ctx, d.Database, do.TALK_RECORDS_DELETE_COLLECTION, bson.M{"record_id": bson.M{"$in": recordsIds}, "user_id": bson.M{"$in": userIds}}, &talkRecordsDeleteList); err != nil {
		return nil, nil, nil, err
	}

	return talkRecordsList, userList, talkRecordsDeleteList, nil
}

// 对话搜索消息
func (d *TalkRecordsDao) SearchTalkRecords() {

}

func (d *TalkRecordsDao) GetTalkRecord(ctx context.Context, recordId int) (*entity.TalkRecords, *entity.User, error) {

	talkRecords, err := d.FindOne(ctx, bson.M{"record_id": recordId})
	if err != nil {
		return nil, nil, err
	}

	if talkRecords.UserId == 0 {
		return talkRecords, nil, err
	}

	user, err := User.FindUserByUserId(ctx, talkRecords.UserId)
	if err != nil {
		return nil, nil, err
	}

	return talkRecords, user, nil
}

// 获取转发消息记录
func (d *TalkRecordsDao) GetForwardRecords(ctx context.Context, uid, recordId int) ([]*entity.TalkRecords, []*entity.User, error) {

	talkRecords, err := d.FindByRecordId(ctx, recordId)
	if err != nil {
		return nil, nil, err
	}

	extra := new(do.TalkRecordExtraForward)
	if err := gjson.Unmarshal([]byte(talkRecords.Extra), &extra); err != nil {
		return nil, nil, err
	}

	talkRecordsList, err := d.Find(ctx, bson.M{"record_id": bson.M{"$in": extra.MsgIds}}, "-sequence")
	if err != nil {
		return nil, nil, err
	}

	userIds := make([]int, 0)
	for _, records := range talkRecordsList {
		userIds = append(userIds, records.UserId)
	}

	userList, err := User.FindUserListByUserIds(ctx, userIds)
	if err != nil {
		return nil, nil, err
	}

	return talkRecordsList, userList, nil
}

func (d *TalkRecordsDao) HandleTalkRecords(ctx context.Context, items []*model.TalkRecordsItem) ([]*model.TalkRecordsItem, error) {

	votes := make([]int, 0)
	for _, item := range items {
		switch item.MsgType {
		case consts.ChatMsgTypeVote:
			votes = append(votes, item.Id)
		}
	}

	hashVotes := make(map[int]*entity.TalkRecordsVote)
	if len(votes) > 0 {

		talkRecordsVoteList, err := TalkRecordsVote.Find(ctx, bson.M{"record_id": bson.M{"$in": votes}})
		if err != nil {
			return nil, err
		}

		for _, vote := range talkRecordsVoteList {
			hashVotes[vote.RecordId] = vote
		}
	}

	newItems := make([]*model.TalkRecordsItem, 0, len(items))
	for _, item := range items {

		data := &model.TalkRecordsItem{
			Id:         item.Id,
			MsgId:      item.MsgId,
			Sequence:   item.Sequence,
			TalkType:   item.TalkType,
			MsgType:    item.MsgType,
			UserId:     item.UserId,
			ReceiverId: item.ReceiverId,
			Nickname:   item.Nickname,
			Avatar:     item.Avatar,
			IsRevoke:   item.IsRevoke,
			IsMark:     item.IsMark,
			IsRead:     item.IsRead,
			Content:    item.Content,
			CreatedAt:  item.CreatedAt,
			Extra:      make(map[string]any),
		}

		_ = gjson.Unmarshal(gjson.MustEncode(item.Extra), &data.Extra)

		switch item.MsgType {
		case consts.ChatMsgTypeVote:

			if value, ok := hashVotes[item.Id]; ok {

				options := make(map[string]any)
				opts := make([]any, 0)

				if err := gjson.Unmarshal([]byte(value.AnswerOption), &options); err == nil {
					arr := make([]string, 0, len(options))
					for k := range options {
						arr = append(arr, k)
					}

					sort.Strings(arr)

					for _, v := range arr {
						opts = append(opts, map[string]any{
							"key":   v,
							"value": options[v],
						})
					}
				}

				users := make([]int, 0)
				if uids, err := TalkRecordsVote.GetVoteAnswerUser(ctx, value.Id); err == nil {
					users = uids
				}

				var statistics any

				if res, err := TalkRecordsVote.GetVoteStatistics(ctx, value.Id); err != nil {
					statistics = map[string]any{
						"count":   0,
						"options": map[string]int{},
					}
				} else {
					statistics = res
				}

				data.Extra = map[string]any{
					"detail": map[string]any{
						"id":            value.Id,
						"record_id":     value.RecordId,
						"title":         value.Title,
						"answer_mode":   value.AnswerMode,
						"status":        value.Status,
						"answer_option": opts,
						"answer_num":    value.AnswerNum,
						"answered_num":  value.AnsweredNum,
					},
					"statistics": statistics,
					"vote_users": users, // 已投票成员
				}
			}
		}

		newItems = append(newItems, data)
	}

	return newItems, nil
}

func (d *TalkRecordsDao) FindByRecordId(ctx context.Context, recordId int) (*entity.TalkRecords, error) {
	return d.FindOne(ctx, bson.M{"record_id": recordId})
}
