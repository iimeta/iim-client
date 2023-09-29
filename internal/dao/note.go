package dao

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/utility/db"
	"github.com/iimeta/iim-client/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"html"
)

var Note = NewNoteDao()

type NoteDao struct {
	*MongoDB[entity.Note]
}

func NewNoteDao(database ...string) *NoteDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &NoteDao{
		MongoDB: NewMongoDB[entity.Note](database[0], do.NOTE_COLLECTION),
	}
}

// 笔记详情
func (d *NoteDao) Detail(ctx context.Context, uid int, noteId string) (*entity.Note, *entity.NoteDetail, error) {

	note, err := d.FindById(ctx, noteId)
	if err != nil {
		return nil, nil, err
	}

	noteDetail, err := NoteDetail.FindOne(ctx, bson.M{"note_id": noteId})
	if err != nil {
		return nil, nil, err
	}

	return note, noteDetail, nil
}

// 新建笔记
func (d *NoteDao) Create(ctx context.Context, create *do.NoteCreate) (string, error) {

	abstract := gstr.SubStr(create.MdContent, 0, 200)

	data := &do.Note{
		UserId:   create.UserId,
		ClassId:  create.ClassId,
		Title:    create.Title,
		Image:    util.ParseHtmlImage(create.Content),
		Abstract: util.Strip(abstract),
		Status:   1,
	}

	id, err := d.Insert(ctx, data)
	if err != nil {
		return "", err
	}

	if _, err := NoteDetail.Insert(ctx, &do.NoteDetail{
		NoteId:    id,
		MdContent: html.EscapeString(create.MdContent),
		Content:   html.EscapeString(create.Content),
	}); err != nil {
		return "", err
	}

	return id, nil
}

// 更新笔记信息
func (d *NoteDao) Update(ctx context.Context, edit *do.NoteEdit) error {

	abstract := gstr.SubStr(edit.MdContent, 0, 200)

	if err := d.UpdateById(ctx, edit.NoteId, &do.Note{
		Title:    edit.Title,
		Image:    util.ParseHtmlImage(edit.Content),
		Abstract: util.Strip(abstract),
	}); err != nil {
		return err
	}

	return NoteDetail.UpdateOne(ctx, bson.M{"note_id": edit.NoteId}, &do.NoteDetail{
		MdContent: html.EscapeString(edit.MdContent),
		Content:   html.EscapeString(edit.Content),
	})
}

// 笔记列表
func (d *NoteDao) List(ctx context.Context, list *do.NoteList) ([]*entity.Note, []*entity.NoteClass, error) {

	filter := bson.M{
		"user_id": list.UserId,
		"status":  1,
	}

	if list.FindType == 2 {
		filter["is_asterisk"] = 1
	} else if list.FindType == 3 {
		if list.Cid != "" {
			filter["class_id"] = list.Cid
		} else {
			filter["class_id"] = nil
		}
	} else if list.FindType == 4 {
		filter["tags_id"] = list.Cid
	}

	if list.FindType == 5 {
		filter["status"] = 2
	}

	if list.Keyword != "" {
		filter["title"] = bson.M{
			"$regex": list.Keyword,
		}
	}

	sortField := "-created_at"
	if list.FindType == 1 {
		sortField = "-updated_at"
	}

	noteList, err := d.FindByPage(ctx, &db.Paging{Page: int64(list.Page), PageSize: 20}, filter, sortField)
	if err != nil {
		return nil, nil, err
	}

	classIds := make([]string, 0)
	for _, note := range noteList {
		classIds = append(classIds, note.ClassId)
	}

	noteClassList, err := NoteClass.FindByIds(ctx, classIds)
	if err != nil {
		return nil, nil, err
	}

	return noteList, noteClassList, nil
}

// 笔记标记星号
func (d *NoteDao) Asterisk(ctx context.Context, uid int, noteId string, mode int) error {

	if mode != 1 {
		mode = 0
	}

	return d.UpdateOne(ctx, bson.M{"_id": noteId, "user_id": uid}, bson.M{
		"is_asterisk": mode,
	})
}

// 更新笔记标签
func (d *NoteDao) Tag(ctx context.Context, uid int, noteId string, tags []int) error {
	return d.UpdateOne(ctx, bson.M{"_id": noteId, "user_id": uid}, bson.M{
		"tags_id": util.ToIds(tags),
	})
}

// 移动笔记分类
func (d *NoteDao) Move(ctx context.Context, uid int, noteId, classId string) error {
	return d.UpdateOne(ctx, bson.M{"_id": noteId, "user_id": uid}, bson.M{
		"class_id": classId,
	})
}

// 修改笔记状态
func (d *NoteDao) UpdateStatus(ctx context.Context, uid int, noteId string, status int) error {

	data := map[string]any{
		"status": status,
	}

	if status == 2 {
		data["deleted_at"] = gtime.Timestamp()
	}

	return d.UpdateOne(ctx, bson.M{"_id": noteId, "user_id": uid}, data)
}

// 永久删除笔记
func (d *NoteDao) ForeverDelete(ctx context.Context, uid int, noteId string) error {

	detail, err := d.FindOne(ctx, bson.M{"_id": noteId, "user_id": uid})
	if err != nil {
		return err
	}

	if detail.Status != 2 {
		return errors.New("笔记不能被删除")
	}

	if _, err = NoteDetail.DeleteOne(ctx, bson.M{"note_id": detail.Id}); err != nil {
		return err
	}

	if err = d.DeleteById(ctx, detail.Id); err != nil {
		return err
	}

	if _, err = NoteAnnex.DeleteOne(ctx, bson.M{"user_id": uid, "note_id": detail.Id}); err != nil {
		return err
	}

	return nil
}
