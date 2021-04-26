package models

import (
	"alta-store/lib/database"
	"alta-store/lib/database/migrations"
	"alta-store/middlewares"

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
	userRespon := migrations.UserRespon{}
	var err error
	if err = database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	userRespon.ID = user.ID
	userRespon.CreatedAt = user.CreatedAt
	userRespon.UpdatedAt = user.UpdatedAt
	userRespon.DeletedAt = user.DeletedAt
	userRespon.Email = user.Email
	userRespon.Password = user.Password
	userRespon.FullName = user.FullName
	userRespon.PhoneNumber = user.PhoneNumber
	userRespon.Gender = user.Gender
	userRespon.DateOfBirth = user.DateOfBirth
	userRespon.District = user.District
	userRespon.SubDistrict = user.SubDistrict
	userRespon.Address = user.Address
	userRespon.Role = user.Role
	userRespon.Token, err = middlewares.CreateToken(int(userRespon.ID))
	if err != nil {
		return nil, err
	}

	return userRespon, nil
}
