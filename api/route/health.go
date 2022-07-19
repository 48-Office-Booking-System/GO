package route

import (
	"KOBA/model"

	"github.com/labstack/echo/v4"
)

func HealthAPI(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSONPretty(200, model.Response{
			Code:    200,
			Message: "sehat gan",
		}, "\t")
	})
}
