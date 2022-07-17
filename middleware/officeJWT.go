package middleware

import {
	echo "github.com/labstack/echo"
	"github.com/labstack/middleware"
}

func JWTMiddlewareOffice() echo.MiddlewareFunc{
	returnmiddleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SighningKey: []byte("rahasia"),
	}) 
} 


func (handler *OfficeHandler) Login(c echo.Context)