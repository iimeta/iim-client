package talk

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
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
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/util"
	"github.com/iimeta/iim-sdk/sdk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

type sTalkSession struct {
	RedisLock      *cache.RedisLock
	ClientStorage  *cache.ClientStorage
	MessageStorage *cache.MessageStorage
	UnreadStorage  *cache.UnreadStorage
}

func init() {
	service.RegisterTalkSession(New())
}

func New() service.ITalkSession {
	return &sTalkSession{
		RedisLock:      cache.NewRedisLock(redis.Client),
		ClientStorage:  cache.NewClientStorage(redis.Client, config.Cfg, cache.NewSidStorage(redis.Client)),
		MessageStorage: cache.NewMessageStorage(redis.Client),
		UnreadStorage:  cache.NewUnreadStorage(redis.Client),
	}
}

// 创建会话列表
func (s *sTalkSession) Create(ctx context.Context, params model.SessionCreateReq) (*model.SessionCreateRes, error) {

	uid := service.Session().GetUid(ctx)
	agent := g.RequestFromCtx(ctx).Request.UserAgent()

	if agent != "" {
		agent = gmd5.MustEncryptString(agent)
	}

	key := fmt.Sprintf("talk:list:%d-%d-%d-%s", uid, params.ReceiverId, params.TalkType, agent)
	if !s.RedisLock.Lock(ctx, key, 10) {
		return nil, errors.New("会话创建失败")
	}

	if service.TalkMessage().VerifyPermission(ctx, &model.VerifyInfo{
		TalkType:   params.TalkType,
		UserId:     uid,
		ReceiverId: params.ReceiverId,
	}) != nil {
		return nil, errors.New("暂无权限")

	}

	// 获取机器人信息, 判断是不是机器人 todo
	robotInfo, err := dao.Robot.GetRobotByUserId(ctx, params.ReceiverId)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(ctx, err)
		return nil, err
	}

	create := &do.TalkSessionCreate{
		UserId:     uid,
		TalkType:   params.TalkType,
		ReceiverId: params.ReceiverId,
	}

	if robotInfo != nil {
		create.IsRobot = 1
		create.IsTalk = robotInfo.IsTalk
	}

	result, err := dao.TalkSession.Create(ctx, create)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	item := &model.TalkSession{
		Id:         result.Id,
		TalkType:   result.TalkType,
		ReceiverId: result.ReceiverId,
		IsRobot:    result.IsRobot,
		UpdatedAt:  gtime.Datetime(),
	}

	if robotInfo != nil {
		item.IsTalk = robotInfo.IsTalk
		item.IsOnline = 1
	}

	if item.TalkType == consts.ChatPrivateMode {
		item.UnreadNum = s.UnreadStorage.Get(ctx, 1, params.ReceiverId, uid)
		item.Remark = dao.Contact.GetFriendRemark(ctx, uid, params.ReceiverId)

		if user, err := dao.User.FindUserByUserId(ctx, result.ReceiverId); err == nil {
			item.Name = user.Nickname
			item.Avatar = user.Avatar
		}
	} else if result.TalkType == consts.ChatGroupMode {
		if group, err := dao.Group.FindGroupByGroupId(ctx, params.ReceiverId); err == nil {
			item.Name = group.GroupName
		}
	}

	// 查询缓存消息
	if msg, err := s.MessageStorage.Get(ctx, result.TalkType, uid, result.ReceiverId); err == nil {
		item.MsgText = msg.Content
		item.UpdatedAt = msg.Datetime
	}

	return &model.SessionCreateRes{
		Id:         item.Id,
		TalkType:   item.TalkType,
		ReceiverId: item.ReceiverId,
		IsTop:      item.IsTop,
		IsDisturb:  item.IsDisturb,
		IsOnline:   item.IsOnline,
		IsRobot:    item.IsRobot,
		Name:       item.Name,
		Avatar:     item.Avatar,
		Remark:     item.Remark,
		UnreadNum:  item.UnreadNum,
		MsgText:    item.MsgText,
		UpdatedAt:  item.UpdatedAt,
		IsTalk:     item.IsTalk,
	}, nil
}

// 删除列表
func (s *sTalkSession) Delete(ctx context.Context, params model.SessionDeleteReq) error {

	if err := dao.TalkSession.Delete(ctx, service.Session().GetUid(ctx), params.ListId); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 置顶列表
func (s *sTalkSession) Top(ctx context.Context, params model.SessionTopReq) error {

	if err := dao.TalkSession.Top(ctx, &do.TalkSessionTop{
		UserId: service.Session().GetUid(ctx),
		Id:     params.ListId,
		Type:   params.Type,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 会话免打扰
func (s *sTalkSession) Disturb(ctx context.Context, params model.SessionDisturbReq) error {

	if err := dao.TalkSession.Disturb(ctx, &do.TalkSessionDisturb{
		UserId:     service.Session().GetUid(ctx),
		TalkType:   params.TalkType,
		ReceiverId: params.ReceiverId,
		IsDisturb:  params.IsDisturb,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 会话列表
func (s *sTalkSession) List(ctx context.Context) (*model.SessionListRes, error) {

	uid := service.Session().GetUid(ctx)

	// 获取未读消息数
	unReads := s.UnreadStorage.All(ctx, uid)
	if len(unReads) > 0 {
		dao.TalkSession.BatchAddList(ctx, uid, unReads)
	}

	talkSessionList, userList, groupList, err := dao.TalkSession.List(ctx, uid)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userMap := util.ToMap(userList, func(t *entity.User) int {
		return t.UserId
	})

	groupMap := util.ToMap(groupList, func(t *entity.Group) int {
		return t.GroupId
	})

	data := make([]*model.SearchTalkSession, 0)
	for _, talkSession := range talkSessionList {

		session := &model.SearchTalkSession{
			Id:            talkSession.Id,
			TalkType:      talkSession.TalkType,
			ReceiverId:    talkSession.ReceiverId,
			IsDelete:      talkSession.IsDelete,
			IsTop:         talkSession.IsTop,
			IsRobot:       talkSession.IsRobot,
			IsDisturb:     talkSession.IsDisturb,
			UpdatedAt:     talkSession.UpdatedAt,
			IsTalk:        talkSession.IsTalk,
			IsOpenContext: talkSession.IsOpenContext,
		}

		if session.TalkType == 1 {
			if userMap[talkSession.ReceiverId] != nil {
				session.UserAvatar = userMap[talkSession.ReceiverId].Avatar
				session.Nickname = userMap[talkSession.ReceiverId].Nickname
			}
		} else if session.TalkType == 2 {
			if groupMap[talkSession.ReceiverId] != nil {
				session.GroupName = groupMap[talkSession.ReceiverId].GroupName
				session.GroupAvatar = groupMap[talkSession.ReceiverId].Avatar
			}
		}

		data = append(data, session)
	}

	friends := make([]int, 0)
	for _, item := range data {
		if item.TalkType == 1 {
			friends = append(friends, item.ReceiverId)
		}
	}

	// 获取好友备注
	remarks, _ := dao.Contact.Remarks(ctx, uid, friends)

	items := make([]*model.TalkSession, 0)
	for _, item := range data {

		value := &model.TalkSession{
			Id:            item.Id,
			TalkType:      item.TalkType,
			ReceiverId:    item.ReceiverId,
			IsTop:         item.IsTop,
			IsDisturb:     item.IsDisturb,
			IsRobot:       item.IsRobot,
			Avatar:        item.UserAvatar,
			MsgText:       "...",
			UpdatedAt:     util.FormatDatetime(item.UpdatedAt),
			IsTalk:        item.IsTalk,
			IsOpenContext: item.IsOpenContext,
		}

		if num, ok := unReads[fmt.Sprintf("%d_%d", item.TalkType, item.ReceiverId)]; ok {
			value.UnreadNum = num
		}

		if item.TalkType == 1 {
			value.Name = item.Nickname
			value.Avatar = item.UserAvatar
			value.Remark = remarks[item.ReceiverId]
			value.IsOnline = util.BoolToInt(s.ClientStorage.IsOnline(ctx, consts.ImChannelChat, strconv.Itoa(value.ReceiverId)))
		} else {
			value.Name = item.GroupName
			value.Avatar = item.GroupAvatar
		}

		// 临时解决数据越权问题 todo
		if item.TalkType == 2 {
			if err := service.TalkMessage().VerifyPermission(ctx, &model.VerifyInfo{
				TalkType:   item.TalkType,
				UserId:     uid,
				ReceiverId: item.ReceiverId,
			}); err != nil {
				items = append(items, value)
				continue
			}
		}

		// 查询缓存消息
		if msg, err := s.MessageStorage.Get(ctx, item.TalkType, uid, item.ReceiverId); err == nil {
			value.MsgText = msg.Content
			value.UpdatedAt = msg.Datetime
		}

		items = append(items, value)
	}

	return &model.SessionListRes{Items: items}, nil
}

// 清除消息未读数
func (s *sTalkSession) ClearUnreadMessage(ctx context.Context, params model.SessionClearUnreadNumReq) error {

	s.UnreadStorage.Reset(ctx, params.TalkType, params.ReceiverId, service.Session().GetUid(ctx))

	return nil
}

// 会话免打扰
func (s *sTalkSession) OpenContext(ctx context.Context, params model.SessionOpenContextReq) error {

	if err := dao.TalkSession.OpenContext(ctx, &do.TalkSessionOpenContext{
		UserId:        service.Session().GetUid(ctx),
		TalkType:      params.TalkType,
		ReceiverId:    params.ReceiverId,
		IsOpenContext: params.IsOpenContext,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 获取会话
func (s *sTalkSession) FindBySession(ctx context.Context, uid int, receiverId int, talkType int) (*model.TalkSession, error) {

	talkSession, err := dao.TalkSession.FindOne(ctx, bson.M{"user_id": uid, "receiver_id": receiverId, "talk_type": talkType})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.TalkSession{
		Id:            talkSession.Id,
		TalkType:      talkSession.TalkType,
		ReceiverId:    talkSession.ReceiverId,
		IsTop:         talkSession.IsTop,
		IsDisturb:     talkSession.IsDisturb,
		IsRobot:       talkSession.IsRobot,
		IsTalk:        talkSession.IsTalk,
		IsOpenContext: talkSession.IsOpenContext,
	}, nil
}

// 清空上下文
func (s *sTalkSession) ClearContext(ctx context.Context, params model.SessionClearContextReq) (err error) {

	uid := service.Session().GetUid(ctx)

	session, err := service.TalkSession().FindBySession(ctx, uid, params.ReceiverId, params.TalkType)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	defer func() {
		if err == nil {
			if params.TalkType == 2 {

				ids := dao.GroupMember.GetMemberIds(ctx, params.ReceiverId)
				robots, err := sdk.Robot.GetRobotsByUserIds(ctx, ids...)
				if err != nil {
					logger.Error(ctx, err)
					return
				}

				for _, robot := range robots {
					message := sdk.NewMessage()
					message.Stype = params.TalkType
					message.Sid = params.ReceiverId
					_, err = sdk.Robot.ClearMessageContext(ctx, robot, message)
					if err != nil {
						logger.Error(ctx, err)
						return
					}
				}

			} else {

				robot, err := sdk.Robot.GetRobotByUserId(ctx, params.ReceiverId)
				if err != nil {
					logger.Error(ctx, err)
					return
				}

				message := sdk.NewMessage()
				message.Stype = params.TalkType
				message.Sid = session.Id
				_, err = sdk.Robot.ClearMessageContext(ctx, robot, message)
				if err != nil {
					logger.Error(ctx, err)
					return
				}
			}
		}
	}()

	if params.TalkType == 2 {
		if !dao.GroupMember.IsLeader(ctx, params.ReceiverId, uid) {
			return errors.New("暂无权限操作, 需要群主/管理员权限")
		}
	} else {

		// 判断对方是否是自己
		if params.TalkType == consts.ChatPrivateMode && params.ReceiverId == uid {
			return nil
		}

		if !dao.Contact.IsFriend(ctx, uid, params.ReceiverId, false) {
			return errors.New("暂无权限操作")
		}
	}

	return nil
}
