// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package article

import (
	"context"

	"github.com/iimeta/iim-client/api/article/v1"
)

type IArticleV1 interface {
	ArticleEdit(ctx context.Context, req *v1.ArticleEditReq) (res *v1.ArticleEditRes, err error)
	ArticleDetail(ctx context.Context, req *v1.ArticleDetailReq) (res *v1.ArticleDetailRes, err error)
	ArticleList(ctx context.Context, req *v1.ArticleListReq) (res *v1.ArticleListRes, err error)
	ArticleDelete(ctx context.Context, req *v1.ArticleDeleteReq) (res *v1.ArticleDeleteRes, err error)
	ArticleRecover(ctx context.Context, req *v1.ArticleRecoverReq) (res *v1.ArticleRecoverRes, err error)
	ArticleMove(ctx context.Context, req *v1.ArticleMoveReq) (res *v1.ArticleMoveRes, err error)
	ArticleAsterisk(ctx context.Context, req *v1.ArticleAsteriskReq) (res *v1.ArticleAsteriskRes, err error)
	ArticleTags(ctx context.Context, req *v1.ArticleTagsReq) (res *v1.ArticleTagsRes, err error)
	ArticleForeverDelete(ctx context.Context, req *v1.ArticleForeverDeleteReq) (res *v1.ArticleForeverDeleteRes, err error)
	ArticleUploadImage(ctx context.Context, req *v1.ArticleUploadImageReq) (res *v1.ArticleUploadImageRes, err error)
	ArticleAnnexUpload(ctx context.Context, req *v1.ArticleAnnexUploadReq) (res *v1.ArticleAnnexUploadRes, err error)
	ArticleAnnexDelete(ctx context.Context, req *v1.ArticleAnnexDeleteReq) (res *v1.ArticleAnnexDeleteRes, err error)
	ArticleAnnexRecover(ctx context.Context, req *v1.ArticleAnnexRecoverReq) (res *v1.ArticleAnnexRecoverRes, err error)
	ArticleAnnexForeverDelete(ctx context.Context, req *v1.ArticleAnnexForeverDeleteReq) (res *v1.ArticleAnnexForeverDeleteRes, err error)
	ArticleAnnexDownload(ctx context.Context, req *v1.ArticleAnnexDownloadReq) (res *v1.ArticleAnnexDownloadRes, err error)
	ArticleAnnexRecoverList(ctx context.Context, req *v1.ArticleAnnexRecoverListReq) (res *v1.ArticleAnnexRecoverListRes, err error)
	ArticleClassList(ctx context.Context, req *v1.ArticleClassListReq) (res *v1.ArticleClassListRes, err error)
	ArticleClassEdit(ctx context.Context, req *v1.ArticleClassEditReq) (res *v1.ArticleClassEditRes, err error)
	ArticleClassDelete(ctx context.Context, req *v1.ArticleClassDeleteReq) (res *v1.ArticleClassDeleteRes, err error)
	ArticleClassSort(ctx context.Context, req *v1.ArticleClassSortReq) (res *v1.ArticleClassSortRes, err error)
	ArticleTagList(ctx context.Context, req *v1.ArticleTagListReq) (res *v1.ArticleTagListRes, err error)
	ArticleTagEdit(ctx context.Context, req *v1.ArticleTagEditReq) (res *v1.ArticleTagEditRes, err error)
	ArticleTagDelete(ctx context.Context, req *v1.ArticleTagDeleteReq) (res *v1.ArticleTagDeleteRes, err error)
}
