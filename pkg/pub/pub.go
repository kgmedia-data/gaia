package pub

import "github.com/kgmedia-data/gaia/pkg/msg"

type IPublisher[T any] interface {
	Publish(message msg.Message[T]) error
	Close() error
}
