package route

import (
	"KOBA/config"
	"KOBA/controller"
	"KOBA/database"
	"KOBA/repository"
	"KOBA/service"

	"github.com/labstack/echo/v4"
)

func RegisterBookingGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)

	repo := repository.NewBooking(db)

	svc := service.NewBookingService(repo, conf)

	cont := controller.BookingServiceController{
		BookingServ: svc,
	}

	apiBooking := e.Group("/booking")

	apiBooking.POST("", cont.CreateBookingController)
	apiBooking.GET("/all", cont.GetBookingsController)
	apiBooking.GET("/:id", cont.GetBookingByIDController)
	apiBooking.PUT("/:id", cont.UpdateBookingController)
	apiBooking.DELETE("/:id", cont.DeleteBookingController)
}
