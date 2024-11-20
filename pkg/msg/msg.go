package msg

import "sync"

type ICoder[T any] interface {
	Encode(data T) ([]byte, error)
	Decode(data []byte) (T, error)
}

type Message[T any] struct {
	Data      T
	Attribute map[string]string
}

type Messages[T any] struct {
	mu       sync.Mutex
	Messages []Message[T]
}

func NewMessages[T any]() Messages[T] {
	return Messages[T]{
		Messages: []Message[T]{},
	}
}

func (m *Messages[T]) Add(msg Message[T]) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Messages = append(m.Messages, msg)
}

func (m *Messages[T]) Flush() []Message[T] {
	m.mu.Lock()
	defer m.mu.Unlock()
	msgs := m.Messages
	m.Messages = []Message[T]{}
	return msgs
}

func (m *Messages[T]) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.Messages)
}
