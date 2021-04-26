package controllers

import (
	"alta-store/lib/database/migrations"
	"alta-store/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	users, err := models.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(users) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get user",
			"status":  "success",
			"users":   users,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "users not found",
		"status":  "fail",
	})

}

func GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := models.GetUser((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(users) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get user detail",
			"status":  "success",
			"users":   users,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "user not found",
		"status":  "fail",
	})
}

func CreateUserController(c echo.Context) error {
	user, e := models.CreateUsers(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create new user",
		"status":  "success",
		"user":    user,
	})
}

func LoginUsersController(c echo.Context) error {
	user := migrations.User{}
	c.Bind(&user)

	userRespon, e := models.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success login",
		"status":  "success",
		"users":   userRespon,
	})
}
