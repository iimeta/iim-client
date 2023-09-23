package note

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/model/do"
	"github.com/iimeta/iim-client/internal/model/entity"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/filesystem"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"html"
)

type sNote struct {
	Filesystem *filesystem.Filesystem
}

func init() {
	service.RegisterNote(New())
}

func New() service.INote {
	return &sNote{
		Filesystem: filesystem.NewFilesystem(config.Cfg),
	}
}

// 文章列表
func (s *sNote) List(ctx context.Context, params model.ArticleListReq) (*model.ArticleListRes, error) {

	noteList, noteClassList, err := dao.Note.List(ctx, &do.NoteList{
		UserId:   service.Session().GetUid(ctx),
		Keyword:  params.Keyword,
		FindType: params.FindType,
		Cid:      params.Cid,
		Page:     params.Page,
	})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	noteClassMap := util.ToMap(noteClassList, func(t *entity.NoteClass) string {
		return t.Id
	})

	items := make([]*model.ArticleListItem, 0)
	for _, note := range noteList {

		articleListItem := &model.ArticleListItem{
			Id:         note.Id,
			UserId:     note.UserId,
			TagsId:     note.TagsId,
			Title:      note.Title,
			Abstract:   note.Abstract,
			Image:      note.Image,
			IsAsterisk: note.IsAsterisk,
			Status:     note.Status,
			CreatedAt:  note.CreatedAt,
			UpdatedAt:  note.UpdatedAt,
		}

		if noteClassMap[note.ClassId] != nil {
			articleListItem.ClassName = noteClassMap[note.ClassId].ClassName
			articleListItem.ClassId = noteClassMap[note.ClassId].Id
		}

		items = append(items, articleListItem)
	}

	list := make([]*model.ArticleListResponse_Item, 0)
	for _, item := range items {
		list = append(list, &model.ArticleListResponse_Item{
			Id:         item.Id,
			ClassId:    item.ClassId,
			TagsId:     item.TagsId,
			Title:      item.Title,
			ClassName:  item.ClassName,
			Image:      item.Image,
			IsAsterisk: item.IsAsterisk,
			Status:     item.Status,
			CreatedAt:  util.FormatDatetime(item.CreatedAt),
			UpdatedAt:  util.FormatDatetime(item.UpdatedAt),
			Abstract:   item.Abstract,
		})
	}

	return &model.ArticleListRes{
		Items: list,
		Paginate: &model.ArticleListResponse_Paginate{
			Page:  1,
			Size:  1000,
			Total: len(list),
		},
	}, nil
}

// 文章详情
func (s *sNote) Detail(ctx context.Context, params model.ArticleDetailReq) (*model.ArticleDetailRes, error) {

	uid := service.Session().GetUid(ctx)

	note, noteDetail, err := dao.Note.Detail(ctx, uid, params.ArticleId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("笔记不存在")
	}

	detail := &model.ArticleDetailInfo{
		Id:         note.Id,
		UserId:     note.UserId,
		ClassId:    note.ClassId,
		TagsId:     note.TagsId,
		Title:      note.Title,
		Abstract:   note.Abstract,
		Image:      note.Image,
		IsAsterisk: note.IsAsterisk,
		Status:     note.Status,
		CreatedAt:  note.CreatedAt,
		UpdatedAt:  note.UpdatedAt,
		MdContent:  html.UnescapeString(noteDetail.MdContent),
		Content:    html.UnescapeString(noteDetail.Content),
	}

	tags := make([]*model.ArticleDetailResponse_Tag, 0)
	for _, id := range gstr.Split(detail.TagsId, ",") {
		tags = append(tags, &model.ArticleDetailResponse_Tag{Id: id})
	}

	files := make([]*model.ArticleDetailResponse_File, 0)
	items, err := dao.NoteAnnex.AnnexList(ctx, uid, params.ArticleId)
	if err == nil {
		for _, item := range items {
			files = append(files, &model.ArticleDetailResponse_File{
				Id:           item.Id,
				Suffix:       item.Suffix,
				Size:         item.Size,
				OriginalName: item.OriginalName,
				CreatedAt:    util.FormatDatetime(item.CreatedAt),
			})
		}
	}

	return &model.ArticleDetailRes{
		Id:         detail.Id,
		ClassId:    detail.ClassId,
		Title:      detail.Title,
		Content:    detail.Content,
		MdContent:  detail.MdContent,
		IsAsterisk: detail.IsAsterisk,
		CreatedAt:  util.FormatDatetime(detail.CreatedAt),
		UpdatedAt:  util.FormatDatetime(detail.UpdatedAt),
		Tags:       tags,
		Files:      files,
	}, nil
}

// 添加或编辑文章
func (s *sNote) Edit(ctx context.Context, params model.ArticleEditReq) (*model.ArticleEditRes, error) {

	if params.ArticleId == "" || params.ArticleId == "0" { // todo
		opt := &do.NoteCreate{
			UserId:    service.Session().GetUid(ctx),
			ArticleId: params.ArticleId,
			ClassId:   params.ClassId,
			Title:     params.Title,
			Content:   params.Content,
			MdContent: params.MdContent,
		}
		id, err := dao.Note.Create(ctx, opt)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
		params.ArticleId = id
	} else {
		opt := &do.NoteEdit{
			UserId:    service.Session().GetUid(ctx),
			ArticleId: params.ArticleId,
			ClassId:   params.ClassId,
			Title:     params.Title,
			Content:   params.Content,
			MdContent: params.MdContent,
		}
		if err := dao.Note.Update(ctx, opt); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	info, err := dao.Note.FindById(ctx, params.ArticleId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.ArticleEditRes{
		Id:       info.Id,
		Title:    info.Title,
		Abstract: info.Abstract,
		Image:    info.Image,
	}, nil
}

// 删除文章
func (s *sNote) Delete(ctx context.Context, params model.ArticleDeleteReq) error {

	err := dao.Note.UpdateStatus(ctx, service.Session().GetUid(ctx), params.ArticleId, 2)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 恢复文章
func (s *sNote) Recover(ctx context.Context, params model.ArticleRecoverReq) error {

	err := dao.Note.UpdateStatus(ctx, service.Session().GetUid(ctx), params.ArticleId, 1)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 文章图片上传
func (s *sNote) Upload(ctx context.Context) (*model.ArticleUploadImageRes, error) {

	_, file, err := g.RequestFromCtx(ctx).Request.FormFile("image")
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("image 字段必传")
	}

	if !util.Include(util.FileSuffix(file.Filename), []string{"png", "jpg", "jpeg", "gif", "webp"}) {
		return nil, errors.New("上传文件格式不正确,仅支持 png、jpg、jpeg、gif 和 webp")
	}

	// 判断上传文件大小(5M)
	if file.Size > 5<<20 {
		return nil, errors.New("上传文件大小不能超过5M")
	}

	stream, err := filesystem.ReadMultipartStream(file)
	if err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("文件上传失败")
	}

	ext := util.FileSuffix(file.Filename)
	meta := util.ReadImageMeta(bytes.NewReader(stream))

	filePath := fmt.Sprintf("public/media/image/note/%s/%s", util.DateNumber(), util.GenImageName(ext, meta.Width, meta.Height))

	if err := s.Filesystem.Default.Write(stream, filePath); err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("文件上传失败")
	}

	return &model.ArticleUploadImageRes{Url: s.Filesystem.Default.PublicUrl(filePath)}, nil
}

// 文章移动
func (s *sNote) Move(ctx context.Context, params model.ArticleMoveReq) error {

	if err := dao.Note.Move(ctx, service.Session().GetUid(ctx), params.ArticleId, params.ClassId); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 标记文章
func (s *sNote) Asterisk(ctx context.Context, params model.ArticleAsteriskReq) error {

	if err := dao.Note.Asterisk(ctx, service.Session().GetUid(ctx), params.ArticleId, params.Type); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 文章标签
func (s *sNote) Tag(ctx context.Context, params model.ArticleTagsReq) error {

	if err := dao.Note.Tag(ctx, service.Session().GetUid(ctx), params.ArticleId, params.Tags); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 永久删除文章
func (s *sNote) ForeverDelete(ctx context.Context, params model.ArticleForeverDeleteReq) error {

	if err := dao.Note.ForeverDelete(ctx, service.Session().GetUid(ctx), params.ArticleId); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
