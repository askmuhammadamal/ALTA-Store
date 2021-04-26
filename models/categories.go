package models

import (
	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/askmuhammadamal/alta-store/lib/database/migrations"
	"github.com/labstack/echo/v4"
)

func GetCategories() ([]migrations.Category, error) {
	var categories []migrations.Category

	if e := database.DB.Find(&categories).Error; e != nil {
		return nil, e
	}

	return categories, nil
}

func GetCategory(id int) ([]migrations.Category, error) {
	var category []migrations.Category

	if e := database.DB.Find(&category, id).Error; e != nil {
		return nil, e
	}

	return category, nil
}

func CreateCategory(c echo.Context) (interface{}, error) {
	category := migrations.Category{}
	c.Bind(&category)

	if e := database.DB.Save(&category).Error; e != nil {
		return nil, e
	}

	return category, nil
}
