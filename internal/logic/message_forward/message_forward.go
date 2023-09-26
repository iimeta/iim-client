package message_forward

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/core"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type sMessageForward struct{}

func init() {
	service.RegisterMessageForward(New())
}

func New() service.IMessageForward {
	return &sMessageForward{}
}

// 验证转发消息合法性
func (s *sMessageForward) Verify(ctx context.Context, uid int, params *model.ForwardMessageReq) error {

	filter := bson.M{
		"record_id": bson.M{
			"$in": params.MessageIds,
		},
		"talk_type": params.Receiver.TalkType,
		"msg_type": bson.M{
			"$in": []int{1, 2, 3, 4, 5, 6, 7, 8, consts.ChatMsgTypeForward},
		},
		"is_revoke": bson.M{
			"$ne": 1,
		},
	}

	if params.Receiver.TalkType == consts.ChatPrivateMode {
		filter["$or"] = bson.A{
			bson.M{"user_id": uid, "receiver_id": params.Receiver.ReceiverId},
			bson.M{"user_id": params.Receiver.ReceiverId, "receiver_id": uid},
		}
	}

	count, err := dao.TalkRecords.CountDocuments(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if int(count) != len(params.MessageIds) {
		logger.Error(ctx, err)
		return errors.New("转发消息异常")
	}

	return nil
}

// 批量合并转发
func (s *sMessageForward) MultiMergeForward(ctx context.Context, uid int, params *model.ForwardMessageReq) ([]*model.ForwardRecord, error) {

	receives := make([]map[string]int, 0)

	for _, userId := range params.Uids {
		receives = append(receives, map[string]int{"receiver_id": userId, "talk_type": 1})
	}

	for _, gid := range params.Gids {
		receives = append(receives, map[string]int{"receiver_id": gid, "talk_type": 2})
	}

	tmpRecords, err := aggregation(ctx, params)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	ids := make([]int, 0)
	for _, id := range params.MessageIds {
		ids = append(ids, id)
	}

	extra := gjson.MustEncodeString(model.TalkRecordExtraForward{
		MsgIds:  ids,
		Records: tmpRecords,
	})

	records := make([]interface{}, 0, len(receives))
	recordIds := make([]int, 0)
	for _, item := range receives {
		data := &do.TalkRecords{
			RecordId:   core.IncrRecordId(ctx),
			MsgId:      util.NewMsgId(),
			TalkType:   item["talk_type"],
			MsgType:    consts.ChatMsgTypeForward,
			UserId:     uid,
			ReceiverId: item["receiver_id"],
			Extra:      extra,
		}

		if data.TalkType == consts.ChatGroupMode {
			data.Sequence = dao.Sequence.Get(ctx, 0, data.ReceiverId)
		} else {
			data.Sequence = dao.Sequence.Get(ctx, uid, data.ReceiverId)
		}

		records = append(records, data)
		recordIds = append(recordIds, data.RecordId)
	}

	_, err = dao.TalkRecords.Inserts(ctx, records)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	list := make([]*model.ForwardRecord, 0, len(records))
	for i, record := range records {
		r := record.(*do.TalkRecords)
		list = append(list, &model.ForwardRecord{
			RecordId:   recordIds[i],
			ReceiverId: r.ReceiverId,
			TalkType:   r.TalkType,
		})
	}

	return list, nil
}

// 批量逐条转发
func (s *sMessageForward) MultiSplitForward(ctx context.Context, uid int, params *model.ForwardMessageReq) ([]*model.ForwardRecord, error) {

	receives := make([]map[string]int, 0)

	for _, userId := range params.Uids {
		receives = append(receives, map[string]int{"receiver_id": userId, "talk_type": consts.TalkRecordTalkTypePrivate})
	}

	for _, gid := range params.Gids {
		receives = append(receives, map[string]int{"receiver_id": gid, "talk_type": consts.TalkRecordTalkTypeGroup})
	}

	records, err := dao.TalkRecords.Find(ctx, bson.M{"record_id": bson.M{"$in": params.MessageIds}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]interface{}, 0, len(receives)*len(records))

	recordsLen := int64(len(records))
	recordIds := make([]int, 0)
	for _, v := range receives {
		var sequences []int64

		if v["talk_type"] == consts.TalkRecordTalkTypeGroup {
			sequences = dao.Sequence.BatchGet(ctx, 0, v["receiver_id"], recordsLen)
		} else {
			sequences = dao.Sequence.BatchGet(ctx, uid, v["receiver_id"], recordsLen)
		}

		for i, item := range records {
			records := &do.TalkRecords{
				RecordId:   core.IncrRecordId(ctx),
				MsgId:      util.NewMsgId(),
				TalkType:   v["talk_type"],
				MsgType:    item.MsgType,
				UserId:     uid,
				ReceiverId: v["receiver_id"],
				Content:    item.Content,
				Sequence:   sequences[i],
				Extra:      item.Extra,
			}
			items = append(items, records)
			recordIds = append(recordIds, records.RecordId)
		}
	}

	_, err = dao.TalkRecords.Inserts(ctx, items)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	list := make([]*model.ForwardRecord, 0, len(items))
	for i, item := range items {
		r := item.(*do.TalkRecords)
		list = append(list, &model.ForwardRecord{
			RecordId:   recordIds[i],
			ReceiverId: r.ReceiverId,
			TalkType:   r.TalkType,
		})
	}

	return list, nil
}

type forwardItem struct {
	MsgType  int    `json:"msg_type"`
	Content  string `json:"content"`
	Nickname string `json:"nickname"`
}

// 聚合转发数据
func aggregation(ctx context.Context, params *model.ForwardMessageReq) ([]map[string]any, error) {

	ids := params.MessageIds
	if len(ids) > 3 {
		ids = ids[:3]
	}

	talkRecordsList, err := dao.TalkRecords.Find(ctx, bson.M{"record_id": bson.M{"$in": ids}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userIds := make([]int, 0)
	for _, talkRecords := range talkRecordsList {
		userIds = append(userIds, talkRecords.UserId)
	}

	userList, err := dao.User.FindUserListByUserIds(ctx, userIds)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userMap := util.ToMap(userList, func(t *entity.User) int {
		return t.UserId
	})

	rows := make([]*forwardItem, 0, 3)
	for _, talkRecords := range talkRecordsList {
		rows = append(rows, &forwardItem{
			MsgType:  talkRecords.MsgType,
			Content:  talkRecords.Content,
			Nickname: userMap[talkRecords.UserId].Nickname,
		})
	}

	data := make([]map[string]any, 0)
	for _, row := range rows {
		item := map[string]any{
			"nickname": row.Nickname,
		}

		switch row.MsgType {
		case consts.ChatMsgTypeText:
			item["text"] = util.MtSubstr(strings.TrimSpace(row.Content), 0, 30)
		case consts.ChatMsgTypeCode:
			item["text"] = "【代码消息】"
		case consts.ChatMsgTypeImage:
			item["text"] = "【图片消息】"
		case consts.ChatMsgTypeAudio:
			item["text"] = "【语音消息】"
		case consts.ChatMsgTypeVideo:
			item["text"] = "【视频消息】"
		case consts.ChatMsgTypeFile:
			item["text"] = "【文件消息】"
		case consts.ChatMsgTypeLocation:
			item["text"] = "【位置消息】"
		}

		data = append(data, item)
	}

	return data, nil
}
