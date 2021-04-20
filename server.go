package main

import (
	"alta-store/config"
	"alta-store/controllers"
	"alta-store/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.Connection()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/users", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserController)

	// Start server
	e.Logger.Fatal(e.Start(config.Env("APP_PORT")))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, config.Env("APP_NAME"))
}
