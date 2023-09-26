package talk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/filesystem"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
)

type sMessage struct {
	Filesystem *filesystem.Filesystem
}

func init() {
	service.RegisterMessage(New())
}

func New() service.IMessage {
	return &sMessage{
		Filesystem: filesystem.NewFilesystem(config.Cfg),
	}
}

// 发送文本消息
func (s *sMessage) Text(ctx context.Context, params model.TextMessageReq) error {

	uid := service.Session().GetUid(ctx)

	// todo auth
	if err := service.Group().IsAuth(ctx, &model.AuthOption{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendText(ctx, uid, &model.TextMessageReq{
		Content: params.Text,
		Receiver: &model.MessageReceiver{
			TalkType:   params.TalkType,
			ReceiverId: params.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 发送代码块消息
func (s *sMessage) Code(ctx context.Context, params model.CodeMessageReq) error {

	uid := service.Session().GetUid(ctx)
	if err := service.Group().IsAuth(ctx, &model.AuthOption{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendCode(ctx, uid, &model.CodeMessageReq{
		Lang: params.Lang,
		Code: params.Code,
		Receiver: &model.MessageReceiver{
			TalkType:   params.TalkType,
			ReceiverId: params.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 发送图片消息
func (s *sMessage) Image(ctx context.Context, params model.ImageMessageReq) error {

	_, file, err := g.RequestFromCtx(ctx).Request.FormFile("image")
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("image 字段必传")
	}

	if !util.Include(util.FileSuffix(file.Filename), []string{"png", "jpg", "jpeg", "gif", "webp"}) {
		return errors.New("上传文件格式不正确,仅支持 png、jpg、jpeg、gif 及 webp")
	}

	// 判断上传文件大小(5M)
	if file.Size > 5<<20 {
		return errors.New("上传文件大小不能超过5M")
	}

	if err = service.Group().IsAuth(ctx, &model.AuthOption{
		TalkType:          params.TalkType,
		UserId:            service.Session().GetUid(ctx),
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	stream, err := filesystem.ReadMultipartStream(file)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	ext := util.FileSuffix(file.Filename)

	meta := util.ReadImageMeta(bytes.NewReader(stream))

	filePath := fmt.Sprintf("public/media/image/talk/%s/%s", util.DateNumber(), util.GenImageName(ext, meta.Width, meta.Height))

	if err = s.Filesystem.Default.Write(stream, filePath); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = service.TalkMessage().SendImage(ctx, service.Session().GetUid(ctx), &model.ImageMessageReq{
		Url:    s.Filesystem.Default.PublicUrl(filePath),
		Width:  meta.Width,
		Height: meta.Height,
		Size:   int(file.Size),
		Receiver: &model.MessageReceiver{
			TalkType:   params.TalkType,
			ReceiverId: params.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 发送文件消息
func (s *sMessage) File(ctx context.Context, params model.FileMessageReq) error {

	uid := service.Session().GetUid(ctx)
	if err := service.Group().IsAuth(ctx, &model.AuthOption{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendFile(ctx, uid, &model.FileMessageReq{
		UploadId: params.UploadId,
		Receiver: &model.MessageReceiver{
			TalkType:   params.TalkType,
			ReceiverId: params.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 发送投票消息
func (s *sMessage) Vote(ctx context.Context, params model.VoteMessageReq) error {

	if len(params.Options) <= 1 {
		return errors.New("options 选项必须大于1")
	}

	if len(params.Options) > 6 {
		return errors.New("options 选项不能超过6个")
	}

	uid := service.Session().GetUid(ctx)
	if err := service.Group().IsAuth(ctx, &model.AuthOption{
		TalkType:          consts.ChatGroupMode,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendVote(ctx, uid, &model.VoteMessageReq{
		Mode:      params.Mode,
		Title:     params.Title,
		Options:   params.Options,
		Anonymous: params.Anonymous,
		Receiver: &model.MessageReceiver{
			TalkType:   consts.ChatGroupMode,
			ReceiverId: params.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 发送表情包消息
func (s *sMessage) Emoticon(ctx context.Context, params model.EmoticonMessageReq) error {

	uid := service.Session().GetUid(ctx)
	if err := service.Group().IsAuth(ctx, &model.AuthOption{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendEmoticon(ctx, uid, &model.EmoticonMessageReq{
		EmoticonId: params.EmoticonId,
		Receiver: &model.MessageReceiver{
			TalkType:   params.TalkType,
			ReceiverId: params.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 发送转发消息
func (s *sMessage) Forward(ctx context.Context, params model.ForwardMessageReq) error {

	if params.ReceiveGroupIds == "" && params.ReceiveUserIds == "" {
		return errors.New("receive_user_ids 和 receive_group_ids 不能都为空")
	}

	uid := service.Session().GetUid(ctx)
	if err := service.Group().IsAuth(ctx, &model.AuthOption{
		TalkType:   params.TalkType,
		UserId:     service.Session().GetUid(ctx),
		ReceiverId: params.ReceiverId,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	data := &model.ForwardMessageReq{
		Mode:       params.ForwardMode,
		MessageIds: make([]int, 0),
		Gids:       make([]int, 0),
		Uids:       make([]int, 0),
		Receiver: &model.MessageReceiver{
			TalkType:   params.TalkType,
			ReceiverId: params.ReceiverId,
		},
	}

	for _, id := range util.ParseIds(params.RecordsIds) {
		data.MessageIds = append(data.MessageIds, id)
	}

	for _, id := range util.ParseIds(params.ReceiveUserIds) {
		data.Uids = append(data.Uids, id)
	}

	for _, id := range util.ParseIds(params.ReceiveGroupIds) {
		data.Gids = append(data.Gids, id)
	}

	if err := service.TalkMessage().SendForward(ctx, uid, data); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 发送用户名片消息
func (s *sMessage) Card(ctx context.Context, params model.CardMessageReq) error {

	uid := service.Session().GetUid(ctx)
	if err := service.Group().IsAuth(ctx, &model.AuthOption{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendBusinessCard(ctx, uid, &model.CardMessageReq{
		UserId: params.ReceiverId,
		Receiver: &model.MessageReceiver{
			TalkType:   params.TalkType,
			ReceiverId: params.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 收藏聊天图片
func (s *sMessage) Collect(ctx context.Context, params model.CollectMessageReq) error {

	if err := service.Talk().Collect(ctx, service.Session().GetUid(ctx), params.RecordId); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 撤销聊天记录
func (s *sMessage) Revoke(ctx context.Context, params model.RevokeMessageReq) error {

	if err := service.TalkMessage().Revoke(ctx, service.Session().GetUid(ctx), params.RecordId); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除聊天记录
func (s *sMessage) Delete(ctx context.Context, params model.DeleteMessageReq) error {

	if err := service.Talk().DeleteRecordList(ctx, &model.RemoveRecordListOpt{
		UserId:     service.Session().GetUid(ctx),
		TalkType:   params.TalkType,
		ReceiverId: params.ReceiverId,
		RecordIds:  params.RecordIds,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 投票处理
func (s *sMessage) HandleVote(ctx context.Context, params model.VoteMessageHandleReq) (*model.VoteStatistics, error) {

	data, err := service.TalkMessage().Vote(ctx, service.Session().GetUid(ctx), params.RecordId, params.Options)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return data, nil
}

// 发送位置消息
func (s *sMessage) Location(ctx context.Context, params model.LocationMessageReq) error {

	uid := service.Session().GetUid(ctx)
	if err := service.Group().IsAuth(ctx, &model.AuthOption{
		TalkType:          params.TalkType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := service.TalkMessage().SendLocation(ctx, uid, &model.LocationMessageReq{
		Longitude:   params.Longitude,
		Latitude:    params.Latitude,
		Description: "", // todo 需完善
		Receiver: &model.MessageReceiver{
			TalkType:   params.TalkType,
			ReceiverId: params.ReceiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
