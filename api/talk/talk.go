// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package talk

import (
	"context"

	"github.com/iimeta/iim-client/api/talk/v1"
)

type ITalkV1 interface {
	TalkSessionCreate(ctx context.Context, req *v1.TalkSessionCreateReq) (res *v1.TalkSessionCreateRes, err error)
	TalkSessionDelete(ctx context.Context, req *v1.TalkSessionDeleteReq) (res *v1.TalkSessionDeleteRes, err error)
	TalkSessionTop(ctx context.Context, req *v1.TalkSessionTopReq) (res *v1.TalkSessionTopRes, err error)
	TalkSessionDisturb(ctx context.Context, req *v1.TalkSessionDisturbReq) (res *v1.TalkSessionDisturbRes, err error)
	TalkSessionList(ctx context.Context, req *v1.TalkSessionListReq) (res *v1.TalkSessionListRes, err error)
	TalkSessionClearUnreadNum(ctx context.Context, req *v1.TalkSessionClearUnreadNumReq) (res *v1.TalkSessionClearUnreadNumRes, err error)
	GetRecords(ctx context.Context, req *v1.GetRecordsReq) (res *v1.GetRecordsRes, err error)
	SearchHistoryRecords(ctx context.Context, req *v1.SearchHistoryRecordsReq) (res *v1.SearchHistoryRecordsRes, err error)
	GetForwardRecords(ctx context.Context, req *v1.GetForwardRecordsReq) (res *v1.GetForwardRecordsRes, err error)
	RecordsFileDownload(ctx context.Context, req *v1.RecordsFileDownloadReq) (res *v1.RecordsFileDownloadRes, err error)
}
