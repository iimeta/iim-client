package dao

import (
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
)

var NoteDetail = NewNoteDetailDao()

type NoteDetailDao struct {
	*MongoDB[entity.NoteDetail]
}

func NewNoteDetailDao(database ...string) *NoteDetailDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &NoteDetailDao{
		MongoDB: NewMongoDB[entity.NoteDetail](database[0], do.NOTE_DETAIL_COLLECTION),
	}
}
