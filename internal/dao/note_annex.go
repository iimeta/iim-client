package dao

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
	"github.com/iimeta/iim-client/utility/filesystem"
	"go.mongodb.org/mongo-driver/bson"
)

var NoteAnnex = NewNoteAnnexDao()

type NoteAnnexDao struct {
	*MongoDB[entity.NoteAnnex]
	filesystem *filesystem.Filesystem
}

func NewNoteAnnexDao(database ...string) *NoteAnnexDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &NoteAnnexDao{
		MongoDB:    NewMongoDB[entity.NoteAnnex](database[0], do.NOTE_ANNEX_COLLECTION),
		filesystem: filesystem.NewFilesystem(config.Cfg),
	}
}

func (d *NoteAnnexDao) AnnexList(ctx context.Context, uid int, articleId string) ([]*entity.NoteAnnex, error) {
	return d.Find(ctx, bson.M{"user_id": uid, "article_id": articleId, "status": 1})
}

func (d *NoteAnnexDao) RecoverList(ctx context.Context, uid int) ([]*entity.NoteAnnex, []*entity.Note, error) {

	noteAnnexList, err := d.Find(ctx, bson.M{"user_id": uid, "status": 2})
	if err != nil {
		return nil, nil, err
	}

	articleIds := make([]string, 0)
	for _, noteAnnex := range noteAnnexList {
		articleIds = append(articleIds, noteAnnex.ArticleId)
	}

	noteList, err := Note.FindByIds(ctx, articleIds)
	if err != nil {
		return nil, nil, err
	}

	return noteAnnexList, noteList, nil
}

// 更新附件状态
func (d *NoteAnnexDao) UpdateStatus(ctx context.Context, uid int, id string, status int) error {

	data := map[string]any{
		"status": status,
	}

	if status == 2 {
		data["deleted_at"] = gtime.Timestamp()
	}

	return d.UpdateOne(ctx, bson.M{"_id": id, "user_id": uid}, data)
}

// 永久删除笔记附件
func (d *NoteAnnexDao) ForeverDelete(ctx context.Context, uid int, id string) error {

	noteAnnex, err := d.FindOne(ctx, bson.M{"_id": id, "user_id": uid})
	if err != nil {
		return err
	}

	switch noteAnnex.Drive {
	case 1:
		_ = d.filesystem.Local.Delete(noteAnnex.Path)
	case 2:
		_ = d.filesystem.Cos.Delete(noteAnnex.Path)
	}

	return d.DeleteById(ctx, id)
}
