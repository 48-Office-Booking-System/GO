package main

import (
	"KOBA/api/route"
	"KOBA/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := config.InitConfiguration()

	e := echo.New()
	e.Use(middleware.CORS())

	route.HealthAPI(e, conf)
	route.RegisterUserGroupAPI(e, conf)
	route.RegisterOfficeGroupAPI(e, conf)
	route.RegisterBookingGroupAPI(e, conf)

	e.Logger.Fatal(e.Start(config.InitConfiguration().SERVER_ADDRESS))
}
