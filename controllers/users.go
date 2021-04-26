package controllers

import (
	"alta-store/lib/database"
	"alta-store/lib/database/migrations"
	"alta-store/models"
	"net/http"
	"strconv"
	"time"

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

	token, e := models.LoginUsers(&user, user.Password)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
		"token":  token,
	})
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// binding data
	user := migrations.User{}
	c.Bind(&user)

	hashPassword, errHash := models.Hash(user.Password)
	if errHash != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errHash.Error())
	}
	user.Password = string(hashPassword)

	err := database.DB.Model(&user).Where("id = ?", id).Take(&migrations.User{}).UpdateColumns(
		map[string]interface{}{
			"full_name":     user.FullName,
			"phone_number":  user.PhoneNumber,
			"email":         user.Email,
			"password":      user.Password,
			"gender":        user.Gender,
			"date_of_birth": user.DateOfBirth,
			"district":      user.District,
			"sub_district":  user.SubDistrict,
			"address":       user.Address,
			"updated_at":    time.Now(),
		},
	).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
		"user":   user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// binding data
	user := migrations.User{}
	c.Bind(&user)

	if err := database.DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
	})
}
