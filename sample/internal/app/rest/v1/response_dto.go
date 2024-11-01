package v1

import (
	"os"

	"github.com/labstack/echo/v4"
)

// ResponseDto is a struct to represent response data in Swagger.
// only used for documentation purpose
type ResponseDto struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
	Version string                 `json:"version"`
}

// DataDoc is a struct to represent data in Swagger
// only used for documentation purpose
type DataDoc struct{}

type ResponseDoc struct {
	Message string  `json:"message"`
	Data    DataDoc `json:"data"`
	Version string  `json:"version"`
}

type ResponsesDto[T any] struct {
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

func NewResponsesDto[T any](msg string, data []T, key string) ResponsesDto[T] {
	version, exist := os.LookupEnv("API_VERSION")
	if !exist {
		version = "-"
	}

	if len(data) > 0 {
		return ResponsesDto[T]{
			Message: msg,
			Data:    map[string][]T{key: data},
			Version: version,
		}
	}
	return ResponsesDto[T]{Message: msg, Data: map[string][]T{key: {}}, Version: version}
}

func unauthorizedResponse(c echo.Context) error {
	resp := NewResponseDto("Unauthorized", nil, "error")
	return c.JSON(401, resp)
}

func forbiddenResponse(c echo.Context) error {
	resp := NewResponseDto("Forbidden", nil, "error")
	return c.JSON(403, resp)
}
