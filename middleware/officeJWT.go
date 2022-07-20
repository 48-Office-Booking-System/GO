package middleware

import (
	"KOBA/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.InitConfiguration().JWT_KEY))
}

// apakan pakai log?
// func LogOffice(e *echo.Echo) {
// 	e.Use(middleware.Logger())
// }