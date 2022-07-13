package route

import (
	"KOBA/config"
	"KOBA/controller"
	"KOBA/database"
	"KOBA/repository"
	"KOBA/service"

	"github.com/labstack/echo/v4"
)

func RegisterOfficeGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)

	repo := repository.NewOffice(db)

	svc := service.NewOfficeService(repo, conf)

	cont := controller.OfficeServiceController{
		OfficeServ: svc,
	}

	apiOffice := e.Group("/office")

	apiOffice.POST("", cont.CreateOfficeController)
	apiOffice.GET("/all", cont.GetAllOfficeController)
	apiOffice.GET("/:id", cont.GetOfficeController)
	apiOffice.PUT("/:id", cont.UpdateOfficeController)
	apiOffice.DELETE("/:id", cont.DeleteOfficeController)
}
