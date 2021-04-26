package models

import (
	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/askmuhammadamal/alta-store/lib/database/migrations"
	"github.com/labstack/echo/v4"
)

func GetUsers() (interface{}, error) {
	var users []migrations.User

	if e := database.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func CreateUsers(c echo.Context) (interface{}, error) {

	user := migrations.User{}
	c.Bind(&user)

	if e := database.DB.Save(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}
