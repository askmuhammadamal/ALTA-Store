package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/askmuhammadamal/alta-store/lib/database/migrations"
	"github.com/askmuhammadamal/alta-store/models"
	"github.com/labstack/echo/v4"
)

func GetCategories(c echo.Context) error {
	categories, err := models.GetCategories()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(categories) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get categories",
			"status":  "success",
			"data":    categories,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "categories not found",
		"status":  "fail",
	})
}

func GetCategoryDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	category, err := models.GetCategory((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(category) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get category detail",
			"status":  "success",
			"data":    category,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "category not found",
		"status":  "fail",
	})
}

func CreateCategory(c echo.Context) error {
	category, e := models.CreateCategory(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create new category",
		"status":  "success",
		"data":    category,
	})
}

func UpdateCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	category := migrations.Category{}
	c.Bind(&category)

	err := database.DB.Model(&category).Where("id = ?", id).Take(&migrations.Category{}).UpdateColumns(
		map[string]interface{}{
			"name":        category.Name,
			"description": category.Description,
			"updated_at":  time.Now(),
		},
	).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
		"data":   category,
	})
}

func DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	category := migrations.Category{}
	c.Bind(&category)

	if err := database.DB.Delete(&category, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
	})
}
