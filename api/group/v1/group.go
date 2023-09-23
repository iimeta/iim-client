package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/model"
)

// 群列表接口请求参数
type GroupListReq struct {
	g.Meta `path:"/list" tags:"group" method:"get" summary:"群列表接口"`
}

// 群列表接口响应参数
type GroupListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupListRes
}

// 创建群聊接口请求参数
type GroupCreateReq struct {
	g.Meta `path:"/create" tags:"group" method:"post" summary:"创建群聊接口"`
	model.GroupCreateReq
}

// 创建群聊接口响应参数
type GroupCreateRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupCreateRes
}

// 群聊详情接口请求参数
type GroupDetailReq struct {
	g.Meta `path:"/detail" tags:"group" method:"get" summary:"群聊详情接口"`
	model.GroupDetailReq
}

// 群聊详情接口响应参数
type GroupDetailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupDetailRes
}

// 群成员列表接口请求参数
type GroupMemberListReq struct {
	g.Meta `path:"/member/list" tags:"group" method:"get" summary:"群成员列表接口"`
	model.GroupMemberListReq
}

// 群成员列表接口响应参数
type GroupMemberListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupMemberListRes
}

// 解散群聊接口请求参数
type GroupDismissReq struct {
	g.Meta `path:"/dismiss" tags:"group" method:"post" summary:"解散群聊接口"`
	model.GroupDismissReq
}

// 解散群聊接口响应参数
type GroupDismissRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 邀请加入群聊接口请求参数
type GroupInviteReq struct {
	g.Meta `path:"/invite" tags:"group" method:"post" summary:"邀请加入群聊接口"`
	model.GroupInviteReq
}

// 邀请加入群聊接口响应参数
type GroupInviteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 获取待审批入群申请列表接口请求参数
type GetInviteFriendsReq struct {
	g.Meta `path:"/member/invites" tags:"group" method:"get" summary:"获取待审批入群申请列表接口"`
	model.GetInviteFriendsReq
}

// 获取待审批入群申请列表接口响应参数
type GetInviteFriendsRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 退出群聊接口请求参数
type GroupSecedeReq struct {
	g.Meta `path:"/secede" tags:"group" method:"post" summary:"退出群聊接口"`
	model.GroupSecedeReq
}

// 退出群聊接口响应参数
type GroupSecedeRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 设置群聊接口请求参数
type GroupSettingReq struct {
	g.Meta `path:"/setting" tags:"group" method:"post" summary:"设置群聊接口"`
	model.GroupSettingReq
}

// 设置群聊接口响应参数
type GroupSettingRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 群聊名片更新接口请求参数
type GroupRemarkUpdateReq struct {
	g.Meta `path:"/member/remark" tags:"group" method:"post" summary:"群聊名片更新接口"`
	model.GroupRemarkUpdateReq
}

// 群聊名片更新接口响应参数
type GroupRemarkUpdateRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 移出群成员接口请求参数
type GroupRemoveMemberReq struct {
	g.Meta `path:"/member/remove" tags:"group" method:"post" summary:"移出群成员接口"`
	model.GroupRemoveMemberReq
}

// 移出群成员接口响应参数
type GroupRemoveMemberRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 公开群聊列表接口请求参数
type GroupOvertListReq struct {
	g.Meta `path:"/overt/list" tags:"group" method:"get" summary:"公开群聊列表接口"`
	model.GroupOvertListReq
}

// 公开群聊列表接口响应参数
type GroupOvertListRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.GroupOvertListRes
}

// 群主更换接口请求参数
type GroupHandoverReq struct {
	g.Meta `path:"/handover" tags:"group" method:"post" summary:"群主更换接口"`
	model.GroupHandoverReq
}

// 群主更换接口请求参数
type GroupHandoverRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 分配管理员接口请求参数
type GroupAssignAdminReq struct {
	g.Meta `path:"/assign-admin" tags:"group" method:"post" summary:"分配管理员接口"`
	model.GroupAssignAdminReq
}

// 分配管理员接口响应参数
type GroupAssignAdminRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 群成员禁言接口请求参数
type GroupNoSpeakReq struct {
	g.Meta `path:"/no-speak" tags:"group" method:"post" summary:"群成员禁言接口"`
	model.GroupNoSpeakReq
}

// 群成员禁言接口响应参数
type GroupNoSpeakRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 全员禁言接口请求参数
type GroupMuteReq struct {
	g.Meta `path:"/mute" tags:"group" method:"post" summary:"全员禁言接口"`
	model.GroupMuteReq
}

// 全员禁言接口响应参数
type GroupMuteRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// 群公开修改接口请求参数
type GroupOvertReq struct {
	g.Meta `path:"/overt" tags:"group" method:"post" summary:"群公开修改接口"`
	model.GroupOvertReq
}

// 群公开修改接口响应参数
type GroupOvertRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
