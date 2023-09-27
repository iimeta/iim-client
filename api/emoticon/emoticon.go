// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package emoticon

import (
	"context"
	
	"github.com/iimeta/iim-client/api/emoticon/v1"
)

type IEmoticonV1 interface {
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
}


