package msg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type xx struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestJsonCoder(t *testing.T) {
	coder := JsonCoder[xx](1)
	x := xx{
		ID:   1,
		Name: "name x",
	}

	data, err := coder.Encode(x)
	assert.NoError(t, err)

	back, err := coder.Decode(data)
	assert.NoError(t, err)

	assert.Equal(t, x, back)
}
