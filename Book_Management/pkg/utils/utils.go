package utils

import (
	"encoding/json"
	"io"
	"github.com/labstack/echo/v4"
)

func ParseBody(c echo.Context, x interface{}) {
	if body, err := io.ReadAll(c.Request().Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
