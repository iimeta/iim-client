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
	IEmoticon interface {
		// CollectList 收藏列表
		CollectList(ctx context.Context) (*model.EmoticonListRes, error)
		// DeleteCollect 删除收藏表情包
		DeleteCollect(ctx context.Context, params model.EmoticonDeleteReq) error
		// Upload 上传自定义表情包
		Upload(ctx context.Context) (*model.EmoticonUploadRes, error)
		// SystemList 系统表情包列表
		SystemList(ctx context.Context) ([]*model.EmoticonSysListResponse_Item, error)
		// SetSystemEmoticon 添加或移除系统表情包
		SetSystemEmoticon(ctx context.Context, params model.EmoticonSetSystemReq) (*model.EmoticonSetSystemRes, error)
	}
)

var (
	localEmoticon IEmoticon
)

func Emoticon() IEmoticon {
	if localEmoticon == nil {
		panic("implement not found for interface IEmoticon, forgot register?")
	}
	return localEmoticon
}

func RegisterEmoticon(i IEmoticon) {
	localEmoticon = i
}
