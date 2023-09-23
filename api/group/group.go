// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package group

import (
	"context"

	"github.com/iimeta/iim-client/api/group/v1"
)

type IGroupV1 interface {
	GroupList(ctx context.Context, req *v1.GroupListReq) (res *v1.GroupListRes, err error)
	GroupCreate(ctx context.Context, req *v1.GroupCreateReq) (res *v1.GroupCreateRes, err error)
	GroupDetail(ctx context.Context, req *v1.GroupDetailReq) (res *v1.GroupDetailRes, err error)
	GroupMemberList(ctx context.Context, req *v1.GroupMemberListReq) (res *v1.GroupMemberListRes, err error)
	GroupDismiss(ctx context.Context, req *v1.GroupDismissReq) (res *v1.GroupDismissRes, err error)
	GroupInvite(ctx context.Context, req *v1.GroupInviteReq) (res *v1.GroupInviteRes, err error)
	GetInviteFriends(ctx context.Context, req *v1.GetInviteFriendsReq) (res *v1.GetInviteFriendsRes, err error)
	GroupSecede(ctx context.Context, req *v1.GroupSecedeReq) (res *v1.GroupSecedeRes, err error)
	GroupSetting(ctx context.Context, req *v1.GroupSettingReq) (res *v1.GroupSettingRes, err error)
	GroupRemarkUpdate(ctx context.Context, req *v1.GroupRemarkUpdateReq) (res *v1.GroupRemarkUpdateRes, err error)
	GroupRemoveMember(ctx context.Context, req *v1.GroupRemoveMemberReq) (res *v1.GroupRemoveMemberRes, err error)
	GroupOvertList(ctx context.Context, req *v1.GroupOvertListReq) (res *v1.GroupOvertListRes, err error)
	GroupHandover(ctx context.Context, req *v1.GroupHandoverReq) (res *v1.GroupHandoverRes, err error)
	GroupAssignAdmin(ctx context.Context, req *v1.GroupAssignAdminReq) (res *v1.GroupAssignAdminRes, err error)
	GroupNoSpeak(ctx context.Context, req *v1.GroupNoSpeakReq) (res *v1.GroupNoSpeakRes, err error)
	GroupMute(ctx context.Context, req *v1.GroupMuteReq) (res *v1.GroupMuteRes, err error)
	GroupOvert(ctx context.Context, req *v1.GroupOvertReq) (res *v1.GroupOvertRes, err error)
	GroupApplyCreate(ctx context.Context, req *v1.GroupApplyCreateReq) (res *v1.GroupApplyCreateRes, err error)
	GroupApplyDelete(ctx context.Context, req *v1.GroupApplyDeleteReq) (res *v1.GroupApplyDeleteRes, err error)
	GroupApplyAgree(ctx context.Context, req *v1.GroupApplyAgreeReq) (res *v1.GroupApplyAgreeRes, err error)
	GroupApplyDecline(ctx context.Context, req *v1.GroupApplyDeclineReq) (res *v1.GroupApplyDeclineRes, err error)
	GroupApplyList(ctx context.Context, req *v1.GroupApplyListReq) (res *v1.GroupApplyListRes, err error)
	GroupApplyAll(ctx context.Context, req *v1.GroupApplyAllReq) (res *v1.GroupApplyAllRes, err error)
	GroupApplyUnread(ctx context.Context, req *v1.GroupApplyUnreadReq) (res *v1.GroupApplyUnreadRes, err error)
	GroupNoticeDelete(ctx context.Context, req *v1.GroupNoticeDeleteReq) (res *v1.GroupNoticeDeleteRes, err error)
	GroupNoticeEdit(ctx context.Context, req *v1.GroupNoticeEditReq) (res *v1.GroupNoticeEditRes, err error)
	GroupNoticeList(ctx context.Context, req *v1.GroupNoticeListReq) (res *v1.GroupNoticeListRes, err error)
}
