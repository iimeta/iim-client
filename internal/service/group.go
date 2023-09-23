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
	IGroup interface {
		IsAuth(ctx context.Context, opt *model.AuthOption) error
		// Create 创建群聊分组
		Create(ctx context.Context, params model.GroupCreateReq) (*model.GroupCreateRes, error)
		// Dismiss 解散群聊
		Dismiss(ctx context.Context, params model.GroupDismissReq) error
		// Invite 邀请好友加入群聊
		Invite(ctx context.Context, params model.GroupInviteReq) error
		// SignOut 退出群聊
		SignOut(ctx context.Context, params model.GroupSecedeReq) error
		// Setting 群设置接口(预留)
		Setting(ctx context.Context, params model.GroupSettingReq) error
		// RemoveMembers 移除指定成员(群聊&管理员权限)
		RemoveMembers(ctx context.Context, params model.GroupRemoveMemberReq) error
		// Detail 获取群聊信息
		Detail(ctx context.Context, params model.GroupDetailReq) (*model.GroupDetailRes, error)
		// UpdateMemberRemark 修改群备注接口
		UpdateMemberRemark(ctx context.Context, params model.GroupRemarkUpdateReq) error
		GetInviteFriends(ctx context.Context, params model.GetInviteFriendsReq) ([]*model.ContactListItem, error)
		List(ctx context.Context) (*model.GroupListRes, error)
		// Members 获取群成员列表
		Members(ctx context.Context, params model.GroupMemberListReq) (*model.GroupMemberListRes, error)
		// OvertList 公开群列表
		OvertList(ctx context.Context, params model.GroupOvertListReq) (*model.GroupOvertListRes, error)
		// Handover 群主交接
		Handover(ctx context.Context, params model.GroupHandoverReq) error
		// AssignAdmin 分配管理员
		AssignAdmin(ctx context.Context, params model.GroupAssignAdminReq) error
		// NoSpeak 禁止发言
		NoSpeak(ctx context.Context, params model.GroupNoSpeakReq) error
		// Mute 全员禁言
		Mute(ctx context.Context, params model.GroupMuteReq) error
		// Overt 公开群
		Overt(ctx context.Context, params model.GroupOvertReq) error
		// Secede 退出群聊[仅管理员及群成员]
		Secede(ctx context.Context, groupId int, uid int) error
	}
)

var (
	localGroup IGroup
)

func Group() IGroup {
	if localGroup == nil {
		panic("implement not found for interface IGroup, forgot register?")
	}
	return localGroup
}

func RegisterGroup(i IGroup) {
	localGroup = i
}
