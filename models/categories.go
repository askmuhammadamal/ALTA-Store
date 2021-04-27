package models

import (
	"strconv"
	"time"

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

func UpdateCategory(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))

	category := migrations.Category{}
	c.Bind(&category)

	categoryDB := migrations.Category{}
	err := database.DB.Model(&category).Where("id = ?", id).Take(&categoryDB).UpdateColumns(
		map[string]interface{}{
			"name":        categoryDB.Name,
			"description": categoryDB.Description,
			"updated_at":  time.Now(),
		},
	).Error

	category.ID = categoryDB.ID
	category.CreatedAt = categoryDB.CreatedAt

	if err != nil {
		return nil, err
	}

	return category, nil
}

func DeleteCategory(id int) error {
	category := migrations.Category{}

	if err := database.DB.Delete(&category, id).Error; err != nil {
		return err
	}
	return nil
}
