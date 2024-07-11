package v1

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyData(t *testing.T) {
	resp := NewResponseDto(MSG_SUCCESS, EMPTY_OBJ, "employee")
	res, err := json.Marshal(resp)
	assert.NoError(t, err)
	assert.Equal(t, `{"message":"Success","data":{"employee":{}},"version":"-"}`, string(res))

	resp = NewResponseDto(MSG_SUCCESS, EMPTY_LIST, "employees")
	res, err = json.Marshal(resp)
	assert.NoError(t, err)
	assert.Equal(t, `{"message":"Success","data":{"employees":[]},"version":"-"}`, string(res))
}
