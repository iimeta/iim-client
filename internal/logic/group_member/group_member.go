package group_member

import (
	"context"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
)

type sGroupMember struct{}

func init() {
	service.RegisterGroupMember(New())
}

func New() service.IGroupMember {
	return &sGroupMember{}
}

// 交接群主权限
func (s *sGroupMember) Handover(ctx context.Context, groupId int, userId int, memberId int) error {

	if err := dao.GroupMember.Handover(ctx, groupId, userId, memberId); err != nil {
		logger.Error(ctx)
		return err
	}

	return nil
}

func (s *sGroupMember) SetLeaderStatus(ctx context.Context, groupId int, userId int, leader int) error {

	if err := dao.GroupMember.SetLeaderStatus(ctx, groupId, userId, leader); err != nil {
		logger.Error(ctx)
		return err
	}

	return nil
}

func (s *sGroupMember) SetMuteStatus(ctx context.Context, groupId int, userId int, status int) error {

	if err := dao.GroupMember.SetMuteStatus(ctx, groupId, userId, status); err != nil {
		logger.Error(ctx)
		return err
	}

	return nil
}
