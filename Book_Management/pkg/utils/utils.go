package utils

import (
	"github.com/labstack/echo/v4"
)

func ParseBody(c echo.Context, x interface{}) {
	if err := c.Bind(x); err != nil {
		return
	}
}
