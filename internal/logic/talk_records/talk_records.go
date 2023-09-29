package talk_records

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/filesystem"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"net/http"
	"slices"
	"time"
)

type sTalkRecords struct {
	Filesystem *filesystem.Filesystem
}

func init() {
	service.RegisterTalkRecords(New())
}

func New() service.ITalkRecords {
	return &sTalkRecords{
		Filesystem: filesystem.NewFilesystem(config.Cfg),
	}
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
func (s *sTalkRecords) GetForwardRecords(ctx context.Context, params model.RecordsForwardReq) ([]*model.TalkRecordsItem, error) {

	uid := service.Session().GetUid(ctx)

	talkRecordsList, userList, err := dao.TalkRecords.GetForwardRecords(ctx, uid, params.RecordId)
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

// 获取会话记录
func (s *sTalkRecords) GetRecords(ctx context.Context, params model.TalkRecordsReq) (*model.TalkRecordsRes, error) {

	if params.TalkType == consts.ChatGroupMode {

		err := service.Group().GroupAuth(ctx, &model.GroupAuth{
			TalkType:   params.TalkType,
			UserId:     service.Session().GetUid(ctx),
			ReceiverId: params.ReceiverId,
		})

		if err != nil {
			items := make([]*model.TalkRecordsItem, 0)
			items = append(items, &model.TalkRecordsItem{
				Content:    "暂无权限查看群消息",
				CreatedAt:  gtime.Datetime(),
				Id:         1,
				MsgId:      util.NewMsgId(),
				MsgType:    consts.ChatMsgSysText,
				ReceiverId: params.ReceiverId,
				TalkType:   params.TalkType,
				UserId:     0,
			})

			return &model.TalkRecordsRes{
				Limit:    params.Limit,
				RecordId: 0,
				Items:    items,
			}, nil
		}
	}

	records, err := s.GetTalkRecords(ctx, &model.QueryTalkRecordsOpt{
		TalkType:   params.TalkType,
		UserId:     service.Session().GetUid(ctx),
		ReceiverId: params.ReceiverId,
		RecordId:   params.RecordId,
		Limit:      params.Limit,
	})

	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	rid := 0
	if length := len(records); length > 0 {
		rid = records[length-1].Sequence
	}

	return &model.TalkRecordsRes{
		Limit:    params.Limit,
		RecordId: rid,
		Items:    records,
	}, nil
}

// 查询下会话记录
func (s *sTalkRecords) SearchHistoryRecords(ctx context.Context, params model.TalkRecordsReq) (*model.TalkRecordsRes, error) {

	uid := service.Session().GetUid(ctx)

	if params.TalkType == consts.ChatGroupMode {
		err := service.Group().GroupAuth(ctx, &model.GroupAuth{
			TalkType:   params.TalkType,
			UserId:     uid,
			ReceiverId: params.ReceiverId,
		})

		if err != nil {
			return &model.TalkRecordsRes{
				Limit:    params.Limit,
				RecordId: 0,
				Items:    make([]*model.TalkRecordsItem, 0),
			}, nil
		}
	}

	m := []int{
		consts.ChatMsgTypeText,
		consts.ChatMsgTypeCode,
		consts.ChatMsgTypeImage,
		consts.ChatMsgTypeVideo,
		consts.ChatMsgTypeAudio,
		consts.ChatMsgTypeFile,
		consts.ChatMsgTypeLocation,
		consts.ChatMsgTypeForward,
		consts.ChatMsgTypeVote,
	}

	if slices.Contains(m, params.MsgType) {
		m = []int{params.MsgType}
	}

	records, err := s.GetTalkRecords(ctx, &model.QueryTalkRecordsOpt{
		TalkType:   params.TalkType,
		MsgType:    m,
		UserId:     service.Session().GetUid(ctx),
		ReceiverId: params.ReceiverId,
		RecordId:   params.RecordId,
		Limit:      params.Limit,
	})

	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	rid := 0
	if length := len(records); length > 0 {
		rid = records[length-1].Sequence
	}

	return &model.TalkRecordsRes{
		Limit:    params.Limit,
		RecordId: rid,
		Items:    records,
	}, nil
}

// 聊天文件下载
func (s *sTalkRecords) Download(ctx context.Context, recordId int) error {

	record, _, err := dao.TalkRecords.GetTalkRecord(ctx, recordId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	uid := service.Session().GetUid(ctx)
	if uid != record.UserId {
		if record.TalkType == consts.ChatPrivateMode {
			if record.ReceiverId != uid {
				logger.Error(ctx, err)
				return errors.New("无访问权限")
			}
		} else {
			if !dao.GroupMember.IsMember(ctx, record.ReceiverId, uid, false) {
				logger.Error(ctx, err)
				return errors.New("无访问权限")
			}
		}
	}

	var fileInfo model.TalkRecordFile
	if err := gjson.Unmarshal([]byte(record.Extra), &fileInfo); err != nil {
		logger.Error(ctx, err)
		return err
	}

	switch fileInfo.Drive {
	case consts.FileDriveLocal:
		g.RequestFromCtx(ctx).Response.ServeFileDownload(s.Filesystem.Local.Path(fileInfo.Path), fileInfo.Name)
	case consts.FileDriveCos:
		g.RequestFromCtx(ctx).Response.RedirectTo(s.Filesystem.Cos.PrivateUrl(fileInfo.Path, 60*time.Second), http.StatusFound)
	default:
		logger.Error(ctx, err)
		return errors.New("未知文件驱动类型")
	}

	return nil
}
