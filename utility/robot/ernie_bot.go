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

type ernieBot struct{}

var ErnieBot *ernieBot

func init() {
	ErnieBot = &ernieBot{}
}

func (o *ernieBot) Chat(ctx context.Context, senderId, receiverId, talkType int, text, model string, mentions ...string) {

	if talkType == 2 {
		content := gstr.Split(text, " ")
		if len(content) > 1 {
			text = content[1]
		}
	}

	if len(text) == 0 {
		return
	}

	messages := make([]sdk.ErnieBotMessage, 0)

	reply, err := g.Redis().LRange(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), 0, -1)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	messagesStr := reply.Strings()

	for _, str := range messagesStr {
		ernieBotMessage := sdk.ErnieBotMessage{}
		if err := json.Unmarshal([]byte(str), &ernieBotMessage); err != nil {
			logger.Error(ctx, err)
			continue
		}
		if ernieBotMessage.Role != openai.ChatMessageRoleSystem {
			messages = append(messages, ernieBotMessage)
		}
	}

	ernieBotMessage := sdk.ErnieBotMessage{
		Role:    sdk.ErnieBotMessageRoleUser,
		Content: text,
	}

	b, err := json.Marshal(ernieBotMessage)
	if err != nil {
		logger.Error(ctx, err)
	}

	logger.Infof(ctx, "ernieBotMessage: %s", string(b))

	messages = append(messages, ernieBotMessage)

	response, err := sdk.ErnieBot(ctx, model, messages)

	if err != nil {
		logger.Error(ctx, err)
		if err = service.TalkMessage().SendText(ctx, senderId, &model2.TextMessageReq{
			Content: err.Error() + ", 发生错误, 请联系作者处理...",
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

	_, err = g.Redis().RPush(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, receiverId, senderId), b)
	if err != nil {
		logger.Error(ctx, err)
	}

	content := response.Result

	ernieBotMessage = sdk.ErnieBotMessage{
		Role:    sdk.ErnieBotMessageRoleAssistant,
		Content: content,
	}

	b, err = json.Marshal(ernieBotMessage)
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
		Receiver: &model2.MessageReceiver{
			TalkType:   talkType,
			ReceiverId: receiverId,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return
	}
}