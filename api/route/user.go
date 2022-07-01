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

	repo := repository.NewUser(db)

	svc := service.NewUserService(repo, conf)

	cont := controller.UserServiceController{
		UserServ: svc,
	}

	apiUser := e.Group("/user")

	apiUser.POST("", cont.CreateUserController)
	apiUser.GET("/all", cont.GetUsersController)
	apiUser.GET("/:id", cont.GetUserByIDController)
	apiUser.PUT("/:id", cont.UpdateUserController)
	apiUser.DELETE("/:id", cont.DeleteUserController)
}
