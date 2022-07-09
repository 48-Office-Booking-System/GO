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

	id, err := sc.UserServ.CreateUserService(user)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	user = model.User{}
	user, err = sc.UserServ.GetUserByIDService(id)

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: "success creating user but user not found" + err.Error(),
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

	role := c.QueryParam("role_id")
	role_id, err := strconv.Atoi(role)
	if err == nil {
		for i := range users {
			if users[i].RoleID != role_id {
				users = append(users[:i], users[i+1:]...)
			}
		}
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting users " + role,
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

	user = model.User{}
	user, err = sc.UserServ.GetUserByIDService(int(intID))

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: "success updating user but user not found" + err.Error(),
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

func (sc *UserServiceController) GetRolesController(c echo.Context) error {
	roles, err := sc.UserServ.GetRolesService()

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting roles",
		Data:    roles,
	}, "\t")
}

func (sc *UserServiceController) DeleteRoleController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.UserServ.DeleteRoleService(int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success deleting role",
		Data:    nil,
	}, "\t")
}
