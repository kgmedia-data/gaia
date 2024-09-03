package v1

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseDto_NewResponseDto(t *testing.T) {

	type (
		args struct {
			msg  string
			data interface{}
			key  string
		}
	)

	testCases := []struct {
		desc     string
		args     args
		expected string
	}{
		{
			desc: "success response with empty data",
			args: args{
				msg:  MSG_SUCCESS,
				data: nil,
				key:  "departments",
			},
			expected: `{"message":"success","data":{"departments":{}},"version":"-"}`,
		}, {
			desc: "success response with data",
			args: args{
				msg:  MSG_SUCCESS,
				data: map[string]int{"id": 1, "age": 2},
				key:  "xx",
			},
			expected: `{"message":"success","data":{"xx":{"id":1,"age":2}},"version":"-"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			resp := NewResponseDto(tc.args.msg, tc.args.data, tc.args.key)
			respByte, err := json.Marshal(resp)
			assert.NoError(t, err)
			assert.JSONEq(t, tc.expected, string(respByte))
		})
	}
}

func TestResponseDto_NewResponsesDto(t *testing.T) {

	type (
		args[T any] struct {
			msg  string
			data []T
			key  string
		}
	)

	testCases := []struct {
		desc     string
		args     args[int]
		expected string
	}{
		{
			desc: "success response with empty data",
			args: args[int]{
				msg:  MSG_SUCCESS,
				data: nil,
				key:  "departments",
			},
			expected: `{"message":"success","data":{"departments":[]},"version":"-"}`,
		}, {
			desc: "success response with data",
			args: args[int]{
				msg:  MSG_SUCCESS,
				data: []int{1, 2},
				key:  "xx",
			},
			expected: `{"message":"success","data":{"xx": [1,2]},"version":"-"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			resp := NewResponsesDto(tc.args.msg, tc.args.data, tc.args.key)
			respByte, err := json.Marshal(resp)
			assert.NoError(t, err)
			assert.JSONEq(t, tc.expected, string(respByte))
		})
	}
}
