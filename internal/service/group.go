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
		// 创建群聊分组
		Create(ctx context.Context, params model.GroupCreateReq) (*model.GroupCreateRes, error)
		// 解散群聊
		Dismiss(ctx context.Context, params model.GroupDismissReq) error
		// 邀请好友加入群聊
		Invite(ctx context.Context, params model.GroupInviteReq) error
		// 退出群聊
		Secede(ctx context.Context, params model.GroupSecedeReq) error
		// 群设置接口(预留)
		Setting(ctx context.Context, params model.GroupSettingReq) error
		// 移除指定成员(群聊&管理员权限)
		RemoveMembers(ctx context.Context, params model.GroupRemoveMemberReq) error
		// 获取群聊信息
		Detail(ctx context.Context, params model.GroupDetailReq) (*model.GroupDetailRes, error)
		// 修改群备注接口
		UpdateMemberRemark(ctx context.Context, params model.GroupRemarkUpdateReq) error
		GetInviteFriends(ctx context.Context, params model.GetInviteFriendsReq) ([]*model.ContactListItem, error)
		List(ctx context.Context) (*model.GroupListRes, error)
		// 获取群成员列表
		Members(ctx context.Context, params model.GroupMemberListReq) (*model.GroupMemberListRes, error)
		// 公开群列表
		OvertList(ctx context.Context, params model.GroupOvertListReq) (*model.GroupOvertListRes, error)
		// 群主交接
		Handover(ctx context.Context, params model.GroupHandoverReq) error
		// 分配管理员
		AssignAdmin(ctx context.Context, params model.GroupAssignAdminReq) error
		// 禁止发言
		NoSpeak(ctx context.Context, params model.GroupNoSpeakReq) error
		// 全员禁言
		Mute(ctx context.Context, params model.GroupMuteReq) error
		// 公开群
		Overt(ctx context.Context, params model.GroupOvertReq) error
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
