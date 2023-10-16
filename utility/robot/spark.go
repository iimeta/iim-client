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
	"github.com/sashabaranov/go-openai"
)

type spark struct{}

var Spark *spark

func init() {
	Spark = &spark{}
}

func (o *spark) Chat(ctx context.Context, senderId, receiverId, talkType int, text, model string, mentions ...string) {

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

	messages := make([]sdk.Text, 0)

	reply, err := g.Redis().LRange(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), 0, -1)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	messagesStr := reply.Strings()

	for _, str := range messagesStr {
		textMessage := sdk.Text{}
		if err := json.Unmarshal([]byte(str), &textMessage); err != nil {
			logger.Error(ctx, err)
			continue
		}
		if textMessage.Role != openai.ChatMessageRoleSystem {
			messages = append(messages, textMessage)
		}
	}

	textMessage := sdk.Text{
		Role:    sdk.SparkMessageRoleUser,
		Content: text,
	}

	b, err := json.Marshal(textMessage)
	if err != nil {
		logger.Error(ctx, err)
	}

	logger.Infof(ctx, "textMessage: %s", string(b))

	messages = append(messages, textMessage)

	response, err := sdk.SparkChat(ctx, model, fmt.Sprintf("%d", receiverId), messages)
	if err != nil {
		logger.Error(ctx, err)
		if err = service.TalkMessage().SendMessage(ctx, &model2.Message{
			MsgType:  consts.MsgTypeText,
			TalkType: talkType,
			Text: &model2.Text{
				Content: err.Error() + ", 发生错误, 请联系作者处理...",
			},
			Sender: &model2.Sender{
				Id: senderId,
			},
			Receiver: &model2.Receiver{
				TalkType:   talkType,
				Id:         receiverId,
				ReceiverId: receiverId,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return
		}
		return
	}

	_, err = g.Redis().RPush(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), b)
	if err != nil {
		logger.Error(ctx, err)
	}

	content := response

	textMessage = sdk.Text{
		Role:    sdk.SparkMessageRoleAssistant,
		Content: content,
	}

	b, err = json.Marshal(textMessage)
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

	if err = service.TalkMessage().SendMessage(ctx, &model2.Message{
		MsgType:  consts.MsgTypeText,
		TalkType: talkType,
		Text: &model2.Text{
			Content: content,
		},
		Sender: &model2.Sender{
			Id: senderId,
		},
		Receiver: &model2.Receiver{
			TalkType:   talkType,
			Id:         receiverId,
			ReceiverId: receiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return
	}
}
