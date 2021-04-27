package controllers

import (
	"net/http"
	"strconv"

	"github.com/askmuhammadamal/alta-store/lib/database/migrations"
	"github.com/askmuhammadamal/alta-store/models"
	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := migrations.User{}
	c.Bind(&user)

	token, e := models.LoginUsers(&user, user.Password)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
		"data":   token,
	})
}

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
			"data":    users,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "users not found",
		"status":  "fail",
	})

}

func GetUserDetailController(c echo.Context) error {
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
			"data":    users,
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
		"data":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	user, err := models.EditUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
		"data":   user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := models.DeleteUser((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
	})
}
