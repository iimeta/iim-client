package robot

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/iim-client/internal/config"
	model2 "github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/sdk"
	"github.com/iimeta/iim-client/utility/util"
	"net/url"
	"strings"
)

type midjourney struct{}

var Midjourney *midjourney

func init() {
	Midjourney = &midjourney{}
}

func (m *midjourney) Image(ctx context.Context, senderId, receiverId, talkType int, text string, proxy string) {

	if talkType == 2 {
		content := gstr.Split(text, " ")
		if len(content) > 1 {
			text = content[1]
		}
	}

	if len(text) == 0 {
		return
	}

	logger.Infof(ctx, "Midjourney Image prompt: %s", text)

	text = gstr.Replace(text, "\n", "")
	text = gstr.Replace(text, "\r", "")
	text = gstr.TrimLeftStr(text, "/mj")
	text = gstr.TrimLeftStr(text, "/imagine")
	text = strings.TrimSpace(text)

	if err := service.TalkMessage().SendText(ctx, senderId, &model2.TextMessageReq{
		Content: "您的请求已收到, 请耐心等待1-5分钟, 精彩马上为您呈献...",
		Receiver: &model2.MessageReceiver{
			TalkType:   talkType,
			ReceiverId: receiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return
	}

	imageURL := ""
	taskId := ""
	var err error
	var imageInfo *util.ImageInfo

	switch proxy {
	case "midjourney_proxy":
		if gstr.HasPrefix(text, "UPSCALE") || gstr.HasPrefix(text, "VARIATION") {
			taskId, imageInfo, imageURL, err = sdk.MidjourneyProxyChanges(ctx, text)
		} else {
			taskId, imageInfo, imageURL, err = sdk.MidjourneyProxy(ctx, text)
		}
	}

	logger.Infof(ctx, "Midjourney Image imageURL: %s", imageURL)

	if err != nil || imageURL == "" {
		logger.Error(ctx, err)
		if err = service.TalkMessage().SendText(ctx, senderId, &model2.TextMessageReq{
			Content: err.Error(),
			Receiver: &model2.MessageReceiver{
				TalkType:   talkType,
				ReceiverId: receiverId,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return
		}
		return
	}

	if imageInfo == nil {

		cdn_url, err := config.Get(ctx, "midjourney.cdn_url")
		if err != nil {
			logger.Error(ctx, err)
		}

		if cdn_url.String() != "" {

			imageInfo = &util.ImageInfo{
				Size:   1024 * 1024 * 5,
				Width:  512,
				Height: 512,
			}

			_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {

				imgBytes := util.HttpDownloadFile(ctx, imageURL, false)

				if len(imgBytes) != 0 {
					_, err = util.SaveImage(ctx, imgBytes, gfile.Ext(imageURL), gfile.Basename(imageURL))
					if err != nil {
						logger.Error(ctx, err)
						return
					}
				} else {
					logger.Errorf(ctx, "HttpDownloadFile %s fail", imageURL)
				}

			}, nil)

			originalUrl, err := url.Parse(imageURL)
			if err != nil {
				logger.Error(ctx, err)
				if err = service.TalkMessage().SendText(ctx, senderId, &model2.TextMessageReq{
					Content: err.Error(),
					Receiver: &model2.MessageReceiver{
						TalkType:   talkType,
						ReceiverId: receiverId,
					},
				}); err != nil {
					logger.Error(ctx, err)
					return
				}
				return
			}

			// 替换CDN
			imageURL = cdn_url.String() + originalUrl.Path

		} else {

			imgBytes := util.HttpDownloadFile(ctx, imageURL, false)

			if len(imgBytes) == 0 {
				if err = service.TalkMessage().SendText(ctx, senderId, &model2.TextMessageReq{
					Content: err.Error(),
					Receiver: &model2.MessageReceiver{
						TalkType:   talkType,
						ReceiverId: receiverId,
					},
				}); err != nil {
					logger.Error(ctx, err)
					return
				}
				return
			}

			imageInfo, err = util.SaveImage(ctx, imgBytes, gfile.Ext(imageURL))
			if err != nil {
				logger.Error(ctx, err)
				return
			}

			domain, err := config.Get(ctx, "filesystem.local.domain")
			if err != nil {
				logger.Error(ctx, err)
				return
			}

			imageURL = domain.String() + "/" + imageInfo.FilePath
		}
	}

	logger.Infof(ctx, "SendImage imageURL: %s, Width: %d, Height: %d, Size: %d", imageURL, imageInfo.Width, imageInfo.Height, imageInfo.Size)

	if err := service.TalkMessage().SendImage(ctx, senderId, &model2.ImageMessageReq{
		Url:    imageURL,
		Width:  imageInfo.Width,
		Height: imageInfo.Height,
		Size:   imageInfo.Size,
		Receiver: &model2.MessageReceiver{
			TalkType:   talkType,
			ReceiverId: receiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return
	}

	if !gstr.HasPrefix(text, "UPSCALE") {

		content := fmt.Sprintf("Prompt: %s\n", text)
		content += fmt.Sprintf("Result: 任务ID: %s, 您可以回复以下内容对图片进行相应操作:\n", taskId)
		content += fmt.Sprintf("第一张: UPSCALE::1::%s , VARIATION::1::%s\n", taskId, taskId)
		content += fmt.Sprintf("第二张: UPSCALE::2::%s , VARIATION::2::%s\n", taskId, taskId)
		content += fmt.Sprintf("第三张: UPSCALE::3::%s , VARIATION::3::%s\n", taskId, taskId)
		content += fmt.Sprintf("第四张: UPSCALE::4::%s , VARIATION::4::%s\n", taskId, taskId)
		content += fmt.Sprintf("操作说明: UPSCALE为放大, VARIATION为微调, 示例: UPSCALE::1::%s\n", taskId)

		if err := service.TalkMessage().SendText(ctx, senderId, &model2.TextMessageReq{
			Content: content,
			Receiver: &model2.MessageReceiver{
				TalkType:   talkType,
				ReceiverId: receiverId,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return
		}
	}
}
