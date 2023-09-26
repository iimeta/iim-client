// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package note

import (
	"context"
	
	"github.com/iimeta/iim-client/api/note/v1"
)

type INoteV1 interface {
	NoteEdit(ctx context.Context, req *v1.NoteEditReq) (res *v1.NoteEditRes, err error)
	NoteDetail(ctx context.Context, req *v1.NoteDetailReq) (res *v1.NoteDetailRes, err error)
	NoteList(ctx context.Context, req *v1.NoteListReq) (res *v1.NoteListRes, err error)
	NoteDelete(ctx context.Context, req *v1.NoteDeleteReq) (res *v1.NoteDeleteRes, err error)
	NoteRecover(ctx context.Context, req *v1.NoteRecoverReq) (res *v1.NoteRecoverRes, err error)
	NoteMove(ctx context.Context, req *v1.NoteMoveReq) (res *v1.NoteMoveRes, err error)
	NoteAsterisk(ctx context.Context, req *v1.NoteAsteriskReq) (res *v1.NoteAsteriskRes, err error)
	NoteTags(ctx context.Context, req *v1.NoteTagsReq) (res *v1.NoteTagsRes, err error)
	NoteForeverDelete(ctx context.Context, req *v1.NoteForeverDeleteReq) (res *v1.NoteForeverDeleteRes, err error)
	NoteUploadImage(ctx context.Context, req *v1.NoteUploadImageReq) (res *v1.NoteUploadImageRes, err error)
	AnnexUpload(ctx context.Context, req *v1.AnnexUploadReq) (res *v1.AnnexUploadRes, err error)
	AnnexDelete(ctx context.Context, req *v1.AnnexDeleteReq) (res *v1.AnnexDeleteRes, err error)
	AnnexRecover(ctx context.Context, req *v1.AnnexRecoverReq) (res *v1.AnnexRecoverRes, err error)
	AnnexForeverDelete(ctx context.Context, req *v1.AnnexForeverDeleteReq) (res *v1.AnnexForeverDeleteRes, err error)
	AnnexDownload(ctx context.Context, req *v1.AnnexDownloadReq) (res *v1.AnnexDownloadRes, err error)
	AnnexRecoverList(ctx context.Context, req *v1.AnnexRecoverListReq) (res *v1.AnnexRecoverListRes, err error)
	ClassList(ctx context.Context, req *v1.ClassListReq) (res *v1.ClassListRes, err error)
	ClassEdit(ctx context.Context, req *v1.ClassEditReq) (res *v1.ClassEditRes, err error)
	ClassDelete(ctx context.Context, req *v1.ClassDeleteReq) (res *v1.ClassDeleteRes, err error)
	ClassSort(ctx context.Context, req *v1.ClassSortReq) (res *v1.ClassSortRes, err error)
	TagList(ctx context.Context, req *v1.TagListReq) (res *v1.TagListRes, err error)
	TagEdit(ctx context.Context, req *v1.TagEditReq) (res *v1.TagEditRes, err error)
	TagDelete(ctx context.Context, req *v1.TagDeleteReq) (res *v1.TagDeleteRes, err error)
}


