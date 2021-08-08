package config

import (
	"net/http"

	"github.com/fauzanmh/online-store/pkg/util"
	"github.com/labstack/echo/v4"
)

func SetEchoErrorDefault(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if !c.Response().Committed {
			util.ErrorDefaultResponse(c, report.Code, report.Message.(string))
		}
	}
}
