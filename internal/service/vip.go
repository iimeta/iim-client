// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/iim-client/internal/model"
)

type (
	IVip interface {
		InitDailyUsage(ctx context.Context)
		GenerateUidUsageKey(ctx context.Context, uid int, date string) string
		GenerateSecretKey(ctx context.Context) (string, error)
		VipInfo(ctx context.Context) (*model.VipInfo, error)
		Vips(ctx context.Context) ([]*model.Vip, error)
		InviteFriends(ctx context.Context) (string, []*model.InviteRecord, error)
		InviteCode(ctx context.Context) string
		InviteCodeToUid(ctx context.Context, inviteCode string) int
	}
)

var (
	localVip IVip
)

func Vip() IVip {
	if localVip == nil {
		panic("implement not found for interface IVip, forgot register?")
	}
	return localVip
}

func RegisterVip(i IVip) {
	localVip = i
}
