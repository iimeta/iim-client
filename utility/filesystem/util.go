package filesystem

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/iim-client/utility/logger"
	"io"
	"mime/multipart"
)

func ReadMultipartStream(file *multipart.FileHeader) ([]byte, error) {

	src, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer func() {
		err := src.Close()
		if err != nil {
			logger.Error(gctx.New(), err)
		}
	}()

	return io.ReadAll(src)
}
