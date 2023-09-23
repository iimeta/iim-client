package group_apply

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sGroupApply struct {
	GroupApplyStorage *cache.GroupApplyStorage
}

func init() {
	service.RegisterGroupApply(New())
}

func New() service.IGroupApply {
	return &sGroupApply{
		GroupApplyStorage: cache.NewGroupApplyStorage(redis.Client),
	}
}

func (s *sGroupApply) Create(ctx context.Context, params model.GroupApplyCreateReq) error {

	groupApply, err := dao.GroupApply.FindOne(ctx, bson.M{"group_id": params.GroupId, "status": model.GroupApplyStatusWait})
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(ctx, err)
		return err
	}

	uid := service.Session().GetUid(ctx)

	if groupApply == nil {
		if _, err = dao.GroupApply.Insert(ctx, &do.GroupApply{
			GroupId: params.GroupId,
			UserId:  uid,
			Status:  model.GroupApplyStatusWait,
			Remark:  params.Remark,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	} else {

		data := g.Map{
			"remark":     params.Remark,
			"updated_at": gtime.Datetime(),
		}

		if err = dao.GroupApply.UpdateOne(ctx, bson.M{"_id": groupApply.Id}, data); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	groupMember, err := dao.GroupMember.FindOne(ctx, bson.M{"group_id": params.GroupId, "leader": 2})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	s.GroupApplyStorage.Incr(ctx, groupMember.UserId)

	if _, err = redis.Publish(ctx, consts.ImTopicChat, gjson.MustEncodeString(g.Map{
		"event": consts.SubEventGroupApply,
		"data": gjson.MustEncodeString(g.Map{
			"group_id": params.GroupId,
			"user_id":  service.Session().GetUid(ctx),
		}),
	})); err != nil {
		logger.Error(ctx, err)
	}

	return nil
}

func (s *sGroupApply) Agree(ctx context.Context, params model.GroupApplyAgreeReq) error {

	uid := service.Session().GetUid(ctx)

	apply, err := dao.GroupApply.FindById(ctx, params.ApplyId)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(ctx, err)
		return err
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return errors.New("申请信息不存在")
	}

	if !dao.GroupMember.IsLeader(ctx, apply.GroupId, uid) {
		return errors.New("无权限访问")
	}

	if apply.Status != model.GroupApplyStatusWait {
		return errors.New("申请信息已被他(她)人处理")
	}

	if !dao.GroupMember.IsMember(ctx, apply.GroupId, apply.UserId, false) {
		if err = dao.Group.Invite(ctx, &do.GroupInvite{
			UserId:    uid,
			GroupId:   apply.GroupId,
			MemberIds: []int{apply.UserId},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	data := bson.M{
		"status":     model.GroupApplyStatusPass,
		"updated_at": gtime.Datetime(),
	}

	if err = dao.GroupApply.UpdateById(ctx, params.ApplyId, data); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (s *sGroupApply) Decline(ctx context.Context, params model.GroupApplyDeclineReq) error {

	apply, err := dao.GroupApply.FindById(ctx, params.ApplyId)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(ctx, err)
		return err
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return errors.New("申请信息不存在")
	}

	if !dao.GroupMember.IsLeader(ctx, apply.GroupId, service.Session().GetUid(ctx)) {
		return errors.New("无权限访问")
	}

	if apply.Status != model.GroupApplyStatusWait {
		return errors.New("申请信息已被他(她)人处理")
	}

	data := bson.M{
		"status":     model.GroupApplyStatusRefuse,
		"reason":     params.Remark,
		"updated_at": gtime.Datetime(),
	}

	err = dao.GroupApply.UpdateById(ctx, params.ApplyId, data)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (s *sGroupApply) List(ctx context.Context, params model.GroupApplyListReq) (*model.GroupApplyListRes, error) {

	if !dao.GroupMember.IsLeader(ctx, params.GroupId, service.Session().GetUid(ctx)) {
		return nil, errors.New("无权限访问")
	}

	groupApplyList, userList, err := dao.GroupApply.List(ctx, []int{params.GroupId})
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("创建群聊失败, 请稍后再试")
	}

	userMap := util.ToMap(userList, func(t *entity.User) int {
		return t.UserId
	})

	list := make([]*model.GroupApplyList, 0)
	for _, apply := range groupApplyList {
		list = append(list, &model.GroupApplyList{
			Id:        apply.Id,
			GroupId:   apply.GroupId,
			UserId:    apply.UserId,
			Remark:    apply.Remark,
			CreatedAt: apply.CreatedAt,
			Nickname:  userMap[apply.UserId].Nickname,
			Avatar:    userMap[apply.UserId].Avatar,
		})
	}

	items := make([]*model.GroupApplyListResponse_Item, 0)
	for _, item := range list {
		items = append(items, &model.GroupApplyListResponse_Item{
			Id:        item.Id,
			UserId:    item.UserId,
			GroupId:   item.GroupId,
			Remark:    item.Remark,
			Avatar:    item.Avatar,
			Nickname:  item.Nickname,
			CreatedAt: util.FormatDatetime(item.CreatedAt),
		})
	}

	return &model.GroupApplyListRes{Items: items}, nil
}

func (s *sGroupApply) All(ctx context.Context) (*model.GroupApplyAllRes, error) {

	all, err := dao.GroupMember.Find(ctx, bson.M{
		"user_id": service.Session().GetUid(ctx),
		"leader":  2,
		"is_quit": bson.M{
			"$ne": 1,
		},
	})

	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("系统异常, 请稍后再试")
	}

	groupIds := make([]int, 0, len(all))
	for _, m := range all {
		groupIds = append(groupIds, m.GroupId)
	}

	resp := &model.GroupApplyAllRes{Items: make([]*model.GroupApplyAllResponse_Item, 0)}

	if len(groupIds) == 0 {
		return resp, nil
	}

	groupApplyList, userList, err := dao.GroupApply.List(ctx, groupIds)
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("系统异常, 请稍后再试")
	}

	userMap := util.ToMap(userList, func(t *entity.User) int {
		return t.UserId
	})

	list := make([]*model.GroupApplyList, 0)
	for _, apply := range groupApplyList {
		list = append(list, &model.GroupApplyList{
			Id:        apply.Id,
			GroupId:   apply.GroupId,
			UserId:    apply.UserId,
			Remark:    apply.Remark,
			CreatedAt: apply.CreatedAt,
			Nickname:  userMap[apply.UserId].Nickname,
			Avatar:    userMap[apply.UserId].Avatar,
		})
	}

	groups, err := dao.Group.Find(ctx, bson.M{"group_id": bson.M{"$in": groupIds}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	groupMap := util.ToMap(groups, func(t *entity.Group) int {
		return t.GroupId
	})

	for _, item := range list {
		resp.Items = append(resp.Items, &model.GroupApplyAllResponse_Item{
			Id:        item.Id,
			UserId:    item.UserId,
			GroupName: groupMap[item.GroupId].GroupName,
			GroupId:   item.GroupId,
			Remark:    item.Remark,
			Avatar:    item.Avatar,
			Nickname:  item.Nickname,
			CreatedAt: util.FormatDatetime(item.CreatedAt),
		})
	}

	s.GroupApplyStorage.Del(ctx, service.Session().GetUid(ctx))

	return resp, nil
}

func (s *sGroupApply) ApplyUnreadNum(ctx context.Context) (*model.ApplyUnreadNumRes, error) {
	return &model.ApplyUnreadNumRes{
		UnreadNum: s.GroupApplyStorage.Get(ctx, service.Session().GetUid(ctx)),
	}, nil
}

func (s *sGroupApply) Delete(ctx context.Context, applyId string, userId int) error {

	if err := dao.GroupApply.Delete(ctx, applyId, userId); err != nil {
		logger.Error(ctx)
		return err
	}

	return nil
}
