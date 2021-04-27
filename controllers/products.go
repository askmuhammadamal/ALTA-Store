package controllers

import (
	"net/http"
	"strconv"

	"github.com/askmuhammadamal/alta-store/models"
	"github.com/labstack/echo/v4"
)

func GetProductsContoller(c echo.Context) error {
	products, err := models.GetProducts()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(products) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get user",
			"status":  "success",
			"data":    products,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "list products not found",
		"status":  "fail",
	})

}

func GetProductDetailContoller(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := models.GetProductDetail((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(product) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get product detail",
			"status":  "success",
			"data":    product,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "data product not found",
		"status":  "fail",
	})
}

func CreateProductContoller(c echo.Context) error {
	product, e := models.CreateProduct(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create new product",
		"status":  "success",
		"data":    product,
	})
}

func UpdateProductContoller(c echo.Context) error {
	product, err := models.UpdateProduct(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
		"data":   product,
	})
}

func DeleteProductContoller(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := models.DeleteProduct((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "success",
	})
}
