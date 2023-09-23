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
	INoteClass interface {
		// List 分类列表
		List(ctx context.Context) (*model.ArticleClassListRes, error)
		// Edit 添加或修改分类
		Edit(ctx context.Context, params model.ArticleClassEditReq) (*model.ArticleClassEditRes, error)
		// Delete 删除分类
		Delete(ctx context.Context, params model.ArticleClassDeleteReq) error
		// Sort 删除分类
		Sort(ctx context.Context, params model.ArticleClassSortReq) error
	}
)

var (
	localNoteClass INoteClass
)

func NoteClass() INoteClass {
	if localNoteClass == nil {
		panic("implement not found for interface INoteClass, forgot register?")
	}
	return localNoteClass
}

func RegisterNoteClass(i INoteClass) {
	localNoteClass = i
}
