package dao

import (
	"context"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
	"go.mongodb.org/mongo-driver/bson"
)

var SplitUpload = NewSplitUploadDao()

type SplitUploadDao struct {
	*MongoDB[entity.SplitUpload]
}

func NewSplitUploadDao(database ...string) *SplitUploadDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &SplitUploadDao{
		MongoDB: NewMongoDB[entity.SplitUpload](database[0], do.SPLIT_UPLOAD_COLLECTION),
	}
}

func (d *SplitUploadDao) GetSplitList(ctx context.Context, uploadId string) ([]*entity.SplitUpload, error) {

	splitUploadList, err := d.Find(ctx, bson.M{"upload_id": uploadId, "type": 2})
	if err != nil {
		return nil, err
	}

	return splitUploadList, nil
}

func (d *SplitUploadDao) GetFile(ctx context.Context, uid int, uploadId string) (*entity.SplitUpload, error) {

	splitUpload, err := d.FindOne(ctx, bson.M{"user_id": uid, "upload_id": uploadId, "type": 1})
	if err != nil {
		return nil, err
	}

	return splitUpload, nil
}
