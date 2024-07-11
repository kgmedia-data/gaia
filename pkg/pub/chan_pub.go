package pub

import (
	"fmt"

	"github.com/kgmedia-data/gaia/pkg/msg"
)

type ChanPublisher[T any] struct {
	subs []chan msg.Message[T]
}

func NewChanPublisher[T any](subscription ...chan msg.Message[T]) *ChanPublisher[T] {
	return &ChanPublisher[T]{
		subs: subscription,
	}
}

func (t *ChanPublisher[T]) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("ChanPublisher[T].(%v)(%v) %w", method, params, err)
}

func (c *ChanPublisher[T]) Publish(message msg.Message[T]) error {
	for _, channel := range c.subs {
		c.send(channel, message)
	}
	return nil
}

func (c *ChanPublisher[T]) send(msgChan chan msg.Message[T], message msg.Message[T]) {
	msgChan <- message
}

func (c *ChanPublisher[T]) Close() error {
	return nil
}
