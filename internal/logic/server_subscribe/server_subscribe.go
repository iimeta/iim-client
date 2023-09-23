package server_subscribe

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/iimeta/iim-client/internal/config"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/cache"
	"github.com/iimeta/iim-client/utility/logger"
	redis2 "github.com/iimeta/iim-client/utility/redis"
	"github.com/iimeta/iim-client/utility/socket"
	"github.com/iimeta/iim-client/utility/socket/adapter"
	"github.com/redis/go-redis/v9"
	"github.com/sourcegraph/conc/pool"
	"golang.org/x/sync/errgroup"
	"net/http"
	"sync"
	"time"
)

type sServerSubscribe struct {
	clientStorage *cache.ClientStorage
	serverStorage *cache.ServerStorage
}

func init() {
	service.RegisterServerSubscribe(New())
}

func New() service.IServerSubscribe {
	return &sServerSubscribe{
		clientStorage: cache.NewClientStorage(redis2.Client, config.Cfg, cache.NewSidStorage(redis2.Client)),
		serverStorage: cache.NewSidStorage(redis2.Client),
	}
}

// 初始化连接
func (s *sServerSubscribe) Conn(w http.ResponseWriter, r *http.Request) error {

	conn, err := adapter.NewWsAdapter(w, r)
	if err != nil {
		logger.Error(r.Context(), "websocket connect error:", err)
		return err
	}

	return s.NewClient(service.Session().GetUid(r.Context()), conn)
}

func (s *sServerSubscribe) NewClient(uid int, conn socket.IConn) error {

	return socket.NewClient(conn, &socket.ClientOption{
		Uid:     uid,
		Channel: socket.Session.Chat,
		Storage: s.clientStorage,
		Buffer:  10,
	}, socket.NewEvent(
		// 连接成功回调
		socket.WithOpenEvent(service.ServerEvent().OnOpen),
		// 接收消息回调
		socket.WithMessageEvent(service.ServerEvent().OnMessage),
		// 关闭连接回调
		socket.WithCloseEvent(service.ServerEvent().OnClose),
	))
}

// todo 不加这个消息会根据会话数量推送重复条数, 不太懂啊...
var once sync.Once

// Start 启动服务
func (s *sServerSubscribe) Start(ctx context.Context, eg *errgroup.Group) {
	once.Do(func() {
		eg.Go(func() error {
			return s.SetupHealthSubscribe(ctx) // 注册健康上报
		})
		eg.Go(func() error {
			return s.SetupMessageSubscribe(ctx) // 注册消息订阅
		})
	})
}

// 注册健康上报
func (s *sServerSubscribe) SetupHealthSubscribe(ctx context.Context) error {

	logger.Info(ctx, "Start HealthSubscribe")

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(10 * time.Second):
			if err := s.serverStorage.Set(ctx, config.Cfg.ServerId(), time.Now().Unix()); err != nil {
				logger.Error(ctx, "Websocket HealthSubscribe Report Err:", err)
			}
		}
	}
}

// 注册消息订阅
func (s *sServerSubscribe) SetupMessageSubscribe(ctx context.Context) error {

	logger.Info(ctx, "Start MessageSubscribe")

	_ = grpool.AddWithRecover(gctx.New(), func(ctx context.Context) {
		s.subscribe(ctx, []string{consts.ImTopicChat, fmt.Sprintf(consts.ImTopicChatPrivate, config.Cfg.ServerId())}, service.ServerConsume())
	}, nil)

	<-ctx.Done()

	return nil
}

func (s *sServerSubscribe) subscribe(ctx context.Context, topic []string, consume service.IServerConsume) {

	sub := redis2.Client.Subscribe(ctx, topic...)
	defer func() {
		err := sub.Close()
		if err != nil {
			logger.Error(ctx, err)
		}
	}()

	worker := pool.New().WithMaxGoroutines(10)

	for data := range sub.Channel() {
		s.handle(ctx, worker, data, consume)
	}

	worker.Wait()
}

func (s *sServerSubscribe) handle(ctx context.Context, worker *pool.Pool, data *redis.Message, consume service.IServerConsume) {
	worker.Go(func() {
		var in model.SubscribeContent
		if err := json.Unmarshal([]byte(data.Payload), &in); err != nil {
			logger.Error(ctx, "SubscribeContent Unmarshal Err:", err)
			return
		}

		defer func() {
			if err := recover(); err != nil {
				logger.Error(ctx, "MessageSubscribe Call Err:", err)
			}
		}()

		consume.Call(ctx, in.Event, []byte(in.Data))
	})
}
