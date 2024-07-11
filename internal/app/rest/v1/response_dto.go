package v1

import (
	"os"

	"github.com/labstack/echo/v4"
)

const (
	MSG_SUCCESS = "Success"
)

var EMPTY_OBJ = map[string]interface{}{}
var EMPTY_LIST = []interface{}{}

type ResponseDto struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
	Version string                 `json:"version"`
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
	return ResponseDto{Message: msg, Data: map[string]interface{}{key: EMPTY_OBJ}, Version: version}
}

func unauthorizedResponse(c echo.Context) error {
	resp := NewResponseDto("Unauthorized", nil, "error")
	return c.JSON(401, resp)
}

func forbiddenResponse(c echo.Context) error {
	resp := NewResponseDto("Forbidden", nil, "error")
	return c.JSON(403, resp)
}
