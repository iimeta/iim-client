package contact_apply

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sContactApply struct{}

func init() {
	service.RegisterContactApply(New())
}

func New() service.IContactApply {
	return &sContactApply{}
}

// 创建好友申请
func (s *sContactApply) Create(ctx context.Context, apply *model.ContactApply) (string, error) {

	user := service.Session().GetUser(ctx)

	contactApply := do.ContactApply{
		UserId:   apply.UserId,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		FriendId: apply.FriendId,
		Remark:   apply.Remarks,
	}

	id, err := dao.Insert(ctx, dao.Contact.Database, contactApply)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	body := map[string]any{
		"event": consts.SubEventContactApply,
		"data": gjson.MustEncodeString(map[string]any{
			"apply_id": id,
			"type":     1,
		}),
	}

	pipe := redis.Pipeline(ctx)
	pipe.Incr(ctx, fmt.Sprintf("im:contact:apply:%d", apply.FriendId))
	pipe.Publish(ctx, consts.ImTopicChat, gjson.MustEncodeString(body))
	_, _ = redis.Pipelined(ctx, pipe)

	return id, nil
}

// 同意好友申请
func (s *sContactApply) Accept(ctx context.Context, apply *model.ContactApply) (*model.ContactApply, error) {

	applyInfo := new(entity.ContactApply)
	if err := dao.FindOne(ctx, dao.Contact.Database, do.CONTACT_APPLY_COLLECTION, bson.M{"_id": apply.ApplyId, "friend_id": apply.UserId}, &applyInfo); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	addFriendFunc := func(uid, fid int, remark string) error {

		contact, err := dao.Contact.FindOne(ctx, bson.M{"user_id": uid, "friend_id": fid})
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
			return err
		}

		// 数据存在则更新
		if contact != nil && contact.Id != "" {
			return dao.Contact.UpdateById(ctx, contact.Id, &do.Contact{
				Remark: remark,
				Status: 1,
			})
		}

		if _, err := dao.Contact.Insert(ctx, &do.Contact{
			UserId:   uid,
			FriendId: fid,
			Remark:   remark,
			Status:   1,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		return nil
	}

	user, err := dao.User.FindUserByUserId(ctx, applyInfo.FriendId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if err = addFriendFunc(applyInfo.UserId, applyInfo.FriendId, user.Nickname); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if err = addFriendFunc(applyInfo.FriendId, applyInfo.UserId, apply.Remarks); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if _, err = dao.DeleteMany(ctx, dao.Contact.Database, do.CONTACT_APPLY_COLLECTION, bson.M{"user_id": applyInfo.UserId, "friend_id": applyInfo.FriendId}); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.ContactApply{
		UserId:   applyInfo.UserId,
		FriendId: applyInfo.FriendId,
	}, nil
}

// 拒绝好友申请
func (s *sContactApply) Decline(ctx context.Context, apply *model.ContactApply) error {

	if _, err := dao.DeleteOne(ctx, dao.Contact.Database, do.CONTACT_APPLY_COLLECTION, bson.M{"_id": apply.ApplyId, "friend_id": apply.UserId}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	body := map[string]any{
		"event": consts.SubEventContactApply,
		"data": gjson.MustEncodeString(map[string]any{
			"apply_id": apply.ApplyId,
			"type":     2,
		}),
	}

	if _, err := redis.Publish(ctx, consts.ImTopicChat, gjson.MustEncodeString(body)); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 好友申请列表
func (s *sContactApply) List(ctx context.Context, uid int) ([]*model.ApplyItem, error) {

	contactApplyList := make([]*entity.ContactApply, 0)
	if err := dao.Find(ctx, dao.Contact.Database, do.CONTACT_APPLY_COLLECTION, bson.M{"friend_id": uid}, &contactApplyList); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userIds := make([]int, 0)
	for _, contactApply := range contactApplyList {
		userIds = append(userIds, contactApply.UserId)
	}

	userList, err := dao.User.FindUserListByUserIds(ctx, userIds)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userMap := make(map[int]*entity.User)
	for _, user := range userList {
		userMap[user.UserId] = user
	}

	items := make([]*model.ApplyItem, 0)
	for _, contactApply := range contactApplyList {

		item := &model.ApplyItem{
			Id:        contactApply.Id,
			UserId:    contactApply.UserId,
			FriendId:  contactApply.FriendId,
			Remark:    contactApply.Remark,
			Nickname:  contactApply.Nickname,
			Avatar:    contactApply.Avatar,
			CreatedAt: contactApply.CreatedAt,
		}

		items = append(items, item)
	}

	return items, nil
}

// 获取申请未读数
func (s *sContactApply) GetApplyUnreadNum(ctx context.Context, uid int) int {

	num, err := redis.GetInt(ctx, fmt.Sprintf("im:contact:apply:%d", uid))
	if err != nil {
		logger.Error(ctx, err)
		return 0
	}

	return num
}

// 清除申请未读数
func (s *sContactApply) ClearApplyUnreadNum(ctx context.Context, uid int) {
	_, _ = redis.Del(ctx, fmt.Sprintf("im:contact:apply:%d", uid))
}
