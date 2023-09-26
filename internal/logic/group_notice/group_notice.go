package group_notice

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
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

// 创建群公告
func (s *sGroupNotice) GroupNoticeCreate(ctx context.Context, edit *model.NoticeEdit) error {

	if _, err := dao.GroupNotice.Insert(ctx, &do.GroupNotice{
		GroupId:      edit.GroupId,
		CreatorId:    edit.UserId,
		Title:        edit.Title,
		Content:      edit.Content,
		IsTop:        edit.IsTop,
		IsConfirm:    edit.IsConfirm,
		ConfirmUsers: "{}",
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新群公告
func (s *sGroupNotice) GroupNoticeUpdate(ctx context.Context, edit *model.NoticeEdit) error {

	if err := dao.GroupNotice.UpdateOne(ctx, bson.M{"_id": edit.NoticeId, "group_id": edit.GroupId}, bson.M{
		"title":      edit.Title,
		"content":    edit.Content,
		"is_top":     edit.IsTop,
		"is_confirm": edit.IsConfirm,
		"updated_at": gtime.Timestamp(),
	}); err != nil {
		logger.Error(ctx)
		return err
	}

	return nil
}

func (s *sGroupNotice) GroupNoticeDelete(ctx context.Context, groupId int, noticeId string) error {

	if err := dao.GroupNotice.UpdateOne(ctx, bson.M{"_id": noticeId, "group_id": groupId}, bson.M{
		"is_delete":  1,
		"deleted_at": gtime.Timestamp(),
		"updated_at": gtime.Timestamp(),
	}); err != nil {
		logger.Error(ctx)
		return err
	}

	return nil
}

// 添加或编辑群公告
func (s *sGroupNotice) CreateAndUpdate(ctx context.Context, params model.NoticeEditReq) (string, error) {

	uid := service.Session().GetUid(ctx)

	if !dao.GroupMember.IsLeader(ctx, params.GroupId, uid) {
		return "", errors.New("无权限操作")
	}

	var (
		msg string
		err error
	)

	if params.NoticeId == "" || params.NoticeId == "0" { // todo
		err = s.GroupNoticeCreate(ctx, &model.NoticeEdit{
			UserId:    uid,
			GroupId:   params.GroupId,
			NoticeId:  params.NoticeId,
			Title:     params.Title,
			Content:   params.Content,
			IsTop:     params.IsTop,
			IsConfirm: params.IsConfirm,
		})
		msg = "添加群公告成功"
	} else {
		err = s.GroupNoticeUpdate(ctx, &model.NoticeEdit{
			GroupId:   params.GroupId,
			NoticeId:  params.NoticeId,
			Title:     params.Title,
			Content:   params.Content,
			IsTop:     params.IsTop,
			IsConfirm: params.IsConfirm,
		})
		msg = "更新群公告成功"
	}

	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	_ = service.TalkMessage().SendSysOther(ctx, &model.TalkRecords{
		TalkType:   consts.TalkRecordTalkTypeGroup,
		MsgType:    consts.ChatMsgSysGroupNotice,
		UserId:     uid,
		ReceiverId: params.GroupId,
		Extra: gjson.MustEncodeString(model.TalkRecordExtraGroupNotice{
			OwnerId:   uid,
			OwnerName: "gzydong", // todo
			Title:     params.Title,
			Content:   params.Content,
		}),
	})

	return msg, nil
}

// 删除群公告
func (s *sGroupNotice) Delete(ctx context.Context, params model.NoticeDeleteReq) (string, error) {

	if err := s.GroupNoticeDelete(ctx, params.GroupId, params.NoticeId); err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	return "群公告删除成功", nil
}

// 获取群公告列表(所有)
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

	all := make([]*model.SearchNoticeItem, 0)
	for _, notice := range groupNoticeList {
		all = append(all, &model.SearchNoticeItem{
			Id:           notice.Id,
			CreatorId:    notice.CreatorId,
			Title:        notice.Title,
			Content:      notice.Content,
			IsTop:        notice.IsTop,
			IsConfirm:    notice.IsConfirm,
			ConfirmUsers: notice.ConfirmUsers,
			CreatedAt:    notice.CreatedAt,
			UpdatedAt:    notice.UpdatedAt,
			Avatar:       userMap[notice.CreatorId].Avatar,
			Nickname:     userMap[notice.CreatorId].Nickname,
		})
	}

	items := make([]*model.NoticeListResponse_Item, 0)
	for i := 0; i < len(all); i++ {
		items = append(items, &model.NoticeListResponse_Item{
			Id:           all[i].Id,
			Title:        all[i].Title,
			Content:      all[i].Content,
			IsTop:        all[i].IsTop,
			IsConfirm:    all[i].IsConfirm,
			ConfirmUsers: all[i].ConfirmUsers,
			Avatar:       all[i].Avatar,
			CreatorId:    all[i].CreatorId,
			CreatedAt:    util.FormatDatetime(all[i].CreatedAt),
			UpdatedAt:    util.FormatDatetime(all[i].UpdatedAt),
		})
	}

	return &model.NoticeListRes{
		Items: items,
	}, nil
}
