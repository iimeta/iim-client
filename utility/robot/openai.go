package robot

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	model2 "github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/sdk"
	"github.com/iimeta/iim-client/utility/util"
	"github.com/sashabaranov/go-openai"
)

type openAI struct{}

var OpenAI *openAI

func init() {
	OpenAI = &openAI{}
}

func (o *openAI) Chat(ctx context.Context, senderId, receiverId, talkType int, text, model string, isOpenContext int, mentions ...string) {

	if talkType == 2 {
		content := gstr.Split(text, " ")
		if len(content) > 1 {
			text = content[1]
		}
	}

	if len(text) == 0 {
		return
	}

	messages := make([]openai.ChatCompletionMessage, 0)

	// 开启上下文
	if isOpenContext == 0 {

		reply, err := redis.LRange(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), 0, -1)
		if err != nil {
			logger.Error(ctx, err)
			return
		}

		messagesStr := reply.Strings()
		if len(messagesStr) == 0 {
			b, err := gjson.Marshal(sdk.ChatMessageRoleSystem)
			if err != nil {
				logger.Error(ctx, err)
			}
			_, err = redis.RPush(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), b)
			if err != nil {
				logger.Error(ctx, err)
			}
			messages = append(messages, sdk.ChatMessageRoleSystem)
		}

		for i, str := range messagesStr {

			chatCompletionMessage := openai.ChatCompletionMessage{}
			if err := gjson.Unmarshal([]byte(str), &chatCompletionMessage); err != nil {
				logger.Error(ctx, err)
				continue
			}

			if i == 0 && chatCompletionMessage.Role != openai.ChatMessageRoleSystem {
				b, err := gjson.Marshal(sdk.ChatMessageRoleSystem)
				if err != nil {
					logger.Error(ctx, err)
				}
				_, err = redis.LPush(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), b)
				if err != nil {
					logger.Error(ctx, err)
				}
				messages = append(messages, sdk.ChatMessageRoleSystem)
			}

			messages = append(messages, chatCompletionMessage)
		}
	} else {
		messages = append(messages, sdk.ChatMessageRoleSystem)
	}

	chatCompletionMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: text,
	}

	b, err := gjson.Marshal(chatCompletionMessage)
	if err != nil {
		logger.Error(ctx, err)
	}

	logger.Infof(ctx, "chatCompletionMessage: %s", string(b))

	messages = append(messages, chatCompletionMessage)

	response, err := sdk.ChatGPTChatCompletion(ctx, model, messages)

	if err != nil {
		logger.Error(ctx, err)

		if gstr.Contains(err.Error(), "Please reduce the length of the messages") {
			start := int64(len(messages) / 2)
			if start > 1 {
				err = redis.LTrim(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), start, -1)
				if err != nil {
					logger.Error(ctx, err)
				} else {
					o.Chat(ctx, senderId, receiverId, talkType, text, model, isOpenContext, mentions...)
					return
				}
			}
		}

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

	_, err = redis.RPush(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), b)
	if err != nil {
		logger.Error(ctx, err)
	}

	content := response.Choices[0].Message.Content

	chatCompletionMessage = openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	}

	b, err = json.Marshal(chatCompletionMessage)
	if err != nil {
		logger.Error(ctx, err)
	}

	_, err = redis.RPush(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), b)
	if err != nil {
		logger.Error(ctx, err)
	}

	if talkType == 2 {
		for i, mention := range mentions {
			if i == 0 {
				content += "\n"
			} else {
				content += " "
			}
			content += "@" + mention
		}
	}

	if err = service.TalkMessage().SendText(ctx, senderId, &model2.TextMessageReq{
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

func (o *openAI) Image(ctx context.Context, senderId, receiverId, talkType int, text string, mentions ...string) {

	if talkType == 2 {
		content := gstr.Split(text, " ")
		if len(content) > 1 {
			text = content[1]
		}
	}

	if len(text) == 0 {
		return
	}

	logger.Infof(ctx, "Image text: %s", text)

	imgBase64, err := sdk.GenImageBase64(ctx, text)
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

	imgBytes, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	imageInfo, err := util.SaveImage(ctx, imgBytes, ".png")
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	domain, err := config.Get(ctx, "filesystem.local.domain")
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	url := domain.String() + "/" + imageInfo.FilePath

	if err := service.TalkMessage().SendImage(ctx, senderId, &model2.ImageMessageReq{
		Url:    url,
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
}
