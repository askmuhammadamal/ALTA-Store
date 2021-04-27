package controllers

import (
	"net/http"
	"strconv"

	"github.com/askmuhammadamal/alta-store/models"
	"github.com/labstack/echo/v4"
)

func CreateTransactionController(c echo.Context) error {
	transaction, e := models.CreateTransaction(c)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create new user",
		"status":  "success",
		"data":    transaction,
	})
}

func GetTransactionController(c echo.Context) error {
	transactions, err := models.GetTransactions(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(transactions) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get transactions",
			"status":  "success",
			"data":    transactions,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "transactions not found",
		"status":  "fail",
	})

}

func GetTransactionDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	transactions, err := models.GetTransaction((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(transactions) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get transaction detail",
			"status":  "success",
			"data":    transactions,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "transaction not found",
		"status":  "fail",
	})
}
