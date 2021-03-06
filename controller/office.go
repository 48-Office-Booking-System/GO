package controller

import (
	"KOBA/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// menambahkan gedung
func (sc *OfficeServiceController) CreateOfficeController(c echo.Context) error {
	office := model.Office{}
	c.Bind(&office)

	_, err := sc.OfficeServ.CreateOfficeService(office)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success creating office",
		Data:    office,
	}, "\t")
}

// menampilkan seluruh gedung
func (sc *OfficeServiceController) GetAllOfficeController(c echo.Context) error {
	office, err := sc.OfficeServ.GetAllOfficeService()

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting office",
		Data:    office,
	}, "\t")
}

// menampilkan gedung berdasarkan id
func (sc *OfficeServiceController) GetOfficeController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	office, err := sc.OfficeServ.GetOfficeService(int(intID))

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting office",
		Data:    office,
	}, "\t")
}

//update gedung
func (sc *OfficeServiceController) UpdateOfficeController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	office := model.Office{}
	office.ID = int(intID)
	c.Bind(&office)

	err := sc.OfficeServ.UpdateOfficeService(office, int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success updating office",
		Data:    office,
	}, "\t")
}

// hapus gedung
func (sc *OfficeServiceController) DeleteOfficeController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.OfficeServ.DeleteOfficeService(int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success deleting office",
		Data:    nil,
	}, "\t")
}
