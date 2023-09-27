package talk

import (
	"context"
	"github.com/iimeta/iim-client/api/talk/v1"
	"github.com/iimeta/iim-client/internal/service"
)

func (c *ControllerV1) RecordsFileDownload(ctx context.Context, req *v1.RecordsFileDownloadReq) (res *v1.RecordsFileDownloadRes, err error) {

	err = service.TalkRecords().Download(ctx, req.RecordId)

	return
}
