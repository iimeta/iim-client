package dao

import (
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
)

var Account = NewAccountDao()

type AccountDao struct {
	*MongoDB[entity.Account]
}

func NewAccountDao(database ...string) *AccountDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &AccountDao{
		MongoDB: NewMongoDB[entity.Account](database[0], do.ACCOUNT_COLLECTION),
	}
}
