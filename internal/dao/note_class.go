package dao

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
	"go.mongodb.org/mongo-driver/bson"
)

var NoteClass = NewNoteClassDao()

type NoteClassDao struct {
	*MongoDB[entity.NoteClass]
}

func NewNoteClassDao(database ...string) *NoteClassDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &NoteClassDao{
		MongoDB: NewMongoDB[entity.NoteClass](database[0], do.NOTE_CLASS_COLLECTION),
	}
}

func (d *NoteClassDao) MaxSort(ctx context.Context, uid int) (int, error) {

	noteClass, err := d.FindOne(ctx, bson.M{"user_id": uid}, "-sort")
	if err != nil {
		return 0, err
	}

	return noteClass.Sort, nil
}

func (d *NoteClassDao) MinSort(ctx context.Context, uid int) (int, error) {

	noteClass, err := d.FindOne(ctx, bson.M{"user_id": uid}, "sort")
	if err != nil {
		return 0, err
	}

	return noteClass.Sort, nil
}

func (d *NoteClassDao) GroupCount(ctx context.Context, uid int) (map[string]int, error) {

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"user_id": uid,
				"status":  1,
			},
		},
		{
			"$group": bson.M{
				"_id":      "$class_id",
				"count":    bson.M{"$sum": 1},
				"class_id": bson.M{"$first": "$class_id"},
			},
		},
	}

	results := make([]map[string]interface{}, 0)
	if err := Note.Aggregate(ctx, pipeline, &results); err != nil {
		return nil, err
	}

	maps := make(map[string]int)
	for _, result := range results {
		maps[gconv.String(result["class_id"])] = gconv.Int(result["count"])
	}

	return maps, nil
}

// 分类列表
func (d *NoteClassDao) List(ctx context.Context, uid int) ([]*entity.NoteClass, error) {

	noteClassList, err := d.Find(ctx, bson.M{"user_id": uid}, "sort")
	if err != nil {
		return nil, err
	}

	return noteClassList, nil
}

// Create 创建分类
func (d *NoteClassDao) Create(ctx context.Context, uid int, name string) (string, error) {

	data := &do.NoteClass{
		UserId:    uid,
		ClassName: name,
		Sort:      1,
	}

	if err := d.UpdateMany(ctx, bson.M{"user_id": uid}, bson.M{
		"$inc": bson.M{
			"sort": 1,
		},
	}); err != nil {
		return "", err
	}

	id, err := d.Insert(ctx, data)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (d *NoteClassDao) Update(ctx context.Context, uid int, cid, name string) error {
	return d.UpdateOne(ctx, bson.M{"_id": cid, "user_id": uid}, bson.M{
		"class_name": name,
	})
}

func (d *NoteClassDao) Delete(ctx context.Context, uid int, cid string) error {

	count, err := Note.CountDocuments(ctx, bson.M{"user_id": uid, "class_id": cid})
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("分类已被使用不能删除")
	}

	if _, err = d.DeleteOne(ctx, bson.M{"_id": cid, "user_id": uid, "is_default": 0}); err != nil {
		return err
	}

	return nil
}

func (d *NoteClassDao) Sort(ctx context.Context, uid int, cid string, mode int) error {

	item, err := d.FindOne(ctx, bson.M{"_id": cid, "user_id": uid})
	if err != nil {
		return err
	}

	if mode == 1 {

		maxSort, err := d.MaxSort(ctx, uid)
		if err != nil {
			return err
		}

		if maxSort == item.Sort {
			return nil
		}

		if err := d.UpdateMany(ctx, bson.M{"user_id": uid, "sort": item.Sort + 1}, bson.M{
			"$inc": bson.M{
				"sort": -1,
			},
		}); err != nil {
			return err
		}

		if err := d.UpdateOne(ctx, bson.M{"_id": cid, "user_id": uid}, bson.M{
			"$inc": bson.M{
				"sort": 1,
			},
		}); err != nil {
			return err
		}

	} else {

		minSort, err := d.MinSort(ctx, uid)
		if err != nil {
			return err
		}

		if minSort == item.Sort {
			return nil
		}

		if err := d.UpdateMany(ctx, bson.M{"user_id": uid, "sort": item.Sort - 1}, bson.M{
			"$inc": bson.M{
				"sort": 1,
			},
		}); err != nil {
			return err
		}

		if err := d.UpdateOne(ctx, bson.M{"_id": cid, "user_id": uid}, bson.M{
			"$inc": bson.M{
				"sort": -1,
			},
		}); err != nil {
			return err
		}
	}

	return nil
}

// 设置默认分类
func (d *NoteClassDao) SetDefaultClass(ctx context.Context, uid int) error {

	count, err := d.CountDocuments(ctx, bson.M{"user_id": uid, "is_default": 1})
	if err != nil {
		return err
	}

	if count == 0 {
		if _, err := d.Insert(ctx, &do.NoteClass{
			UserId:    uid,
			ClassName: "默认分类",
			Sort:      1,
			IsDefault: 1,
		}); err != nil {
			return err
		}
	}

	return nil
}
