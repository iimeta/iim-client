package talk

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
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/filesystem"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"net/http"
	"time"
)

type sRecords struct {
	Filesystem *filesystem.Filesystem
}

func init() {
	service.RegisterRecords(New())
}

func New() service.IRecords {
	return &sRecords{
		Filesystem: filesystem.NewFilesystem(config.Cfg),
	}
}

// 获取会话记录
func (s *sRecords) GetRecords(ctx context.Context, params model.GetTalkRecordsReq) (*model.GetTalkRecordsRes, error) {

	if params.TalkType == consts.ChatGroupMode {

		err := service.Group().IsAuth(ctx, &model.AuthOption{
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

			return &model.GetTalkRecordsRes{
				Limit:    params.Limit,
				RecordId: 0,
				Items:    items,
			}, nil
		}
	}

	records, err := service.TalkRecords().GetTalkRecords(ctx, &model.QueryTalkRecordsOpt{
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

	return &model.GetTalkRecordsRes{
		Limit:    params.Limit,
		RecordId: rid,
		Items:    records,
	}, nil
}

// 查询下会话记录
func (s *sRecords) SearchHistoryRecords(ctx context.Context, params model.GetTalkRecordsReq) (*model.GetTalkRecordsRes, error) {

	uid := service.Session().GetUid(ctx)

	if params.TalkType == consts.ChatGroupMode {
		err := service.Group().IsAuth(ctx, &model.AuthOption{
			TalkType:   params.TalkType,
			UserId:     uid,
			ReceiverId: params.ReceiverId,
		})

		if err != nil {
			return &model.GetTalkRecordsRes{
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

	if util.Include(params.MsgType, m) {
		m = []int{params.MsgType}
	}

	records, err := service.TalkRecords().GetTalkRecords(ctx, &model.QueryTalkRecordsOpt{
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

	return &model.GetTalkRecordsRes{
		Limit:    params.Limit,
		RecordId: rid,
		Items:    records,
	}, nil
}

// 获取转发记录
func (s *sRecords) GetForwardRecords(ctx context.Context, params model.GetForwardTalkRecordReq) (*model.GetTalkRecordsRes, error) {

	records, err := service.TalkRecords().GetForwardRecords(ctx, service.Session().GetUid(ctx), params.RecordId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.GetTalkRecordsRes{
		Items: records,
	}, nil
}

// 聊天文件下载
func (s *sRecords) Download(ctx context.Context, recordId int) error {

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

	var fileInfo model.TalkRecordExtraFile
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
