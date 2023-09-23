package sdk

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
)

func init() {

	ctx := gctx.New()

	models, err := config.Get(ctx, "ernie_bot.models")
	if err != nil {
		logger.Error(ctx, err)
		panic(err)
	}

	strVar := models.MapStrVar()

	for model, cfg := range strVar {
		setModel(model, cfg)
	}
}

const ErnieBotMessageRoleUser = "user"
const ErnieBotMessageRoleAssistant = "assistant"
const ERNIE_BOT_ACCESS_TOKEN_KEY = "ernie_bot:access_token"

type ErnieBotMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ErnieBotReq struct {
	Messages []ErnieBotMessage `json:"messages"`
}
type ErnieBotRes struct {
	Id               string `json:"id"`
	Object           string `json:"object"`
	Created          int    `json:"created"`
	Result           string `json:"result"`
	IsTruncated      bool   `json:"is_truncated"`
	NeedClearHistory bool   `json:"need_clear_history"`
	Usage            struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func ErnieBot(ctx context.Context, model string, messages []ErnieBotMessage, retry ...int) (*ErnieBotRes, error) {

	logger.Infof(ctx, "model: %s, ErnieBot...", model)

	now := gtime.Now().Unix()

	defer func() {
		logger.Infof(ctx, "ErnieBot 总耗时: %d", gtime.Now().Unix()-now)
	}()

	req := ErnieBotReq{
		Messages: messages,
	}

	ernieBotRes := new(ErnieBotRes)
	err := util.HttpPost(ctx, fmt.Sprintf("%s?access_token=%s", getModel(model).MapStrVar()["url"].String(), GetAccessToken(ctx)), nil, req, &ernieBotRes)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	logger.Infof(ctx, "ErnieBot ernieBotRes: %s", gjson.MustEncodeString(ernieBotRes))

	if ernieBotRes.ErrorCode != 0 {
		return nil, gerror.Newf("ErrorCode: %d, ErrorMsg: %s", ernieBotRes.ErrorCode, ernieBotRes.ErrorMsg)
	}

	return ernieBotRes, nil
}

type GetAccessTokenRes struct {
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int64  `json:"expires_in"`
	SessionKey       string `json:"session_key"`
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	SessionSecret    string `json:"session_secret"`
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

func GetAccessToken(ctx context.Context, appid ...string) string {

	reply, err := g.Redis().Get(ctx, ERNIE_BOT_ACCESS_TOKEN_KEY)
	if err == nil && reply.String() != "" {
		return reply.String()
	}

	access_token_url, err := config.Get(ctx, "ernie_bot.access_token_url")
	if err != nil {
		logger.Error(ctx, err)
		return ""
	}

	apps, err := config.Get(ctx, "ernie_bot.apps")
	if err != nil {
		logger.Error(ctx, err)
		return ""
	}

	appMap := make(map[string]string)

	if len(appid) > 0 {
		appMap = apps.MapStrVar()[appid[0]].MapStrStr()
	} else {
		for _, value := range apps.Vars()[0].MapStrVar() {
			appMap = value.MapStrStr()
		}
	}

	data := g.Map{
		"grant_type":    "client_credentials",
		"client_id":     appMap["api_key"],
		"client_secret": appMap["secret_key"],
	}

	getAccessTokenRes := new(GetAccessTokenRes)
	err = util.HttpPost2(ctx, access_token_url.String(), nil, data, &getAccessTokenRes)
	if err != nil {
		logger.Error(ctx, err)
		return ""
	}

	if getAccessTokenRes.Error != "" {
		logger.Error(ctx, getAccessTokenRes.Error)
		return ""
	}

	_ = g.Redis().SetEX(ctx, ERNIE_BOT_ACCESS_TOKEN_KEY, getAccessTokenRes.AccessToken, getAccessTokenRes.ExpiresIn)

	return getAccessTokenRes.AccessToken
}
