package models

import (
	"strconv"
	"time"

	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/askmuhammadamal/alta-store/lib/database/migrations"
	"github.com/labstack/echo/v4"
)

func GetProducts() ([]migrations.Product, error) {
	var products []migrations.Product

	if e := database.DB.Find(&products).Error; e != nil {
		return nil, e
	}
	return products, nil
}

func GetProductDetail(id int) ([]migrations.Product, error) {
	var product []migrations.Product

	if err := database.DB.Find(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func CreateProduct(c echo.Context) (interface{}, error) {

	product := migrations.Product{}
	c.Bind(&product)

	if e := database.DB.Save(&product).Error; e != nil {
		return nil, e
	}

	return product, nil
}

func UpdateProduct(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))

	product := migrations.Product{}
	c.Bind(&product)

	productDB := migrations.Product{}
	err := database.DB.Model(&product).Where("id = ?", id).Take(&productDB).UpdateColumns(
		map[string]interface{}{
			"name":        productDB.Name,
			"description": productDB.Description,
			"stock":       productDB.Stock,
			"price":       productDB.Price,
			"category":    productDB.Category,
			"updated_at":  time.Now(),
		},
	).Error

	product.ID = productDB.ID
	product.CreatedAt = productDB.CreatedAt

	if err != nil {
		return nil, err
	}

	return product, nil
}

func DeleteProduct(id int) error {
	product := migrations.Product{}

	if err := database.DB.Delete(&product, id).Error; err != nil {
		return err
	}
	return nil
}
