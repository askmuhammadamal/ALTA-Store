package models

import (
	"strconv"
	"time"

	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/askmuhammadamal/alta-store/lib/database/migrations"
	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) ([]migrations.Product, error) {
	var products []migrations.Product
	name := c.QueryParam("name")
	categoryName := c.QueryParam("categoryName")
	categoryId := c.QueryParam("categoryId")

	if name != "" {
		if e := database.DB.Model(&migrations.Product{}).Where("name LIKE ?", "%"+name+"%").Find(&products).Error; e != nil {
			return nil, e
		}
	} else if categoryName != "" {
		var category migrations.Category
		if e := database.DB.Model(&migrations.Category{}).Where("name LIKE ?", "%"+categoryName+"%").First(&category).Error; e != nil {
			return nil, e
		}
		if category.ID != 0 {
			if e := database.DB.Model(&migrations.Product{}).Where("category_id = ?", category.ID).Find(&products).Error; e != nil {
				return nil, e
			}
		}
	} else if categoryId != "" {
		var category migrations.Category
		if e := database.DB.Model(&migrations.Category{}).Where("id = ?", categoryId).First(&category).Error; e != nil {
			return nil, e
		}
		if category.ID != 0 {
			if e := database.DB.Model(&migrations.Product{}).Where("category_id = ?", category.ID).Find(&products).Error; e != nil {
				return nil, e
			}
		}
	} else {
		if e := database.DB.Find(&products).Error; e != nil {
			return nil, e
		}
	}

	return getCategoryData(products)
}

func GetProductDetail(id int) ([]migrations.Product, error) {
	var product []migrations.Product

	if err := database.DB.Find(&product, id).Error; err != nil {
		return nil, err
	}
	return getCategoryData(product)
}

func CreateProduct(c echo.Context) (interface{}, error) {

	product := migrations.Product{}
	c.Bind(&product)

	if e := database.DB.Save(&product).Error; e != nil {
		return nil, e
	}

	return getCategoryData([]migrations.Product{product})
}

func UpdateProduct(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))

	product := migrations.Product{}
	c.Bind(&product)

	productDB := migrations.Product{}
	err := database.DB.Model(&product).Where("id = ?", id).Take(&productDB).UpdateColumns(
		map[string]interface{}{
			"name":        product.Name,
			"description": product.Description,
			"stock":       product.Stock,
			"price":       product.Price,
			"category_id": product.CategoryID,
			"updated_at":  time.Now(),
		},
	).Error

	product.ID = productDB.ID
	product.CreatedAt = productDB.CreatedAt

	if err != nil {
		return nil, err
	}

	return getCategoryData([]migrations.Product{product})
}

func DeleteProduct(id int) error {
	product := migrations.Product{}

	if err := database.DB.Delete(&product, id).Error; err != nil {
		return err
	}
	return nil
}

func getCategoryData(products []migrations.Product) ([]migrations.Product, error) {
	if len(products) > 0 {
		for i := range products {
			err := database.DB.Model(&migrations.Category{}).Where("id = ?", products[i].CategoryID).Take(&products[i].Category).Error
			if err != nil {
				return nil, err
			}
		}
	}

	return products, nil
}
