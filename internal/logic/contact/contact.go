package contact

import (
	"context"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sContact struct {
	ClientStorage *cache.ClientStorage
}

func init() {
	service.RegisterContact(New())
}

func New() service.IContact {
	return &sContact{
		ClientStorage: cache.NewClientStorage(redis.Client, config.Cfg, cache.NewSidStorage(redis.Client)),
	}
}

// List 联系人列表
func (s *sContact) List(ctx context.Context) (*model.ContactListRes, error) {

	contactList, userList, err := dao.Contact.List(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userMap := util.ToMap(userList, func(t *entity.User) int {
		return t.UserId
	})

	items := make([]*model.ContactListResponse_Item, 0)
	for _, contact := range contactList {
		items = append(items, &model.ContactListResponse_Item{
			Id:       contact.FriendId,
			Nickname: userMap[contact.FriendId].Nickname,
			Gender:   userMap[contact.FriendId].Gender,
			Motto:    userMap[contact.FriendId].Motto,
			Avatar:   userMap[contact.FriendId].Avatar,
			Remark:   contact.Remark,
			GroupId:  contact.GroupId,
		})
	}

	return &model.ContactListRes{Items: items}, nil
}

// Delete 删除联系人
func (s *sContact) Delete(ctx context.Context, params model.ContactDeleteReq) error {

	uid := service.Session().GetUid(ctx)
	if err := dao.Contact.Delete(ctx, uid, params.FriendId); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 解除好友关系后需添加一条聊天记录
	_ = service.TalkMessage().SendSystemText(ctx, uid, &model.TextMessageReq{
		Content: "你与对方已经解除了好友关系",
		Receiver: &model.MessageReceiver{
			TalkType:   consts.ChatPrivateMode,
			ReceiverId: params.FriendId,
		},
	})

	// 删除聊天会话
	sid, err := dao.TalkSession.FindBySessionId(ctx, uid, params.FriendId, consts.ChatPrivateMode)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := dao.TalkSession.Delete(ctx, service.Session().GetUid(ctx), sid); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// Search 查找联系人
func (s *sContact) Search(ctx context.Context, params model.ContactSearchReq) (*model.ContactSearchRes, error) {

	user, err := dao.User.FindUserByAccount(ctx, params.Mobile)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
			return nil, errors.New("用户不存在")
		}

		logger.Error(ctx, err)
		return nil, err
	}

	return &model.ContactSearchRes{
		Id:       user.UserId,
		Mobile:   user.Mobile,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Gender:   user.Gender,
		Motto:    user.Motto,
	}, nil
}

// Remark 编辑联系人备注
func (s *sContact) Remark(ctx context.Context, params model.ContactEditRemarkReq) error {

	if err := dao.Contact.UpdateRemark(ctx, service.Session().GetUid(ctx), params.FriendId, params.Remark); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// Detail 联系人详情信息
func (s *sContact) Detail(ctx context.Context, params model.ContactDetailReq) (*model.ContactDetailRes, error) {

	uid := service.Session().GetUid(ctx)

	user, err := dao.User.FindUserByUserId(ctx, params.UserId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("用户不存在")
		}

		logger.Error(ctx, err)
		return nil, err
	}

	data := model.ContactDetailRes{
		Id:           user.UserId,
		Mobile:       user.Mobile,
		Email:        user.Email,
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		Gender:       user.Gender,
		Motto:        user.Motto,
		FriendApply:  0,
		FriendStatus: 0, // 朋友关系[0:本人;1:陌生人;2:朋友;]
	}

	if uid != user.UserId {

		data.FriendStatus = 1

		contact, err := dao.Contact.FindOne(ctx, bson.M{"user_id": uid, "friend_id": user.UserId})
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
			return nil, err
		}

		if err == nil && contact.Status == 1 {
			if dao.Contact.IsFriend(ctx, uid, user.UserId, false) {
				data.FriendStatus = 2
				data.GroupId = contact.GroupId
				data.Remark = contact.Remark
			}
		}
	}

	return &data, nil
}

// MoveGroup 移动好友分组
func (s *sContact) MoveGroup(ctx context.Context, params model.ContactChangeGroupReq) error {

	err := dao.Contact.MoveGroup(ctx, service.Session().GetUid(ctx), params.UserId, params.GroupId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
