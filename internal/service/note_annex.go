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
		Create(ctx context.Context, data *model.NoteAnnex) error
		// 上传附件
		Upload(ctx context.Context, params model.AnnexUploadReq) (*model.AnnexUploadRes, error)
		// 删除附件
		Delete(ctx context.Context, params model.AnnexDeleteReq) error
		// 恢复附件
		Recover(ctx context.Context, params model.AnnexRecoverReq) error
		// 附件回收站列表
		RecoverList(ctx context.Context) (*model.AnnexRecoverListRes, error)
		// 永久删除附件
		ForeverDelete(ctx context.Context, params model.AnnexForeverDeleteReq) error
		// 下载笔记附件
		Download(ctx context.Context, params model.AnnexDownloadReq) error
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
