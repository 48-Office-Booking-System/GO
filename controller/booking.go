package controller

import (
	"KOBA/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (sc *BookingServiceController) CreateBookingController(c echo.Context) error {
	booking := model.Booking{}
	c.Bind(&booking)

	_, err := sc.BookingServ.CreateBookingService(booking)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success creating booking",
		Data:    booking,
	}, "\t")
}

func (sc *BookingServiceController) GetBookingsController(c echo.Context) error {
	bookings, err := sc.BookingServ.GetBookingsService()

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting bookings",
		Data:    bookings,
	}, "\t")
}

func (sc *BookingServiceController) GetBookingByIDController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.ParseInt(id, 10, 64)
	booking, err := sc.BookingServ.GetBookingByIDService(int(intID))

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting booking",
		Data:    booking,
	}, "\t")
}
