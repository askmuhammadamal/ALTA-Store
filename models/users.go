package models

import (
	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/askmuhammadamal/alta-store/lib/database/migrations"
	"github.com/askmuhammadamal/alta-store/middlewares"

	"github.com/labstack/echo/v4"
)

func GetUsers() ([]migrations.User, error) {
	var users []migrations.User

	if e := database.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUser(userId int) ([]migrations.User, error) {
	var users []migrations.User

	if err := database.DB.Find(&users, userId).Error; err != nil {
		return nil, err
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

func LoginUsers(user *migrations.User) (interface{}, error) {
	var err error
	token := migrations.Token{}
	if err = database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	token.Data, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	return token, nil
}
