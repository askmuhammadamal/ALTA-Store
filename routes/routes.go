package routes

import (
	"net/http"

	"github.com/askmuhammadamal/alta-store/controllers"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello this is echo")
	})

	e.GET("/users", controllers.GetUser)
	e.POST("/users", controllers.CreateUser)

	return e
}
