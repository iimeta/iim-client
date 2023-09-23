// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package emoticon

import (
	"context"

	"github.com/iimeta/iim-client/api/emoticon/v1"
)

type IEmoticonV1 interface {
	EmoticonSetSystem(ctx context.Context, req *v1.EmoticonSetSystemReq) (res *v1.EmoticonSetSystemRes, err error)
	EmoticonDelete(ctx context.Context, req *v1.EmoticonDeleteReq) (res *v1.EmoticonDeleteRes, err error)
	EmoticonSysList(ctx context.Context, req *v1.EmoticonSysListReq) (res *v1.EmoticonSysListRes, err error)
	EmoticonList(ctx context.Context, req *v1.EmoticonListReq) (res *v1.EmoticonListRes, err error)
	EmoticonUpload(ctx context.Context, req *v1.EmoticonUploadReq) (res *v1.EmoticonUploadRes, err error)
}
