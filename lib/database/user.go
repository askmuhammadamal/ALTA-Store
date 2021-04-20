package database

import (
	"alta-store/models"

	"github.com/labstack/echo/v4"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func CreateUsers(c echo.Context) (interface{}, error) {

	user := models.User{}
	c.Bind(&user)

	if e := DB.Save(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}
