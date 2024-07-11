package msg

import "encoding/json"

type JsonCoder[T any] int

func (j JsonCoder[T]) Encode(data T) ([]byte, error) {
	return json.Marshal(data)
}

func (j JsonCoder[T]) Decode(data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return result, err
}
