package office

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

func NewOfficeHandler(e *echo.Echo, ) {// code apa aku ini )
	e.POST("/offices", middleware.JWTMiddleware())
	e.POST("/login", handler.Login)
}

func (handler *OfficeHandler) Check(c echo.Context) error{
	var request JWTForOffice

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if requestt.Usrname != "admin" && request.Password != "admin"{
		return c.JSON(http.StatusUnauthorized, "username atau password salah")
	}

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour *24).Unix()

	token := jwtNewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := token.SignedString([]byte("rahasia"))
	if err != nill{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, jwt)
}