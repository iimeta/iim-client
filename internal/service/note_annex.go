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
	INoteAnnex interface {
		Create(ctx context.Context, data *model.ArticleAnnex) error
		// Upload 上传附件
		Upload(ctx context.Context, params model.ArticleAnnexUploadReq) (*model.ArticleAnnexUploadRes, error)
		// Delete 删除附件
		Delete(ctx context.Context, params model.ArticleAnnexDeleteReq) error
		// Recover 恢复附件
		Recover(ctx context.Context, params model.ArticleAnnexRecoverReq) error
		// RecoverList 附件回收站列表
		RecoverList(ctx context.Context) (*model.ArticleAnnexRecoverListRes, error)
		// ForeverDelete 永久删除附件
		ForeverDelete(ctx context.Context, params model.ArticleAnnexForeverDeleteReq) error
		// Download 下载笔记附件
		Download(ctx context.Context, params model.ArticleAnnexDownloadReq) error
	}
)

var (
	localNoteAnnex INoteAnnex
)

func NoteAnnex() INoteAnnex {
	if localNoteAnnex == nil {
		panic("implement not found for interface INoteAnnex, forgot register?")
	}
	return localNoteAnnex
}

func RegisterNoteAnnex(i INoteAnnex) {
	localNoteAnnex = i
}
