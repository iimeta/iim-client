package util

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iimeta/iim-client/utility/logger"
	"image"
	"io/fs"
	"os"
)

type ImageInfo struct {
	Md5Sum   string
	FilePath string
	Size     int
	Width    int
	Height   int
}

func SaveImage(ctx context.Context, imgBytes []byte, ext string, fileName ...string) (*ImageInfo, error) {

	basePath := fmt.Sprintf("public/media/image/talk/%s", DateNumber())

	err := os.MkdirAll("./resource/"+basePath, fs.ModePerm)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	md5Sum := fmt.Sprintf("%x", md5.Sum(imgBytes))

	var filePath string

	if len(fileName) > 0 {
		filePath = fmt.Sprintf("%s/%s", basePath, fileName[0])
	} else {
		filePath = fmt.Sprintf("%s/%s%s", basePath, md5Sum, ext)
	}

	file, err := os.Create("./resource/" + filePath)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			logger.Error(ctx, err)
		}
	}()

	size, err := file.Write(imgBytes)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	width := 512
	height := 512
	if ext != ".webp" {
		reader, err := os.Open("./resource/" + filePath)
		if err != nil {
			logger.Error(ctx, err)
		} else {

			defer func() {
				err := reader.Close()
				if err != nil {
					logger.Error(ctx, err)
				}
			}()

			img, _, err := image.Decode(reader)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}

			bounds := img.Bounds()
			if bounds.Dx() != 0 {
				width = bounds.Dx()
			}

			if bounds.Dy() != 0 {
				height = bounds.Dy()
			}
		}
	}

	imageInfo := &ImageInfo{
		Md5Sum:   md5Sum,
		FilePath: filePath,
		Size:     size,
		Width:    width,
		Height:   height,
	}

	logger.Infof(ctx, "SaveImage imageInfo: %s, size: %d", gjson.MustEncodeString(imageInfo), imageInfo.Size)

	return imageInfo, nil
}
