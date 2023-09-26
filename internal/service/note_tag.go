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
	INoteTag interface {
		// List 标签列表
		List(ctx context.Context) (*model.TagListRes, error)
		// Edit 添加或修改标签
		Edit(ctx context.Context, params model.TagEditReq) (*model.TagEditRes, error)
		// Delete 删除标签
		Delete(ctx context.Context, params model.TagDeleteReq) error
	}
)

var (
	localNoteTag INoteTag
)

func NoteTag() INoteTag {
	if localNoteTag == nil {
		panic("implement not found for interface INoteTag, forgot register?")
	}
	return localNoteTag
}

func RegisterNoteTag(i INoteTag) {
	localNoteTag = i
}
