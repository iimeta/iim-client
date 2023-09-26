// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/iim-client/internal/model"
)

type (
	INote interface {
		// 笔记列表
		List(ctx context.Context, params model.NoteListReq) (*model.NoteListRes, error)
		// 笔记详情
		Detail(ctx context.Context, params model.NoteDetailReq) (*model.NoteDetailRes, error)
		// 添加或编辑笔记
		Edit(ctx context.Context, params model.NoteEditReq) (*model.NoteEditRes, error)
		// 删除笔记
		Delete(ctx context.Context, params model.NoteDeleteReq) error
		// 恢复笔记
		Recover(ctx context.Context, params model.NoteRecoverReq) error
		// 笔记图片上传
		Upload(ctx context.Context) (*model.NoteUploadImageRes, error)
		// 笔记移动
		Move(ctx context.Context, params model.NoteMoveReq) error
		// 标记笔记
		Asterisk(ctx context.Context, params model.NoteAsteriskReq) error
		// 笔记标签
		Tag(ctx context.Context, params model.NoteTagsReq) error
		// 永久删除笔记
		ForeverDelete(ctx context.Context, params model.NoteForeverDeleteReq) error
	}
)

var (
	localNote INote
)

func Note() INote {
	if localNote == nil {
		panic("implement not found for interface INote, forgot register?")
	}
	return localNote
}

func RegisterNote(i INote) {
	localNote = i
}
