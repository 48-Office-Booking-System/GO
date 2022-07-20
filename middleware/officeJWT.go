package middleware

import (
	"KOBA/config"

	"github.com/dgrijalva/jwt-go"
	echo "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/middleware"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.InitConfiguration().JWT_KEY))
}

// apakan pakai log?
// func LogOffice(e *echo.Echo) {
// 	e.Use(middleware.Logger())
// }