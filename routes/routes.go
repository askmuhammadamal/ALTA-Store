package routes

import (
	"github.com/askmuhammadamal/alta-store/config"
	"github.com/askmuhammadamal/alta-store/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// JWT Group
	jwtGroup := e.Group("")
	jwtGroup.Use(middleware.JWT([]byte(config.Env("SECRET_KEY"))))

	// User Routes
	jwtGroup.GET("/users/:id", controllers.GetUserDetailController)
	jwtGroup.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUserController)
	jwtGroup.PUT("/users/:id", controllers.UpdateUserController)
	jwtGroup.DELETE("/users/:id", controllers.DeleteUserController)

	// Category Routes
	e.GET("/categories", controllers.GetCategories)
	e.GET("/categories/:id", controllers.GetCategoryDetail)
	e.POST("/categories", controllers.CreateCategory)
	e.PUT("/categories/:id", controllers.UpdateCategory)
	e.DELETE("/categories/:id", controllers.DeleteCategory)
	return e
}
