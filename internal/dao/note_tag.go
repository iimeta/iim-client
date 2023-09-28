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

var NoteTag = NewNoteTagDao()

type NoteTagDao struct {
	*MongoDB[entity.NoteTag]
}

func NewNoteTagDao(database ...string) *NoteTagDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &NoteTagDao{
		MongoDB: NewMongoDB[entity.NoteTag](database[0], do.NOTE_TAG_COLLECTION),
	}
}

func (d *NoteTagDao) Create(ctx context.Context, uid int, tag string) (string, error) {

	data := &do.NoteTag{
		UserId:  uid,
		TagName: tag,
		Sort:    1,
	}

	id, err := d.Insert(ctx, data)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (d *NoteTagDao) Update(ctx context.Context, uid int, tagId string, tag string) error {
	return d.UpdateOne(ctx, bson.M{"_id": tagId, "user_id": uid}, bson.M{
		"tag_name": tag,
	})
}

func (d *NoteTagDao) Delete(ctx context.Context, uid int, tagId string) error {

	num, err := Note.CountDocuments(ctx, bson.M{"user_id": uid, "tags_id": tagId})
	if err != nil {
		return err
	}

	if num > 0 {
		return errors.New("标签已被使用不能删除")
	}

	_, err = d.DeleteOne(ctx, bson.M{"_id": tagId, "user_id": uid})
	if err != nil {
		return err
	}

	return nil
}

func (d *NoteTagDao) List(ctx context.Context, uid int) ([]*entity.NoteTag, map[string]int, error) {

	noteTagList, err := d.Find(ctx, bson.M{"user_id": uid})
	if err != nil {
		return nil, nil, err
	}

	tagsIds := make([]string, 0)
	for _, noteTag := range noteTagList {
		tagsIds = append(tagsIds, noteTag.Id)
	}

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"user_id": uid,
				"status":  1,
				"tags_id": bson.M{
					"$in": tagsIds,
				},
			},
		},
		{
			"$group": bson.M{
				"_id":     "$tags_id",
				"count":   bson.M{"$sum": 1},
				"tags_id": bson.M{"$first": "$tags_id"},
			},
		},
	}

	results := make([]map[string]interface{}, 0)
	if err := Note.Aggregate(ctx, pipeline, &results); err != nil {
		return nil, nil, err
	}

	countResults := make(map[string]int)
	for _, result := range results {
		countResults[gconv.String(result["tags_id"])] = gconv.Int(result["count"])
	}

	return noteTagList, countResults, nil
}
