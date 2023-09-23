package socket

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/iim-client/utility/logger"
)

type IEvent interface {
	Open(client IClient)
	Message(client IClient, data []byte)
	Close(client IClient, code int, text string)
	Destroy(client IClient)
}

type (
	OpenEvent    func(client IClient)
	MessageEvent func(client IClient, data []byte)
	CloseEvent   func(client IClient, code int, text string)
	DestroyEvent func(client IClient)
	EventOption  func(event *Event)
)

type Event struct {
	open    OpenEvent
	message MessageEvent
	close   CloseEvent
	destroy DestroyEvent
}

func NewEvent(opts ...EventOption) IEvent {

	o := &Event{}

	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (c *Event) Open(client IClient) {

	if c.open == nil {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			logger.Error(gctx.New(), "open event callback exception: ", client.Uid(), client.Cid(), client.Channel().Name(), err)
		}
	}()

	c.open(client)
}

func (c *Event) Message(client IClient, data []byte) {

	if c.message == nil {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			logger.Error(gctx.New(), "message event callback exception: ", client.Uid(), client.Cid(), client.Channel().Name(), err)
		}
	}()

	c.message(client, data)
}

func (c *Event) Close(client IClient, code int, text string) {

	if c.close == nil {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			logger.Error(gctx.New(), "close event callback exception: ", client.Uid(), client.Cid(), client.Channel().Name(), err)
		}
	}()

	c.close(client, code, text)
}

func (c *Event) Destroy(client IClient) {

	if c.destroy == nil {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			logger.Error(gctx.New(), "destroy event callback exception: ", client.Uid(), client.Cid(), client.Channel().Name(), err)
		}
	}()

	c.destroy(client)
}

// WithOpenEvent 连接成功回调事件
func WithOpenEvent(e OpenEvent) EventOption {
	return func(event *Event) {
		event.open = e
	}
}

// WithMessageEvent 消息回调事件
func WithMessageEvent(e MessageEvent) EventOption {
	return func(event *Event) {
		event.message = e
	}
}

// WithCloseEvent 连接关闭回调事件
func WithCloseEvent(e CloseEvent) EventOption {
	return func(event *Event) {
		event.close = e
	}
}

// WithDestroyEvent 连接销毁回调事件
func WithDestroyEvent(e DestroyEvent) EventOption {
	return func(event *Event) {
		event.destroy = e
	}
}
