package note_annex

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/filesystem"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"math"
	"net/http"
	"time"
)

type sNoteAnnex struct {
	Filesystem *filesystem.Filesystem
}

func init() {
	service.RegisterNoteAnnex(New())
}

func New() service.INoteAnnex {
	return &sNoteAnnex{
		Filesystem: filesystem.NewFilesystem(config.Cfg),
	}
}

func (s *sNoteAnnex) Create(ctx context.Context, data *model.NoteAnnex) error {

	if _, err := dao.NoteAnnex.Insert(ctx, &do.NoteAnnex{
		UserId:       data.UserId,
		NoteId:       data.NoteId,
		Drive:        data.Drive,
		Suffix:       data.Suffix,
		Size:         data.Size,
		Path:         data.Path,
		OriginalName: data.OriginalName,
		Status:       data.Status,
		CreatedAt:    gtime.Timestamp(),
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 上传附件
func (s *sNoteAnnex) Upload(ctx context.Context, params model.AnnexUploadReq) (*model.AnnexUploadRes, error) {

	_, file, err := g.RequestFromCtx(ctx).Request.FormFile("annex")
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("annex 字段必传")
	}

	// 判断上传文件大小(10M)
	if file.Size > 10<<20 {
		return nil, errors.New("附件大小不能超过10M")
	}

	stream, err := filesystem.ReadMultipartStream(file)
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("附件上传失败")
	}

	ext := gfile.ExtName(file.Filename)

	filePath := fmt.Sprintf("private/files/note/%s/%s", util.DateNumber(), util.GenFileName(ext))

	if err := s.Filesystem.Default.Write(stream, filePath); err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("附件上传失败")
	}

	data := &do.NoteAnnex{
		UserId:       service.Session().GetUid(ctx),
		NoteId:       params.NoteId,
		Drive:        consts.FileDriveMode(s.Filesystem.Driver()),
		Suffix:       ext,
		Size:         int(file.Size),
		Path:         filePath,
		OriginalName: file.Filename,
		Status:       1,
	}

	id, err := dao.NoteAnnex.Insert(ctx, data)
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("附件上传失败")
	}

	return &model.AnnexUploadRes{
		Id:           id,
		Size:         data.Size,
		Path:         data.Path,
		Suffix:       data.Suffix,
		OriginalName: data.OriginalName,
	}, nil
}

// 删除附件
func (s *sNoteAnnex) Delete(ctx context.Context, params model.AnnexDeleteReq) error {

	err := dao.NoteAnnex.UpdateStatus(ctx, service.Session().GetUid(ctx), params.AnnexId, 2)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 恢复附件
func (s *sNoteAnnex) Recover(ctx context.Context, params model.AnnexRecoverReq) error {

	err := dao.NoteAnnex.UpdateStatus(ctx, service.Session().GetUid(ctx), params.AnnexId, 1)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 附件回收站列表
func (s *sNoteAnnex) RecoverList(ctx context.Context) (*model.AnnexRecoverListRes, error) {

	noteAnnexList, noteList, err := dao.NoteAnnex.RecoverList(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	noteMap := util.ToMap(noteList, func(t *entity.Note) string {
		return t.Id
	})

	items := make([]*model.NoteAnnex, 0)
	for _, annex := range noteAnnexList {
		at := gtime.NewFromTimeStamp(annex.DeletedAt).Add(time.Hour * 24 * 30)
		items = append(items, &model.NoteAnnex{
			Id:           annex.Id,
			NoteId:       annex.NoteId,
			Title:        noteMap[annex.NoteId].Title,
			OriginalName: annex.OriginalName,
			DeletedAt:    util.FormatDatetime(annex.DeletedAt),
			Day:          int(math.Ceil(float64(at.Second() / 86400))), // todo 有没有更好的方法
		})
	}

	return &model.AnnexRecoverListRes{
		Items: items,
		Paginate: &model.Paginate{
			Page:  1,
			Size:  10000,
			Total: len(items),
		},
	}, nil
}

// 永久删除附件
func (s *sNoteAnnex) ForeverDelete(ctx context.Context, params model.AnnexForeverDeleteReq) error {

	if err := dao.NoteAnnex.ForeverDelete(ctx, service.Session().GetUid(ctx), params.AnnexId); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 下载笔记附件
func (s *sNoteAnnex) Download(ctx context.Context, params model.AnnexDownloadReq) error {

	info, err := dao.NoteAnnex.FindById(ctx, params.AnnexId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if info.UserId != service.Session().GetUid(ctx) {
		return errors.New("无权限下载")
	}

	switch info.Drive {
	case consts.FileDriveLocal:
		g.RequestFromCtx(ctx).Response.ServeFileDownload(s.Filesystem.Local.Path(info.Path), info.OriginalName)
	case consts.FileDriveCos:
		g.RequestFromCtx(ctx).Response.RedirectTo(s.Filesystem.Cos.PrivateUrl(info.Path, 60*time.Second), http.StatusFound)

	default:
		logger.Error(ctx, err)
		return errors.New("未知文件驱动类型")
	}

	return nil
}
