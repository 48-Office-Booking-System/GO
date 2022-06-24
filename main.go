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

/*

1. pindah ke branch masing masing
2. pull dari development (biar dapet yg paling update)
3. ngerjain masing masing kerjaan di branch
4. kalo udah push ke branch masing masing tersebut

//bisa diurus sama ario yang no. 5
5. checkout ke development dan pull dari branch kita tadi lalu push ke development

*/
