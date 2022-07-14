package middleware

import (
	"KOBA/config"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.InitConfiguration().JWT_KEY))
}
