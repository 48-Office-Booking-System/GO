package route

import (
	"KOBA/config"
	"KOBA/controller"
	"KOBA/database"
	"KOBA/repository"
	"KOBA/service"

	"github.com/labstack/echo/v4"
)

func RegisterReviewGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)

	repo := repository.NewReview(db)

	svc := service.NewReviewService(repo, conf)

	cont := controller.ReviewServiceController{
		ReviewServ: svc,
	}

	apiReview := e.Group("/review")

	apiReview.POST("", cont.CreateReviewController)
	apiReview.GET("/all", cont.GetReviewsController)
	apiReview.GET("/:id", cont.GetReviewByIDController)
	apiReview.PUT("/:id", cont.UpdateReviewController)
	apiReview.DELETE("/:id", cont.DeleteReviewController)
}
