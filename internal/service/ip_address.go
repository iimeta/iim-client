// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IIpAddress interface {
		FindAddress(ctx context.Context, ip string) (string, error)
	}
)

var (
	localIpAddress IIpAddress
)

func IpAddress() IIpAddress {
	if localIpAddress == nil {
		panic("implement not found for interface IIpAddress, forgot register?")
	}
	return localIpAddress
}

func RegisterIpAddress(i IIpAddress) {
	localIpAddress = i
}
