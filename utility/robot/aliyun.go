package robot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/iim-client/internal/consts"
	model2 "github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/sdk"
)

type aliyun struct{}

var Aliyun *aliyun

func init() {
	Aliyun = &aliyun{}
}

func (o *aliyun) Chat(ctx context.Context, senderId, receiverId, talkType int, text, model string, mentions ...string) {

	if talkType == 2 {
		content := gstr.Split(text, " ")
		if len(content) > 1 {
			text = content[1]
		} else {
			content = gstr.Split(text, " ")
			if len(content) > 1 {
				text = content[1]
			}
		}
	}

	if len(text) == 0 {
		return
	}

	messages := make([]sdk.QwenChatCompletionMessage, 0)

	reply, err := g.Redis().LRange(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), 0, -1)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	messagesStr := reply.Strings()

	for _, str := range messagesStr {
		qwenChatCompletionMessage := sdk.QwenChatCompletionMessage{}
		if err := json.Unmarshal([]byte(str), &qwenChatCompletionMessage); err != nil {
			logger.Error(ctx, err)
			continue
		}
		messages = append(messages, qwenChatCompletionMessage)
	}

	qwenChatCompletionMessage := sdk.QwenChatCompletionMessage{
		User: text,
	}

	b, err := json.Marshal(qwenChatCompletionMessage)
	if err != nil {
		logger.Error(ctx, err)
	}

	logger.Infof(ctx, "qwenChatCompletionMessage: %s", string(b))

	messages = append(messages, qwenChatCompletionMessage)

	response, err := sdk.QwenChatCompletion(ctx, model, messages)

	if err != nil {
		logger.Error(ctx, err)
		if err = service.TalkMessage().SendText(ctx, senderId, &model2.TextMessageReq{
			Content: err.Error(),
			Receiver: &model2.Receiver{
				TalkType:   talkType,
				ReceiverId: receiverId,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return
		}
		return
	}

	content := response.Output.Text

	qwenChatCompletionMessage.Bot = content

	b, err = json.Marshal(qwenChatCompletionMessage)
	if err != nil {
		logger.Error(ctx, err)
	}

	_, err = g.Redis().RPush(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), b)
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
		Receiver: &model2.Receiver{
			TalkType:   talkType,
			ReceiverId: receiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return
	}
}
