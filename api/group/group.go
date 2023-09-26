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
	ApplyCreate(ctx context.Context, req *v1.ApplyCreateReq) (res *v1.ApplyCreateRes, err error)
	ApplyDelete(ctx context.Context, req *v1.ApplyDeleteReq) (res *v1.ApplyDeleteRes, err error)
	ApplyAgree(ctx context.Context, req *v1.ApplyAgreeReq) (res *v1.ApplyAgreeRes, err error)
	ApplyDecline(ctx context.Context, req *v1.ApplyDeclineReq) (res *v1.ApplyDeclineRes, err error)
	ApplyList(ctx context.Context, req *v1.ApplyListReq) (res *v1.ApplyListRes, err error)
	ApplyAll(ctx context.Context, req *v1.ApplyAllReq) (res *v1.ApplyAllRes, err error)
	ApplyUnread(ctx context.Context, req *v1.ApplyUnreadReq) (res *v1.ApplyUnreadRes, err error)
	NoticeDelete(ctx context.Context, req *v1.NoticeDeleteReq) (res *v1.NoticeDeleteRes, err error)
	NoticeEdit(ctx context.Context, req *v1.NoticeEditReq) (res *v1.NoticeEditRes, err error)
	NoticeList(ctx context.Context, req *v1.NoticeListReq) (res *v1.NoticeListRes, err error)
}


