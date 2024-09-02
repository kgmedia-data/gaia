package v1

import (
	"os"

	"github.com/labstack/echo/v4"
)

type ResponseDto struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
	Version string                 `json:"version"`
}

type ResponseDtos[T any] struct {
	Message string         `json:"message"`
	Data    map[string][]T `json:"data"`
	Version string         `json:"version"`
}

func NewResponseDto(msg string, data interface{}, key string) ResponseDto {
	version, exist := os.LookupEnv("API_VERSION")
	if !exist {
		version = "-"
	}

	if data != nil {
		return ResponseDto{
			Message: msg,
			Data:    map[string]interface{}{key: data},
			Version: version,
		}
	}
	return ResponseDto{Message: msg, Data: map[string]interface{}{key: map[string]interface{}{}}, Version: version}
}

func NewResponseDtos[T any](msg string, data []T, key string) ResponseDtos[T] {
	version, exist := os.LookupEnv("API_VERSION")
	if !exist {
		version = "-"
	}

	if len(data) > 0 {
		return ResponseDtos[T]{
			Message: msg,
			Data:    map[string][]T{key: data},
			Version: version,
		}
	}
	return ResponseDtos[T]{Message: msg, Data: map[string][]T{key: {}}, Version: version}
}

func unauthorizedResponse(c echo.Context) error {
	resp := NewResponseDto("Unauthorized", nil, "error")
	return c.JSON(401, resp)
}

func forbiddenResponse(c echo.Context) error {
	resp := NewResponseDto("Forbidden", nil, "error")
	return c.JSON(403, resp)
}
