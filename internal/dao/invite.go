package dao

import (
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
)

var Invite = NewInviteDao()

type InviteDao struct {
	*MongoDB[entity.InviteRecord]
}

func NewInviteDao(database ...string) *InviteDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &InviteDao{
		MongoDB: NewMongoDB[entity.InviteRecord](database[0], do.INVITE_RECORD_COLLECTION),
	}
}
