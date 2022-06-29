package controller

import (
	"KOBA/model"
	"net/http"
	"strconv"

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

func (sc *UserServiceController) GetUsersController(c echo.Context) error {
	users, err := sc.UserServ.GetUsersService()

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting users",
		Data:    users,
	}, "\t")
}

func (sc *UserServiceController) GetUserByIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	user, err := sc.UserServ.GetUserByIDService(int(intID))

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting user",
		Data:    user,
	}, "\t")
}

func (sc *UserServiceController) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	user := model.User{}
	user.ID = int(intID)
	c.Bind(&user)

	err := sc.UserServ.UpdateUserService(user, int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success updating user",
		Data:    user,
	}, "\t")
}

func (sc *UserServiceController) DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.UserServ.DeleteUserService(int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success deleting user",
		Data:    nil,
	}, "\t")
}
