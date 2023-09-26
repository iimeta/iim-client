// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package contact

import (
	"context"
	
	"github.com/iimeta/iim-client/api/contact/v1"
)

type IContactV1 interface {
	ContactList(ctx context.Context, req *v1.ContactListReq) (res *v1.ContactListRes, err error)
	ContactDelete(ctx context.Context, req *v1.ContactDeleteReq) (res *v1.ContactDeleteRes, err error)
	ContactEditRemark(ctx context.Context, req *v1.ContactEditRemarkReq) (res *v1.ContactEditRemarkRes, err error)
	ContactDetail(ctx context.Context, req *v1.ContactDetailReq) (res *v1.ContactDetailRes, err error)
	ContactSearch(ctx context.Context, req *v1.ContactSearchReq) (res *v1.ContactSearchRes, err error)
	ContactChangeGroup(ctx context.Context, req *v1.ContactChangeGroupReq) (res *v1.ContactChangeGroupRes, err error)
	ApplyCreate(ctx context.Context, req *v1.ApplyCreateReq) (res *v1.ApplyCreateRes, err error)
	ApplyAccept(ctx context.Context, req *v1.ApplyAcceptReq) (res *v1.ApplyAcceptRes, err error)
	ApplyDecline(ctx context.Context, req *v1.ApplyDeclineReq) (res *v1.ApplyDeclineRes, err error)
	ApplyList(ctx context.Context, req *v1.ApplyListReq) (res *v1.ApplyListRes, err error)
	ApplyUnreadNum(ctx context.Context, req *v1.ApplyUnreadNumReq) (res *v1.ApplyUnreadNumRes, err error)
	GroupCreate(ctx context.Context, req *v1.GroupCreateReq) (res *v1.GroupCreateRes, err error)
	GroupUpdate(ctx context.Context, req *v1.GroupUpdateReq) (res *v1.GroupUpdateRes, err error)
	GroupDelete(ctx context.Context, req *v1.GroupDeleteReq) (res *v1.GroupDeleteRes, err error)
	GroupSort(ctx context.Context, req *v1.GroupSortReq) (res *v1.GroupSortRes, err error)
	GroupList(ctx context.Context, req *v1.GroupListReq) (res *v1.GroupListRes, err error)
	GroupSave(ctx context.Context, req *v1.GroupSaveReq) (res *v1.GroupSaveRes, err error)
}


