// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package vip

import (
	"context"
	
	"github.com/iimeta/iim-client/api/vip/v1"
)

type IVipV1 interface {
	VipInfo(ctx context.Context, req *v1.VipInfoReq) (res *v1.VipInfoRes, err error)
	GenerateSecretKey(ctx context.Context, req *v1.GenerateSecretKeyReq) (res *v1.GenerateSecretKeyRes, err error)
}


