package dao

import (
	"context"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
	"go.mongodb.org/mongo-driver/bson"
)

var ContactGroup = NewContactGroupDao()

type ContactGroupDao struct {
	*MongoDB[entity.ContactGroup]
}

func NewContactGroupDao(database ...string) *ContactGroupDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &ContactGroupDao{
		MongoDB: NewMongoDB[entity.ContactGroup](database[0], do.CONTACT_GROUP_COLLECTION),
	}
}

func (d *ContactGroupDao) FindContactGroupList(ctx context.Context, userId int) ([]*entity.ContactGroup, error) {

	contactGroupList, err := d.Find(ctx, bson.M{"user_id": userId}, "sort")
	if err != nil {
		return nil, err
	}

	return contactGroupList, nil
}
