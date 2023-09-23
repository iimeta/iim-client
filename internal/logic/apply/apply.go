package contact

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"time"
)

type sApply struct{}

func init() {
	service.RegisterApply(New())
}

func New() service.IApply {
	return &sApply{}
}

// 获取好友申请未读数
func (s *sApply) ApplyUnreadNum(ctx context.Context) (int, error) {
	return service.ContactApply().GetApplyUnreadNum(ctx, service.Session().GetUid(ctx)), nil
}

// 创建联系人申请
func (s *sApply) Create(ctx context.Context, params model.ContactApplyCreateReq) error {

	uid := service.Session().GetUid(ctx)
	if dao.Contact.IsFriend(ctx, uid, params.FriendId, false) {
		return nil
	}

	if _, err := service.ContactApply().Create(ctx, &model.ContactApply{
		UserId:   service.Session().GetUid(ctx),
		Remarks:  params.Remark,
		FriendId: params.FriendId,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// Accept 同意联系人添加申请
func (s *sApply) Accept(ctx context.Context, params model.ContactApplyAcceptReq) error {

	uid := service.Session().GetUid(ctx)
	applyInfo, err := service.ContactApply().Accept(ctx, &model.ContactApply{
		Remarks: params.Remark,
		ApplyId: params.ApplyId,
		UserId:  uid,
	})

	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	err = service.TalkMessage().SendSystemText(ctx, applyInfo.UserId, &model.TextMessageReq{
		Content: "你们已成为好友, 可以开始聊天咯",
		Receiver: &model.MessageReceiver{
			TalkType:   consts.ChatPrivateMode,
			ReceiverId: applyInfo.FriendId,
		},
	})

	if err != nil {
		logger.Error(ctx, "Apply Accept Err", err.Error())
	}

	return nil
}

// Decline 拒绝联系人添加申请
func (s *sApply) Decline(ctx context.Context, params model.ContactApplyDeclineReq) error {

	if err := service.ContactApply().Decline(ctx, &model.ContactApply{
		UserId:  service.Session().GetUid(ctx),
		Remarks: params.Remark,
		ApplyId: params.ApplyId,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// List 获取联系人申请列表
func (s *sApply) List(ctx context.Context) (*model.ContactApplyListRes, error) {

	list, err := service.ContactApply().List(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.ContactApplyListResponse_Item, 0, len(list))
	for _, item := range list {
		items = append(items, &model.ContactApplyListResponse_Item{
			Id:        item.Id,
			UserId:    item.UserId,
			FriendId:  item.FriendId,
			Remark:    item.Remark,
			Nickname:  item.Nickname,
			Avatar:    item.Avatar,
			CreatedAt: gtime.NewFromTimeStamp(item.CreatedAt).Format(time.DateTime),
		})
	}

	service.ContactApply().ClearApplyUnreadNum(ctx, service.Session().GetUid(ctx))

	return &model.ContactApplyListRes{Items: items}, nil
}
