package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func requestTestHelper[T any](method string, data T, queryParam string) (echo.Context,
	*httptest.ResponseRecorder) {

	e := echo.New()
	var req *http.Request
	url := fmt.Sprintf("/%v", queryParam)

	if &data != nil {
		jsonData, _ := json.Marshal(data)
		req = httptest.NewRequest(method, url, bytes.NewReader(jsonData))
	} else {
		req = httptest.NewRequest(method, url, nil)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}
