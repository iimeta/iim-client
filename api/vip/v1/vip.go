package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 会员信息接口请求参数
type VipInfoReq struct {
	g.Meta `path:"/info" tags:"vip" method:"get" summary:"会员信息接口"`
}

// 会员信息接口响应参数
type VipInfoRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.VipInfo
}

// 生成密钥接口请求参数
type GenerateSecretKeyReq struct {
	g.Meta `path:"/generate_secret_key" tags:"vip" method:"get" summary:"生成密钥接口"`
}

// 生成密钥接口响应参数
type GenerateSecretKeyRes struct {
	g.Meta    `mime:"application/json" example:"json"`
	SecretKey string `json:"secret_key"`
}

// 会员权益接口请求参数
type VipsReq struct {
	g.Meta `path:"/vips" tags:"vip" method:"get" summary:"会员权益接口"`
}

// 会员权益接口响应参数
type VipsRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.VipsRes
}
