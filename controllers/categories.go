package controllers

import (
	"net/http"
	"strconv"

	"github.com/askmuhammadamal/alta-store/models"
	"github.com/labstack/echo/v4"
)

func GetCategoriesContoller(c echo.Context) error {
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

func GetCategoryDetailContoller(c echo.Context) error {
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

func CreateCategoryContoller(c echo.Context) error {
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

func UpdateCategoryContoller(c echo.Context) error {
	category, err := models.UpdateCategory(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
		"data":   category,
	})
}

func DeleteCategoryContoller(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := models.DeleteCategory((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
	})
}
