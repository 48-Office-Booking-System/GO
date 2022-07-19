package controller

import (
	"KOBA/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//create review
func (sc *ReviewServiceController) CreateReviewController(c echo.Context) error {
	review := model.Review{}
	c.Bind(&review)

	_, err := sc.ReviewServ.CreateReviewService(review)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success creating review",
		Data:    review,
	}, "\t")
}

//get all review
func (sc *ReviewServiceController) GetReviewsController(c echo.Context) error {
	review := model.Review{}
	hidden, err := strconv.Atoi(c.QueryParam("hidden"))
	if err == nil {
		review.Hidden = hidden
	}

	office_id, err := strconv.Atoi(c.QueryParam("office_id"))
	if err == nil {
		review.OfficeID = office_id
	}

	reviews, err := sc.ReviewServ.GetReviewsService(review)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting reviews",
		Data:    reviews,
	}, "\t")
}

//get review by id
func (sc *ReviewServiceController) GetReviewByIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	review, err := sc.ReviewServ.GetReviewByIDService(int(intID))

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting review",
		Data:    review,
	}, "\t")
}

//update review
func (sc *ReviewServiceController) UpdateReviewController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	review := model.Review{}
	review.ID = int(intID)
	c.Bind(&review)

	err := sc.ReviewServ.UpdateReviewService(review, int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success updating review",
		Data:    review,
	}, "\t")
}

//delete review
func (sc *ReviewServiceController) DeleteReviewController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.ReviewServ.DeleteReviewService(int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success deleting review",
		Data:    nil,
	}, "\t")
}
