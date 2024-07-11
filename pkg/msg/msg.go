package msg

type ICoder[T any] interface {
	Encode(data T) ([]byte, error)
	Decode(data []byte) (T, error)
}

type Message[T any] struct {
	Data      T
	Attribute map[string]string
}
