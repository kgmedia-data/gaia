package handler

import "github.com/kgmedia-data/gaia/pkg/msg"

type IProcessor[T any] interface {
	Execute(message msg.Message[T]) error
}
