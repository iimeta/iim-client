package group_notice

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

type sGroupNotice struct{}

func init() {
	service.RegisterGroupNotice(New())
}

func New() service.IGroupNotice {
	return &sGroupNotice{}
}

// 群公告列表
func (s *sGroupNotice) List(ctx context.Context, params model.NoticeListReq) (*model.NoticeListRes, error) {

	// 判断是否是群成员
	if !dao.GroupMember.IsMember(ctx, params.GroupId, service.Session().GetUid(ctx), true) {
		return nil, errors.New("无获取数据权限")
	}

	groupNoticeList, userList, err := dao.GroupNotice.GetListAll(ctx, params.GroupId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userMap := util.ToMap(userList, func(t *entity.User) int {
		return t.UserId
	})

	items := make([]*model.Notice, 0)
	for _, notice := range groupNoticeList {
		items = append(items, &model.Notice{
			Id:           notice.Id,
			CreatorId:    notice.CreatorId,
			Title:        notice.Title,
			Content:      notice.Content,
			IsTop:        notice.IsTop,
			IsConfirm:    notice.IsConfirm,
			ConfirmUsers: notice.ConfirmUsers,
			Nickname:     userMap[notice.CreatorId].Nickname,
			Avatar:       userMap[notice.CreatorId].Avatar,
			CreatedAt:    util.FormatDatetime(notice.CreatedAt),
			UpdatedAt:    util.FormatDatetime(notice.UpdatedAt),
		})
	}

	return &model.NoticeListRes{
		Items: items,
	}, nil
}

// 发布或更新群公告
func (s *sGroupNotice) CreateAndUpdate(ctx context.Context, params model.NoticeEditReq) error {

	uid := service.Session().GetUid(ctx)

	if !dao.GroupMember.IsLeader(ctx, params.GroupId, uid) {
		return errors.New("无权限操作")
	}

	if params.NoticeId == "" {
		if _, err := dao.GroupNotice.Insert(ctx, &do.GroupNotice{
			GroupId:      params.GroupId,
			CreatorId:    uid,
			Title:        params.Title,
			Content:      params.Content,
			IsTop:        params.IsTop,
			IsConfirm:    params.IsConfirm,
			ConfirmUsers: "{}",
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	} else {
		if err := dao.GroupNotice.UpdateOne(ctx, bson.M{"_id": params.NoticeId, "group_id": params.GroupId}, bson.M{
			"title":      params.Title,
			"content":    params.Content,
			"is_top":     params.IsTop,
			"is_confirm": params.IsConfirm,
			"updated_at": gtime.Timestamp(),
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	_ = service.TalkMessage().SendSysMessage(ctx, &model.SysMessage{
		MsgType:  consts.MsgSysGroupNotice,
		TalkType: consts.TalkRecordTalkTypeGroup,
		Sender: &model.Sender{
			Id: uid,
		},
		Receiver: &model.Receiver{
			Id:         params.GroupId,
			ReceiverId: params.GroupId,
		},
		GroupNotice: &model.GroupNotice{
			OwnerId:   uid,
			OwnerName: "", // todo
			Title:     params.Title,
			Content:   params.Content,
		},
	})

	return nil
}

// 删除群公告
func (s *sGroupNotice) Delete(ctx context.Context, params model.NoticeDeleteReq) error {

	if err := dao.GroupNotice.UpdateOne(ctx, bson.M{"_id": params.NoticeId, "group_id": params.GroupId}, bson.M{
		"is_delete":  1,
		"deleted_at": gtime.Timestamp(),
		"updated_at": gtime.Timestamp(),
	}); err != nil {
		logger.Error(ctx)
		return err
	}

	return nil
}
