package router

import (
	"go-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.Signup)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	return e
}
