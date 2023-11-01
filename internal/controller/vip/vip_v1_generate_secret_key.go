package vip

import (
	"context"
	"github.com/iimeta/iim-client/internal/service"

	"github.com/iimeta/iim-client/api/vip/v1"
)

func (c *ControllerV1) GenerateSecretKey(ctx context.Context, req *v1.GenerateSecretKeyReq) (res *v1.GenerateSecretKeyRes, err error) {

	secretKey, err := service.Vip().GenerateSecretKey(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.GenerateSecretKeyRes{
		SecretKey: secretKey,
	}

	return
}
