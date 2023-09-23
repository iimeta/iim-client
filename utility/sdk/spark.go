package sdk

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gorilla/websocket"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-client/utility/util"
	"net/url"
	"time"
)

const SparkMessageRoleUser = "user"
const SparkMessageRoleAssistant = "assistant"

type Header struct {
	// req
	AppId string `json:"app_id"`
	Uid   string `json:"uid"`
	// res
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Sid     string `json:"sid,omitempty"`
	Status  int    `json:"status,omitempty"`
}

type Parameter struct {
	// req
	Chat *Chat `json:"chat"`
}

type Chat struct {
	// req
	Domain          string `json:"domain"`
	RandomThreshold int    `json:"random_threshold"`
	MaxTokens       int    `json:"max_tokens"`
	Auditing        string `json:"auditing"`
}

type Payload struct {
	// req
	Message *Message `json:"message"`
	// res
	Choices *Choices `json:"choices,omitempty"`
	Usage   *Usage   `json:"usage,omitempty"`
}

type Message struct {
	// req
	Text []Text `json:"text"`
}

type Text struct {
	// req res
	Role    string `json:"role"`
	Content string `json:"content"`

	// Choices
	Index int `json:"index,omitempty"`

	// Usage
	QuestionTokens   int `json:"question_tokens,omitempty"`
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}

type Choices struct {
	// res
	Status int    `json:"status,omitempty"`
	Seq    int    `json:"seq,omitempty"`
	Text   []Text `json:"text,omitempty"`
}

type Usage struct {
	// res
	Text *Text `json:"text,omitempty"`
}

type SparkReq struct {
	Header    Header    `json:"header"`
	Parameter Parameter `json:"parameter"`
	Payload   Payload   `json:"payload"`
}
type SparkRes struct {
	Header  Header  `json:"header"`
	Payload Payload `json:"payload"`
}

func SparkChat(ctx context.Context, model string, text []Text, retry ...int) (string, error) {

	app_id, err := config.Get(ctx, "xfyun.spark.app_id")
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	domain, err := config.Get(ctx, "xfyun.spark.domain")
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	sparkReq := SparkReq{
		Header: Header{
			AppId: app_id.String(),
			Uid:   "loka",
		},
		Parameter: Parameter{
			Chat: &Chat{
				Domain:          domain.String(),
				RandomThreshold: 0,
				MaxTokens:       4096,
				Auditing:        "default",
			},
		},
		Payload: Payload{
			Message: &Message{
				Text: text,
			},
		},
	}

	data, err := gjson.Marshal(sparkReq)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	url := getAuthorizationUrl(ctx)

	logger.Infof(ctx, "getAuthorizationUrl: %s", url)

	result := make(chan []byte)
	var conn *websocket.Conn

	_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {
		conn, err = util.WebSocketClient(ctx, url, websocket.TextMessage, data, result)
		if err != nil {
			logger.Error(ctx, err)
		}
	}, nil)

	defer func() {
		err = conn.Close()
		if err != nil {
			logger.Error(ctx, err)
		}
	}()

	responseContent := ""
	for {
		select {
		case message := <-result:

			sparkRes := new(SparkRes)
			err := gjson.Unmarshal(message, &sparkRes)
			if err != nil {
				logger.Error(ctx, err)
				return "", err
			}

			if sparkRes.Header.Code != 0 {
				return "", errors.New(gjson.MustEncodeString(sparkRes))
			}

			responseContent += sparkRes.Payload.Choices.Text[0].Content

			if sparkRes.Header.Status == 2 {
				logger.Infof(ctx, "responseContent: %s", responseContent)
				return responseContent, nil
			}
		}
	}
}

func SparkStreaming(ctx context.Context, model string, text []Text, responseContent chan Payload, retry ...int) {

	app_id, err := config.Get(ctx, "xfyun.spark.app_id")
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	domain, err := config.Get(ctx, "xfyun.spark.domain")
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	sparkReq := SparkReq{
		Header: Header{
			AppId: app_id.String(),
			Uid:   "loka",
		},
		Parameter: Parameter{
			Chat: &Chat{
				Domain:          domain.String(),
				RandomThreshold: 0,
				MaxTokens:       4096,
				Auditing:        "default",
			},
		},
		Payload: Payload{
			Message: &Message{
				Text: text,
			},
		},
	}

	data, err := gjson.Marshal(sparkReq)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	url := getAuthorizationUrl(ctx)

	logger.Infof(ctx, "getAuthorizationUrl: %s", url)

	result := make(chan []byte)
	var conn *websocket.Conn

	_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {
		conn, err = util.WebSocketClient(ctx, url, websocket.TextMessage, data, result)
		if err != nil {
			logger.Error(ctx, err)
		}
	}, nil)

	defer func() {
		err = conn.Close()
		if err != nil {
			logger.Error(ctx, err)
		}
	}()

	for {
		select {
		case message := <-result:

			sparkRes := new(SparkRes)
			err := gjson.Unmarshal(message, &sparkRes)
			if err != nil {
				logger.Error(ctx, err)
				return
			}

			responseContent <- sparkRes.Payload

			if sparkRes.Header.Status == 2 {
				return
			}
		}
	}
}

func getAuthorizationUrl(ctx context.Context) string {

	api_key, err := config.Get(ctx, "xfyun.spark.api_key")
	if err != nil {
		logger.Error(ctx, err)
		return ""
	}

	api_secret, err := config.Get(ctx, "xfyun.spark.api_secret")
	if err != nil {
		logger.Error(ctx, err)
		return ""
	}

	chat_url, err := config.Get(ctx, "xfyun.spark.chat_url")
	if err != nil {
		logger.Error(ctx, err)
		return ""
	}

	parse, err := url.Parse(chat_url.String())
	if err != nil {
		logger.Error(ctx, err)
		return ""
	}

	now := gtime.Now()
	loc, _ := time.LoadLocation("GMT")
	zone, _ := now.ToZone(loc.String())
	date := zone.Layout("Mon, 02 Jan 2006 15:04:05 GMT")

	tmp := "host: " + parse.Host + "\n"
	tmp += "date: " + date + "\n"
	tmp += "GET " + parse.Path + " HTTP/1.1"

	hash := hmac.New(sha256.New, api_secret.Bytes())

	_, err = hash.Write([]byte(tmp))
	if err != nil {
		logger.Error(ctx, err)
		return ""
	}

	signature := gbase64.EncodeToString(hash.Sum(nil))

	authorizationOrigin := gbase64.EncodeToString([]byte(fmt.Sprintf("api_key=\"%s\",algorithm=\"%s\",headers=\"%s\",signature=\"%s\"", api_key.String(), "hmac-sha256", "host date request-line", signature)))

	wsURL := gstr.Replace(gstr.Replace(chat_url.String(), "https://", "wss://"), "http://", "ws://")

	return fmt.Sprintf("%s?authorization=%s&date=%s&host=%s", wsURL, authorizationOrigin, gurl.RawEncode(date), parse.Host)
}
