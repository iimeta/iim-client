package util

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gorilla/websocket"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/utility/logger"
	"io"
	"net/http"
	"net/url"
	"time"
)

var ProxyOpen bool
var ProxyURL string

func init() {

	ctx := gctx.New()

	proxy_open, err := config.Get(ctx, "http.proxy_open")
	if err != nil {
		logger.Error(ctx, err)
	}
	ProxyOpen = proxy_open.Bool()

	proxy_url, err := config.Get(ctx, "http.proxy_url")
	if err != nil {
		logger.Error(ctx, err)
	}

	ProxyURL = proxy_url.String()
}

func HttpGet(ctx context.Context, url string, header map[string]string, data g.Map, result interface{}, cookies ...map[string]string) error {

	logger.Infof(ctx, "HttpGet url: %s, header: %+v, data: %+v", url, header, data)

	client := g.Client().Timeout(60 * time.Second)
	if header != nil {
		client.SetHeaderMap(header)
	}

	if len(cookies) > 0 {
		client.Cookie(cookies[0])
	}

	response, err := client.Get(ctx, url, data)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	defer func() {
		err = response.Close()
		if err != nil {
			logger.Error(ctx, err)
		}
	}()

	bytes := response.ReadAll()
	logger.Infof(ctx, "HttpGet url: %s, header: %+v, data: %+v, response: %s", url, header, data, string(bytes))

	if bytes != nil && len(bytes) > 0 {
		err = gjson.Unmarshal(bytes, result)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

func HttpPost(ctx context.Context, url string, header map[string]string, data, result interface{}, cookies ...map[string]string) error {

	logger.Infof(ctx, "HttpPost url: %s, header: %+v, data: %+v, cookies: %+v", url, header, data, cookies)

	client := g.Client().Timeout(60 * time.Second)
	if header != nil {
		client.SetHeaderMap(header)
	}

	if len(cookies) > 0 {
		client.Cookie(cookies[0])
	}
	response, err := client.ContentJson().Post(ctx, url, data)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	defer func() {
		err = response.Close()
		if err != nil {
			logger.Error(ctx, err)
		}
	}()

	bytes := response.ReadAll()
	logger.Infof(ctx, "HttpPost url: %s, header: %+v, data: %+v, cookies: %+v, response: %s", url, header, data, cookies, string(bytes))

	if bytes != nil && len(bytes) > 0 {
		err = gjson.Unmarshal(bytes, result)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

func HttpPost2(ctx context.Context, url string, header map[string]string, data, result interface{}, cookies ...map[string]string) error {

	logger.Infof(ctx, "HttpPost url: %s, header: %+v, data: %+v, cookies: %+v", url, header, data, cookies)

	client := g.Client().Timeout(60 * time.Second)
	if header != nil {
		client.SetHeaderMap(header)
	}

	if len(cookies) > 0 {
		client.Cookie(cookies[0])
	}
	response, err := client.Post(ctx, url, data)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	defer func() {
		err = response.Close()
		if err != nil {
			logger.Error(ctx, err)
		}
	}()

	bytes := response.ReadAll()
	logger.Infof(ctx, "HttpPost url: %s, header: %+v, data: %+v, cookies: %+v, response: %s", url, header, data, cookies, string(bytes))

	if bytes != nil && len(bytes) > 0 {
		err = gjson.Unmarshal(bytes, result)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

func WebSocketClientOnlyReceive(ctx context.Context, url string, result chan []byte) (*websocket.Conn, error) {

	logger.Infof(ctx, "WebSocketClientOnlyReceive url: %s", url)

	client := gclient.NewWebSocket()
	//client.HandshakeTimeout = time.Second // 设置超时时间
	//client.Proxy = http.ProxyFromEnvironment // 设置代理
	//client.TLSClientConfig = &tls.Config{}   // 设置 tls 配置

	conn, _, err := client.Dial(url, nil)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	entry := gtimer.AddSingleton(ctx, 30*time.Second, func(ctx context.Context) {
		logger.Debugf(ctx, "WebSocketClientOnlyReceive url: %s, ping...", url)
		err = conn.WriteMessage(websocket.PingMessage, []byte("ping"))
		if err != nil {
			logger.Error(ctx, err)
			return
		}
	})

	_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {

		defer entry.Close()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				logger.Error(ctx, err)
				return
			}
			logger.Infof(ctx, "messageType: %d, message: %s", messageType, string(message))

			_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {
				result <- message
			}, nil)
		}
	}, nil)

	return conn, nil
}

func WebSocketClient(ctx context.Context, url string, messageType int, message []byte, result chan []byte) (*websocket.Conn, error) {

	logger.Infof(ctx, "WebSocketClient url: %s", url)

	client := gclient.NewWebSocket()
	//client.HandshakeTimeout = time.Second // 设置超时时间
	//client.Proxy = http.ProxyFromEnvironment // 设置代理
	//client.TLSClientConfig = &tls.Config{}   // 设置 tls 配置

	conn, _, err := client.Dial(url, nil)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	err = conn.WriteMessage(messageType, message)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				logger.Error(ctx, err)
				return
			}
			logger.Infof(ctx, "messageType: %d, message: %s", messageType, string(message))

			_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {
				result <- message
			}, nil)
		}
	}, nil)

	return conn, nil
}

func SSEServer(ctx context.Context, event string, content any) error {

	r := g.RequestFromCtx(ctx)
	rw := r.Response.RawWriter()
	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return errors.New("Streaming unsupported!")
	}

	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")

	_, err := fmt.Fprintf(rw, "event: %s\ndata: %s\n\n", event, content)
	if err != nil {
		return err
	}

	flusher.Flush()

	return nil
}

func SSEClient(ctx context.Context, method, url string, header map[string]string, data interface{}, result chan []byte, cookies ...map[string]string) error {

	logger.Infof(ctx, "SSEClient method: %s, url: %s, header: %+v, data: %+v, cookies: %+v", method, url, header, data, cookies)

	client := g.Client().Timeout(600 * time.Second)
	if header != nil {
		client.SetHeaderMap(header)
	}

	client.SetHeader("Accept", "text/event-stream")

	if len(cookies) > 0 {
		client.Cookie(cookies[0])
	}

	response, err := client.DoRequest(ctx, method, url, data)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	defer func() {
		err = response.Close()
		if err != nil {
			logger.Error(ctx, err)
		}
	}()

	// 使用bufio.NewReader读取响应正文
	reader := bufio.NewReader(response.Body)

	isClose := false
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				logger.Infof(ctx, "SSEClient method: %s, url: %s, header: %+v, data: %+v, cookies: %+v done", method, url, header, data, cookies)
				return nil
			}
			logger.Error(ctx, err)
			return err
		}

		logger.Infof(ctx, "SSEClient method: %s, url: %s, header: %+v, data: %+v, cookies: %+v, message: %s", method, url, header, data, cookies, message)

		_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {
			result <- []byte(message)
		}, nil)

		if isClose {
			logger.Infof(ctx, "SSEClient method: %s, url: %s, header: %+v, data: %+v, cookies: %+v done", method, url, header, data, cookies)
			return nil
		}

		isClose = message == "event: close"
	}
}

func HttpDownloadFile(ctx context.Context, fileURL string, useProxy ...bool) []byte {

	logger.Infof(ctx, "HttpDownloadFile fileURL: %s", fileURL)

	client := g.Client().Timeout(600 * time.Second)

	transport := &http.Transport{}

	if ProxyOpen && len(ProxyURL) > 0 && (len(useProxy) == 0 || useProxy[0]) {

		logger.Infof(ctx, "HttpDownloadFile ProxyURL: %s", ProxyURL)

		proxyUrl, err := url.Parse(ProxyURL)
		if err != nil {
			logger.Error(ctx, err)
		}

		transport.Proxy = http.ProxyURL(proxyUrl)
		client.Transport = transport
	}

	return client.GetBytes(ctx, fileURL)
}

func isChanClose(ch chan []byte) bool {
	select {
	case _, ok := <-ch:
		return !ok
	default:
	}
	return false
}
