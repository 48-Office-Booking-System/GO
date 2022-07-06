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

	id, err := sc.OfficeServ.CreateOfficeService(office)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	office = model.Office{}
	office, err = sc.OfficeServ.GetOfficeService(id)

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: "success creating office but office not found" + err.Error(),
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

	office = model.Office{}
	office, err = sc.OfficeServ.GetOfficeService(int(intID))

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: "success updating office but office not found" + err.Error(),
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

func (sc *OfficeServiceController) GetTypesController(c echo.Context) error {
	types, err := sc.OfficeServ.GetTypesService()

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting types",
		Data:    types,
	}, "\t")
}

func (sc *OfficeServiceController) DeleteTypeController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.OfficeServ.DeleteTypeService(int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success deleting type",
		Data:    nil,
	}, "\t")
}

func (sc *OfficeServiceController) GetFacilitationsController(c echo.Context) error {
	facilitations, err := sc.OfficeServ.GetFacilitationsService()

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting facilitations",
		Data:    facilitations,
	}, "\t")
}

func (sc *OfficeServiceController) DeleteFacilitationController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.OfficeServ.DeleteFacilitationService(int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success deleting facilitation",
		Data:    nil,
	}, "\t")
}

func (sc *OfficeServiceController) GetTagsController(c echo.Context) error {
	tags, err := sc.OfficeServ.GetTagsService()

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
		Data:    tags,
	}, "\t")
}

func (sc *OfficeServiceController) DeleteTagController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.OfficeServ.DeleteTagService(int(intID))
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
