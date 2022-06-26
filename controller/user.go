package controller

import (
	"KOBA/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (sc *UserServiceController) CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	_, err := sc.UserServ.CreateUserService(user)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "Success Creating User",
		Data:    user,
	}, "\t")
}
