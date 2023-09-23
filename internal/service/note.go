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
		// List 文章列表
		List(ctx context.Context, params model.ArticleListReq) (*model.ArticleListRes, error)
		// Detail 文章详情
		Detail(ctx context.Context, params model.ArticleDetailReq) (*model.ArticleDetailRes, error)
		// Edit 添加或编辑文章
		Edit(ctx context.Context, params model.ArticleEditReq) (*model.ArticleEditRes, error)
		// Delete 删除文章
		Delete(ctx context.Context, params model.ArticleDeleteReq) error
		// Recover 恢复文章
		Recover(ctx context.Context, params model.ArticleRecoverReq) error
		// Upload 文章图片上传
		Upload(ctx context.Context) (*model.ArticleUploadImageRes, error)
		// Move 文章移动
		Move(ctx context.Context, params model.ArticleMoveReq) error
		// Asterisk 标记文章
		Asterisk(ctx context.Context, params model.ArticleAsteriskReq) error
		// Tag 文章标签
		Tag(ctx context.Context, params model.ArticleTagsReq) error
		// ForeverDelete 永久删除文章
		ForeverDelete(ctx context.Context, params model.ArticleForeverDeleteReq) error
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
