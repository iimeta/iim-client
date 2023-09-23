package talk_records

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
)

type sTalkRecords struct{}

func init() {
	service.RegisterTalkRecords(New())
}

func New() service.ITalkRecords {
	return &sTalkRecords{}
}

// 获取对话消息
func (s *sTalkRecords) GetTalkRecords(ctx context.Context, opt *model.QueryTalkRecordsOpt) ([]*model.TalkRecordsItem, error) {

	talkRecordsList, userList, deleteList, err := dao.TalkRecords.GetTalkRecords(ctx, &do.TalkRecordsQuery{
		TalkType:   opt.TalkType,
		UserId:     opt.UserId,
		ReceiverId: opt.ReceiverId,
		MsgType:    opt.MsgType,
		RecordId:   opt.RecordId,
		Limit:      opt.Limit,
	})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userMap := util.ToMap(userList, func(t *entity.User) int {
		return t.UserId
	})

	deleteMap := util.ToMap(deleteList, func(t *entity.TalkRecordsDelete) int {
		return t.RecordId
	})

	items := make([]*model.TalkRecordsItem, 0)
	for _, records := range talkRecordsList {

		if deleteMap[records.RecordId] == nil {

			talkRecordsItem := &model.TalkRecordsItem{
				Id:         records.RecordId,
				Sequence:   records.Sequence,
				MsgId:      records.MsgId,
				TalkType:   records.TalkType,
				MsgType:    records.MsgType,
				UserId:     records.UserId,
				ReceiverId: records.ReceiverId,
				IsRevoke:   records.IsRevoke,
				IsMark:     records.IsMark,
				IsRead:     records.IsRead,
				Content:    records.Content,
				CreatedAt:  util.FormatDatetime(records.CreatedAt),
				Extra:      make(map[string]any),
			}

			if records.Extra != "" {
				talkRecordsItem.Extra, err = gjson.Decode(records.Extra)
				if err != nil {
					logger.Error(ctx, err)
					return nil, err
				}
			}

			if userMap[records.UserId] != nil {
				talkRecordsItem.Nickname = userMap[records.UserId].Nickname
				talkRecordsItem.Avatar = userMap[records.UserId].Avatar
			}

			items = append(items, talkRecordsItem)
		}
	}

	return s.HandleTalkRecords(ctx, items)
}

// 对话搜索消息
func (s *sTalkRecords) SearchTalkRecords() {
	// todo ???
	dao.TalkRecords.SearchTalkRecords()
}

func (s *sTalkRecords) GetTalkRecord(ctx context.Context, recordId int) (*model.TalkRecordsItem, error) {

	record, user, err := dao.TalkRecords.GetTalkRecord(ctx, recordId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	talkRecordsItem := &model.TalkRecordsItem{
		Id:         record.RecordId,
		Sequence:   record.Sequence,
		MsgId:      record.MsgId,
		TalkType:   record.TalkType,
		MsgType:    record.MsgType,
		UserId:     record.UserId,
		ReceiverId: record.ReceiverId,
		IsRevoke:   record.IsRevoke,
		IsMark:     record.IsMark,
		IsRead:     record.IsRead,
		Content:    record.Content,
		CreatedAt:  util.FormatDatetime(record.CreatedAt),
		Extra:      make(map[string]any),
	}

	if user != nil {
		talkRecordsItem.Nickname = user.Nickname
		talkRecordsItem.Avatar = user.Avatar
	}

	if record.Extra != "" {
		talkRecordsItem.Extra, err = gjson.Decode(record.Extra)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	items, err := s.HandleTalkRecords(ctx, []*model.TalkRecordsItem{talkRecordsItem})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return items[0], nil
}

// 获取转发消息记录
func (s *sTalkRecords) GetForwardRecords(ctx context.Context, uid, recordId int) ([]*model.TalkRecordsItem, error) {

	talkRecordsList, userList, err := dao.TalkRecords.GetForwardRecords(ctx, uid, recordId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userMap := util.ToMap(userList, func(t *entity.User) int {
		return t.UserId
	})

	items := make([]*model.TalkRecordsItem, 0)
	for _, records := range talkRecordsList {

		talkRecordsItem := &model.TalkRecordsItem{
			Id:         records.RecordId,
			Sequence:   records.Sequence,
			MsgId:      records.MsgId,
			TalkType:   records.TalkType,
			MsgType:    records.MsgType,
			UserId:     records.UserId,
			ReceiverId: records.ReceiverId,
			Nickname:   userMap[records.UserId].Nickname,
			Avatar:     userMap[records.UserId].Avatar,
			IsRevoke:   records.IsRevoke,
			IsMark:     records.IsMark,
			IsRead:     records.IsRead,
			Content:    records.Content,
			CreatedAt:  util.FormatDatetime(records.CreatedAt),
			Extra:      make(map[string]any),
		}

		if records.Extra != "" {
			talkRecordsItem.Extra, err = gjson.Decode(records.Extra)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
		}

		items = append(items, talkRecordsItem)
	}

	return items, nil
}

func (s *sTalkRecords) HandleTalkRecords(ctx context.Context, items []*model.TalkRecordsItem) ([]*model.TalkRecordsItem, error) {

	talkRecordsItems, err := dao.TalkRecords.HandleTalkRecords(ctx, items)
	if err != nil {
		logger.Error(ctx)
		return nil, err
	}

	return talkRecordsItems, nil
}
