package routes

import (
	"github.com/askmuhammadamal/alta-store/config"
	"github.com/askmuhammadamal/alta-store/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	// JWT Group
	jwtGroup := e.Group("")
	jwtGroup.Use(middleware.JWT([]byte(config.Env("SECRET_KEY"))))

	// User Routes
	e.POST("/login", controllers.LoginUserController)
	e.POST("/register", controllers.CreateUserController)
	e.POST("/users", controllers.CreateUserController)
	jwtGroup.GET("/users/:id", controllers.GetUserDetailController)
	jwtGroup.GET("/users", controllers.GetUserController)
	jwtGroup.PUT("/users/:id", controllers.UpdateUserController)
	jwtGroup.DELETE("/users/:id", controllers.DeleteUserController)

	// Category Routes
	e.GET("/categories", controllers.GetCategories)
	e.GET("/categories/:id", controllers.GetCategoryDetail)
	jwtGroup.POST("/categories", controllers.CreateCategory)
	jwtGroup.PUT("/categories/:id", controllers.UpdateCategory)
	jwtGroup.DELETE("/categories/:id", controllers.DeleteCategory)

	// Product Routes
	e.GET("/products", controllers.GetProducts)
	e.GET("/products/:id", controllers.GetProductDetail)
	jwtGroup.POST("/products", controllers.CreateProduct)
	jwtGroup.PUT("/products/:id", controllers.UpdateProduct)
	jwtGroup.DELETE("/products/:id", controllers.DeleteProduct)

	// Transaction Routes
	jwtGroup.POST("/transactions", controllers.CreateTransactionController)
	jwtGroup.GET("/transactions", controllers.GetTransactionController)
	jwtGroup.GET("/transactions/:id", controllers.GetTransactionDetailController)

	return e
}
