package route

import (
	"KOBA/config"
	"KOBA/controller"
	"KOBA/database"
	"KOBA/repository"
	"KOBA/service"

	"github.com/labstack/echo/v4"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)

	repo := repository.NewRecipe(db)

	svc := service.NewServiceRepo(repo, conf)

	cont := controller.ServiceController{
		Sa: svc,
	}

	apiUser := e.Group("/user")

	apiUser.POST("", cont.CreateUserController)
}
