// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package emoticon

import (
	"context"
	
	"github.com/iimeta/iim-client/api/emoticon/v1"
)

type IEmoticonV1 interface {
	SetSystem(ctx context.Context, req *v1.SetSystemReq) (res *v1.SetSystemRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	SysList(ctx context.Context, req *v1.SysListReq) (res *v1.SysListRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
}


