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
	ContactApplyCreate(ctx context.Context, req *v1.ContactApplyCreateReq) (res *v1.ContactApplyCreateRes, err error)
	ContactApplyAccept(ctx context.Context, req *v1.ContactApplyAcceptReq) (res *v1.ContactApplyAcceptRes, err error)
	ContactApplyDecline(ctx context.Context, req *v1.ContactApplyDeclineReq) (res *v1.ContactApplyDeclineRes, err error)
	ContactApplyList(ctx context.Context, req *v1.ContactApplyListReq) (res *v1.ContactApplyListRes, err error)
	ContactApplyUnreadNum(ctx context.Context, req *v1.ContactApplyUnreadNumReq) (res *v1.ContactApplyUnreadNumRes, err error)
	ContactGroupCreate(ctx context.Context, req *v1.ContactGroupCreateReq) (res *v1.ContactGroupCreateRes, err error)
	ContactGroupUpdate(ctx context.Context, req *v1.ContactGroupUpdateReq) (res *v1.ContactGroupUpdateRes, err error)
	ContactGroupDelete(ctx context.Context, req *v1.ContactGroupDeleteReq) (res *v1.ContactGroupDeleteRes, err error)
	ContactGroupSort(ctx context.Context, req *v1.ContactGroupSortReq) (res *v1.ContactGroupSortRes, err error)
	ContactGroupList(ctx context.Context, req *v1.ContactGroupListReq) (res *v1.ContactGroupListRes, err error)
	ContactGroupSave(ctx context.Context, req *v1.ContactGroupSaveReq) (res *v1.ContactGroupSaveRes, err error)
}
