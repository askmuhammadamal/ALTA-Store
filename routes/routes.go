package routes

import (
	"alta-store/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUserController)
	e.GET("/users/:id", controllers.GetUserDetailControllers)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUsersController)
	return e
}
