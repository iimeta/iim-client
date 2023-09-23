package sdk

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
)

var qwenApiKeysRoundrobin = new(util.Roundrobin)

func init() {

	ctx := gctx.New()

	models, err := config.Get(ctx, "aliyun.models")
	if err != nil {
		logger.Error(ctx, err)
		panic(err)
	}

	strVar := models.MapStrVar()

	for model, cfg := range strVar {
		setModel(model, cfg)
	}

}

func getQwenApiKey(ctx context.Context, model string) string {

	logger.Infof(ctx, "model: %s", model)

	cfg := getModel(model).MapStrVar()

	apiKeys := cfg["api_keys"].Strings()

	apiKey := qwenApiKeysRoundrobin.Roundrobin(apiKeys)
	logger.Infof(ctx, "apiKey: %s", apiKey)

	return apiKey
}

type QwenChatCompletionMessage struct {
	User string `json:"user"`
	Bot  string `json:"bot"`
}
type QwenChatCompletionReq struct {
	Model      string `json:"model"`
	Input      Input  `json:"input"`
	Parameters struct {
	} `json:"parameters"`
}
type Input struct {
	Prompt  string                      `json:"prompt"`
	History []QwenChatCompletionMessage `json:"history"`
}
type QwenChatCompletionRes struct {
	Output struct {
		FinishReason string `json:"finish_reason"`
		Text         string `json:"text"`
	} `json:"output"`
	Usage struct {
		OutputTokens int `json:"output_tokens"`
		InputTokens  int `json:"input_tokens"`
	} `json:"usage"`
	RequestId string `json:"request_id"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

func QwenChatCompletion(ctx context.Context, model string, messages []QwenChatCompletionMessage, retry ...int) (*QwenChatCompletionRes, error) {

	logger.Infof(ctx, "model: %s, QwenChatCompletion...", model)

	now := gtime.Now().Unix()

	defer func() {
		logger.Infof(ctx, "QwenChatCompletion 总耗时: %d", gtime.Now().Unix()-now)
	}()

	l := len(messages)
	prompt := messages[l-1].User
	qwenChatCompletionReq := QwenChatCompletionReq{
		Model: model,
		Input: Input{
			Prompt: prompt,
		},
	}

	if l > 1 {
		qwenChatCompletionReq.Input.History = messages[:l-1]
	}

	header := make(map[string]string)
	header["Authorization"] = "Bearer " + getQwenApiKey(ctx, model)

	qwenChatCompletionRes := new(QwenChatCompletionRes)
	err := util.HttpPost(ctx, getModel(model).MapStrVar()["url"].String(), header, qwenChatCompletionReq, &qwenChatCompletionRes)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	logger.Infof(ctx, "QwenChatCompletion response: %s", gjson.MustEncodeString(qwenChatCompletionRes))

	if qwenChatCompletionRes.Code != "" {
		return qwenChatCompletionRes, gerror.Newf("Code: %s, Message: %s, 发生错误, 请联系作者处理...", qwenChatCompletionRes.Code, qwenChatCompletionRes.Message)
	}

	return qwenChatCompletionRes, nil
}
