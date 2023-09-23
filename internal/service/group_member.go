// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IGroupMember interface {
		// Handover 交接群主权限
		Handover(ctx context.Context, groupId int, userId int, memberId int) error
		SetLeaderStatus(ctx context.Context, groupId int, userId int, leader int) error
		SetMuteStatus(ctx context.Context, groupId int, userId int, status int) error
	}
)

var (
	localGroupMember IGroupMember
)

func GroupMember() IGroupMember {
	if localGroupMember == nil {
		panic("implement not found for interface IGroupMember, forgot register?")
	}
	return localGroupMember
}

func RegisterGroupMember(i IGroupMember) {
	localGroupMember = i
}
