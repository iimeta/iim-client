package talk

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
)

type sTalk struct{}

func init() {
	service.RegisterTalk(New())
}

func New() service.ITalk {
	return &sTalk{}
}

// 删除消息记录
func (s *sTalk) DeleteRecordList(ctx context.Context, opt *model.RemoveRecordListOpt) error {

	if err := dao.TalkRecords.DeleteRecordList(ctx, &do.RemoveRecord{
		UserId:     opt.UserId,
		TalkType:   opt.TalkType,
		ReceiverId: opt.ReceiverId,
		RecordIds:  opt.RecordIds,
	}); err != nil {
		logger.Error(ctx)
		return err
	}

	return nil
}

// 收藏表情包
func (s *sTalk) Collect(ctx context.Context, uid int, recordId int) error {

	record, err := dao.TalkRecords.FindById(ctx, recordId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if record.MsgType != consts.ChatMsgTypeImage {
		return errors.New("当前消息不支持收藏")
	}

	if record.IsRevoke == 1 {
		return errors.New("当前消息不支持收藏")
	}

	if record.TalkType == consts.ChatPrivateMode {
		if record.UserId != uid && record.ReceiverId != uid {
			return consts.ErrPermissionDenied
		}
	} else if record.TalkType == consts.ChatGroupMode {
		if !dao.GroupMember.IsMember(ctx, record.ReceiverId, uid, true) {
			return consts.ErrPermissionDenied
		}
	}

	var file model.TalkRecordExtraImage
	if err = gjson.Unmarshal([]byte(record.Extra), &file); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.Insert(ctx, dao.Emoticon.Database, &do.EmoticonItem{
		UserId:     uid,
		Url:        file.Url,
		FileSuffix: file.Suffix,
		FileSize:   file.Size,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
