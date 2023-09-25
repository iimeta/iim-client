package dao

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

var TalkSession = NewTalkSessionDao()

type TalkSessionDao struct {
	*MongoDB[entity.TalkSession]
}

func NewTalkSessionDao(database ...string) *TalkSessionDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &TalkSessionDao{
		MongoDB: NewMongoDB[entity.TalkSession](database[0], do.TALK_SESSION_COLLECTION),
	}
}

func (d *TalkSessionDao) List(ctx context.Context, uid int) ([]*entity.TalkSession, []*entity.User, []*entity.Group, error) {

	talkSessionList, err := d.Find(ctx, bson.M{"user_id": uid, "is_delete": 0}, "-updated_at")
	if err != nil {
		return nil, nil, nil, err
	}

	userReceiverIds := make([]int, 0)
	groupReceiverIds := make([]int, 0)
	for _, talkSession := range talkSessionList {
		if talkSession.TalkType == 1 {
			userReceiverIds = append(userReceiverIds, talkSession.ReceiverId)
		} else if talkSession.TalkType == 2 {
			groupReceiverIds = append(groupReceiverIds, talkSession.ReceiverId)
		}
	}

	userList, err := User.FindUserListByUserIds(ctx, userReceiverIds)
	if err != nil {
		return nil, nil, nil, err
	}

	groupList, err := Group.FindGroupListByGroupIds(ctx, groupReceiverIds)
	if err != nil {
		return nil, nil, nil, err
	}

	return talkSessionList, userList, groupList, nil
}

// 创建会话列表
func (d *TalkSessionDao) Create(ctx context.Context, create *do.TalkSessionCreate) (*entity.TalkSession, error) {

	talkSession, err := d.FindOne(ctx, bson.M{"talk_type": create.TalkType, "user_id": create.UserId, "receiver_id": create.ReceiverId})
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {

		talkSession = &entity.TalkSession{
			TalkType:   create.TalkType,
			UserId:     create.UserId,
			ReceiverId: create.ReceiverId,
			IsRobot:    create.IsRobot,
			IsTalk:     create.IsTalk,
		}

		if _, err = d.Insert(ctx, &do.TalkSession{
			TalkType:   talkSession.TalkType,
			UserId:     talkSession.UserId,
			ReceiverId: talkSession.ReceiverId,
			IsRobot:    talkSession.IsRobot,
			IsTalk:     talkSession.IsTalk,
		}); err != nil {
			return nil, err
		}

	} else {

		talkSession.IsTop = 0
		talkSession.IsDelete = 0
		talkSession.IsDisturb = 0

		talkSession.IsRobot = create.IsRobot
		talkSession.IsTalk = create.IsTalk

		if err = d.UpdateById(ctx, talkSession.Id, &do.TalkSession{
			IsTop:     talkSession.IsTop,
			IsDisturb: talkSession.IsDisturb,
			IsDelete:  talkSession.IsDelete,
			IsRobot:   talkSession.IsRobot,
			IsTalk:    talkSession.IsTalk,
			UpdatedAt: gtime.Timestamp(),
		}); err != nil {
			return nil, err
		}
	}

	return talkSession, nil
}

// 删除会话
func (d *TalkSessionDao) Delete(ctx context.Context, uid int, id string) error {

	if err := d.UpdateOne(ctx, bson.M{"_id": id, "user_id": uid}, bson.M{
		"is_delete": 1,
	}); err != nil {
		return err
	}

	return nil
}

// 会话置顶
func (d *TalkSessionDao) Top(ctx context.Context, top *do.TalkSessionTop) error {

	if err := d.UpdateOne(ctx, bson.M{"_id": top.Id, "user_id": top.UserId}, bson.M{
		"is_top": util.BoolToInt(top.Type == 1),
	}); err != nil {
		return err
	}

	return nil
}

// 会话免打扰
func (d *TalkSessionDao) Disturb(ctx context.Context, disturb *do.TalkSessionDisturb) error {

	if err := d.UpdateOne(ctx, bson.M{"user_id": disturb.UserId, "receiver_id": disturb.ReceiverId, "talk_type": disturb.TalkType}, bson.M{
		"is_disturb": disturb.IsDisturb,
	}); err != nil {
		return err
	}

	return nil
}

// 是否开启会话免打扰
func (d *TalkSessionDao) IsDisturb(ctx context.Context, uid int, receiverId int, talkType int) bool {

	talkSession, err := d.FindOne(ctx, bson.M{"user_id": uid, "receiver_id": receiverId, "talk_type": talkType})
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
		}
		return false
	}

	return talkSession.IsDisturb == 1
}

// 批量添加会话列表
func (d *TalkSessionDao) BatchAddList(ctx context.Context, uid int, values map[string]int) {

	ctime := gtime.Timestamp()

	sessionList, err := d.Find(ctx, bson.M{"user_id": uid})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	sessionMap := util.ToMap(sessionList, func(t *entity.TalkSession) int {
		return t.ReceiverId
	})

	talkSessionList := make([]interface{}, 0)
	for k, v := range values {
		if v == 0 {
			continue
		}

		value := strings.Split(k, "_")
		if len(value) != 2 {
			continue
		}

		if sessionMap[gconv.Int(value[1])] != nil {
			continue
		}

		// 获取机器人信息, 判断是不是机器人 todo
		robotInfo, err := Robot.GetRobotByUserId(ctx, gconv.Int(value[1]))
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
			continue
		}

		talkSession := &do.TalkSession{
			TalkType:   gconv.Int(value[0]),
			UserId:     uid,
			ReceiverId: gconv.Int(value[1]),
			CreatedAt:  ctime,
			UpdatedAt:  ctime,
		}

		if robotInfo != nil {
			talkSession.IsRobot = 1
			talkSession.IsTalk = robotInfo.IsTalk
		}

		talkSessionList = append(talkSessionList, talkSession)
	}

	if len(talkSessionList) == 0 {
		return
	}

	if _, err := d.Inserts(ctx, talkSessionList); err != nil {
		logger.Error(ctx, err)
	}
}

func (d *TalkSessionDao) FindBySessionId(ctx context.Context, uid int, receiverId int, talkType int) (string, error) {

	talkSession, err := d.FindOne(ctx, bson.M{"user_id": uid, "receiver_id": receiverId, "talk_type": talkType})
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	return talkSession.Id, nil
}
