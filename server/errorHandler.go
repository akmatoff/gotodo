package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError

	message := map[string]string{
		"error": err.Error(),
	}

	if e, ok := err.(*echo.HTTPError); ok {
		code = e.Code
		message["error"] = e.Message.(string)
	}

	c.Logger().Error(err)

	c.JSON(code, message)
}
