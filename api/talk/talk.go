// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package talk

import (
	"context"
	
	"github.com/iimeta/iim-client/api/talk/v1"
)

type ITalkV1 interface {
	MessagePublish(ctx context.Context, req *v1.MessagePublishReq) (res *v1.MessagePublishRes, err error)
	MessageFile(ctx context.Context, req *v1.MessageFileReq) (res *v1.MessageFileRes, err error)
	MessageVote(ctx context.Context, req *v1.MessageVoteReq) (res *v1.MessageVoteRes, err error)
	MessageVoteHandle(ctx context.Context, req *v1.MessageVoteHandleReq) (res *v1.MessageVoteHandleRes, err error)
	MessageCollect(ctx context.Context, req *v1.MessageCollectReq) (res *v1.MessageCollectRes, err error)
	MessageRevoke(ctx context.Context, req *v1.MessageRevokeReq) (res *v1.MessageRevokeRes, err error)
	MessageDelete(ctx context.Context, req *v1.MessageDeleteReq) (res *v1.MessageDeleteRes, err error)
	Records(ctx context.Context, req *v1.RecordsReq) (res *v1.RecordsRes, err error)
	RecordsSearchHistory(ctx context.Context, req *v1.RecordsSearchHistoryReq) (res *v1.RecordsSearchHistoryRes, err error)
	RecordsForward(ctx context.Context, req *v1.RecordsForwardReq) (res *v1.RecordsForwardRes, err error)
	RecordsFileDownload(ctx context.Context, req *v1.RecordsFileDownloadReq) (res *v1.RecordsFileDownloadRes, err error)
	SessionCreate(ctx context.Context, req *v1.SessionCreateReq) (res *v1.SessionCreateRes, err error)
	SessionDelete(ctx context.Context, req *v1.SessionDeleteReq) (res *v1.SessionDeleteRes, err error)
	SessionTop(ctx context.Context, req *v1.SessionTopReq) (res *v1.SessionTopRes, err error)
	SessionDisturb(ctx context.Context, req *v1.SessionDisturbReq) (res *v1.SessionDisturbRes, err error)
	SessionList(ctx context.Context, req *v1.SessionListReq) (res *v1.SessionListRes, err error)
	SessionClearUnreadNum(ctx context.Context, req *v1.SessionClearUnreadNumReq) (res *v1.SessionClearUnreadNumRes, err error)
	SessionClearContext(ctx context.Context, req *v1.SessionClearContextReq) (res *v1.SessionClearContextRes, err error)
	SessionOpenContext(ctx context.Context, req *v1.SessionOpenContextReq) (res *v1.SessionOpenContextRes, err error)
}


