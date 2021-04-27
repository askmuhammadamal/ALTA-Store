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
	e.GET("/categories", controllers.GetCategoriesContoller)
	e.GET("/categories/:id", controllers.GetCategoryDetailContoller)
	jwtGroup.POST("/categories", controllers.CreateCategoryContoller)
	jwtGroup.PUT("/categories/:id", controllers.UpdateCategoryContoller)
	jwtGroup.DELETE("/categories/:id", controllers.DeleteCategoryContoller)

	// Product Routes
	e.GET("/products", controllers.GetProductsContoller)
	e.GET("/products/:id", controllers.GetProductDetailContoller)
	jwtGroup.POST("/products", controllers.CreateProductContoller)
	jwtGroup.PUT("/products/:id", controllers.UpdateProductContoller)
	jwtGroup.DELETE("/products/:id", controllers.DeleteProductContoller)

	// Transaction Routes
	jwtGroup.POST("/transactions", controllers.CreateTransactionController)
	jwtGroup.GET("/transactions", controllers.GetTransactionController)
	jwtGroup.GET("/transactions/:id", controllers.GetTransactionDetailController)

	return e
}
