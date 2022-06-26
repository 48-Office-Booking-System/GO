package main

import (
	"KOBA/api/route"
	"KOBA/config"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfiguration()

	e := echo.New()

	route.HealthAPI(e, conf)
	route.RegisterUserGroupAPI(e, conf)
	route.RegisterBookingGroupAPI(e, conf)

	e.Logger.Fatal(e.Start(config.InitConfiguration().SERVER_ADDRESS))
}
