package controllers

import (
	"net/http"

	"github.com/askmuhammadamal/alta-store/models"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	users, e := models.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
		"data":   users,
	})
}

func CreateUser(c echo.Context) error {
	user, e := models.CreateUsers(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success create new user",
		"data":   user,
	})
}
