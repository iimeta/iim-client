package dao

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
	"github.com/iimeta/iim-client/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"slices"
)

var Emoticon = NewEmoticonDao()

type EmoticonDao struct {
	*MongoDB[entity.Emoticon]
}

func NewEmoticonDao(database ...string) *EmoticonDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &EmoticonDao{
		MongoDB: NewMongoDB[entity.Emoticon](database[0], do.EMOTICON_COLLECTION),
	}
}

// 获取用户激活的表情包
func (d *EmoticonDao) GetUserInstallIds(ctx context.Context, uid int) []string {

	userEmoticon := new(entity.UserEmoticon)
	if err := FindOne(ctx, d.Database, do.USER_EMOTICON_COLLECTION, bson.M{"user_id": uid}, &userEmoticon); err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
		}
		return []string{}
	}

	return gconv.Strings(gstr.Split(userEmoticon.EmoticonIds, ","))
}

// 获取系统表情包分组列表
func (d *EmoticonDao) GetSystemEmoticonList(ctx context.Context) ([]*entity.Emoticon, error) {

	emoticonList, err := d.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	return emoticonList, nil
}

// 获取系统表情包分组详情列表
func (d *EmoticonDao) GetDetailsAll(ctx context.Context, emoticonId string, uid int) ([]*entity.EmoticonItem, error) {

	filter := bson.M{}
	if emoticonId != "" && emoticonId != "0" {
		filter["emoticon_id"] = emoticonId
	}

	if uid != 0 {
		filter["user_id"] = uid
	}

	emoticonItemList := make([]*entity.EmoticonItem, 0)
	if err := Find(ctx, d.Database, do.EMOTICON_ITEM_COLLECTION, filter, &emoticonItemList, "-created_at"); err != nil {
		return nil, err
	}

	return emoticonItemList, nil
}

func (d *EmoticonDao) RemoveUserSysEmoticon(ctx context.Context, uid int, emoticonId string) error {

	ids := d.GetUserInstallIds(ctx, uid)
	if !slices.Contains(ids, emoticonId) {
		return errors.New("数据不存在")
	}

	items := make([]string, 0, len(ids)-1)
	for _, id := range ids {
		if id != emoticonId {
			items = append(items, id)
		}
	}

	if err := UpdateOne(ctx, d.Database, do.USER_EMOTICON_COLLECTION, bson.M{"user_id": uid}, bson.M{
		"emoticon_ids": gstr.JoinAny(items, ","),
	}); err != nil {
		return err
	}

	return nil
}

func (d *EmoticonDao) AddUserSysEmoticon(ctx context.Context, uid int, emoticonId string) error {

	ids := d.GetUserInstallIds(ctx, uid)
	if slices.Contains(ids, emoticonId) {
		return nil
	}

	ids = append(ids, emoticonId)

	if err := UpdateOne(ctx, d.Database, do.USER_EMOTICON_COLLECTION, bson.M{"user_id": uid}, bson.M{
		"emoticon_ids": gstr.JoinAny(ids, ","),
	}); err != nil {
		return err
	}

	return nil
}

// 删除自定义表情包
func (d *EmoticonDao) DeleteCollect(ctx context.Context, uid int, ids []int) (int64, error) {
	return DeleteOne(ctx, d.Database, do.EMOTICON_ITEM_COLLECTION, bson.M{"_id": bson.M{"$in": ids}, "emoticon_id": "", "user_id": uid})
}
