package models

import (
	"errors"
	"time"

	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/askmuhammadamal/alta-store/lib/database/migrations"
	"github.com/askmuhammadamal/alta-store/middlewares"
	"github.com/labstack/echo/v4"
)

func CreateTransaction(c echo.Context) (interface{}, error) {

	user := migrations.User{}
	product := migrations.Product{}
	transaction := migrations.Transaction{}
	transactionDetail := migrations.TransactionDetail{}
	transactionRequest := migrations.TransactionRequest{}
	c.Bind(&transactionRequest)

	userId := middlewares.ExtractTokenUserId(c)
	if userId == 0 {
		return nil, errors.New("user is not found")
	}
	if err := database.DB.Find(&user, userId).Error; err != nil {
		return nil, errors.New("user is not found")
	}
	if err := database.DB.Find(&product, transactionRequest.ProductID).Error; err != nil {
		return nil, errors.New("product is not found")
	}
	if product.ID == 0 {
		return nil, errors.New("product is not found")
	}

	database.DB.Where("user_id = ? AND status = 'cart'", userId).First(&transaction)

	if transaction.UserID == 0 {
		transaction.User = user
		transaction.UserID = user.ID
		transaction.TransactionDate = time.Now()
		transaction.Shipping = 10000
		transaction.Status = "cart"
	}
	transaction.Total += product.Price * float64(transactionRequest.Quantity)

	if err := database.DB.Save(&transaction).Error; err != nil {
		return nil, err
	}

	database.DB.Where("transaction_id = ? AND product_id = ?", transaction.ID, product.ID).First(&transactionDetail)

	if transactionDetail.ID == 0 {
		transactionDetail.TransactionID = int(transaction.ID)
		transactionDetail.ProductID = int(product.ID)
	}
	transactionDetail.Quantity += int(transactionRequest.Quantity)

	if err := database.DB.Save(&transactionDetail).Error; err != nil {
		return nil, err
	}

	product.Stock -= int(transactionRequest.Quantity)
	if err := database.DB.Save(&product).Error; err != nil {
		return nil, err
	}

	return getDetailData([]migrations.Transaction{transaction})
}

func GetTransactions(c echo.Context) ([]migrations.TransactionResponse, error) {
	var transactions []migrations.Transaction
	user := migrations.User{}
	userId := middlewares.ExtractTokenUserId(c)
	if userId == 0 {
		return nil, errors.New("user is not found")
	}
	if err := database.DB.Find(&user, userId).Error; err != nil {
		return nil, errors.New("user is not found")
	}

	if err := database.DB.Where("user_id = ?", userId).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return getDetailData(transactions)
}

func GetTransaction(transactionId int) ([]migrations.TransactionResponse, error) {
	var transactions []migrations.Transaction

	if err := database.DB.Find(&transactions, transactionId).Error; err != nil {
		return nil, err
	}

	return getDetailData(transactions)
}

func EditTransaction(c echo.Context) (interface{}, error) {

	user := migrations.User{}
	transaction := migrations.Transaction{}
	updateStatusRequest := migrations.UpdateStatusRequest{}
	c.Bind(&updateStatusRequest)

	userId := middlewares.ExtractTokenUserId(c)
	if userId == 0 {
		return nil, errors.New("user is not found")
	}
	if err := database.DB.Find(&user, userId).Error; err != nil {
		return nil, errors.New("user is not found")
	}

	if updateStatusRequest.Status == "checkout" {
		database.DB.Where("user_id = ? AND status = 'cart'", userId).First(&transaction)

		if transaction.UserID == 0 {
			return nil, errors.New("cart is empty")
		}
		transaction.Status = "checkout"
	} else if updateStatusRequest.Status == "paid" {
		database.DB.Where("user_id = ? AND status = 'checkout'", userId).First(&transaction)

		if transaction.UserID == 0 {
			return nil, errors.New("checkout transaction is not found")
		}
		transaction.Status = "paid"
	}

	if err := database.DB.Save(&transaction).Error; err != nil {
		return nil, err
	}

	return getDetailData([]migrations.Transaction{transaction})
}

func DeleteTransaction(c echo.Context) error {

	user := migrations.User{}
	product := migrations.Product{}
	transaction := migrations.Transaction{}
	transactionDetail := migrations.TransactionDetail{}
	transactionRequest := migrations.TransactionRequest{}
	c.Bind(&transactionRequest)

	userId := middlewares.ExtractTokenUserId(c)
	if userId == 0 {
		return errors.New("user is not found")
	}
	if err := database.DB.Find(&user, userId).Error; err != nil {
		return errors.New("user is not found")
	}
	if err := database.DB.Find(&product, transactionRequest.ProductID).Error; err != nil {
		return errors.New("product is not found")
	}
	if product.ID == 0 {
		return errors.New("product is not found")
	}

	database.DB.Where("user_id = ? AND status = 'cart'", userId).First(&transaction)
	if transaction.ID == 0 {
		return errors.New("cart is empty")
	}

	database.DB.Where("transaction_id = ? AND product_id = ?", transaction.ID, product.ID).First(&transactionDetail)
	if transactionDetail.ID == 0 {
		return errors.New("product is not found in cart")
	}
	if transactionRequest.Quantity == 0 || transactionDetail.Quantity < int(transactionRequest.Quantity) {
		return errors.New("quantity is not valid")
	}

	// update quantity and stock
	transactionDetail.Quantity -= int(transactionRequest.Quantity)
	transaction.Total -= product.Price * float64(transactionRequest.Quantity)
	product.Stock += int(transactionRequest.Quantity)
	if err := database.DB.Save(&transactionDetail).Error; err != nil {
		return err
	}
	if err := database.DB.Save(&transaction).Error; err != nil {
		return err
	}
	if err := database.DB.Save(&product).Error; err != nil {
		return err
	}

	// delete transaction & transaction detail if necessary
	if transactionDetail.Quantity <= 0 {
		if err := database.DB.Delete(&transactionDetail, transactionDetail.ID).Error; err != nil {
			return err
		}
		tempTransDetail := migrations.TransactionDetail{}
		database.DB.Where("transaction_id = ?", transaction.ID).First(&tempTransDetail)
		if tempTransDetail.ID == 0 {
			if err := database.DB.Delete(&transaction, transaction.ID).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func getDetailData(transactions []migrations.Transaction) ([]migrations.TransactionResponse, error) {
	var transactionsResponse []migrations.TransactionResponse
	if len(transactions) > 0 {
		for i := range transactions {
			err := database.DB.Model(&migrations.User{}).Where("id = ?", transactions[i].UserID).Take(&transactions[i].User).Error
			if err != nil {
				return nil, err
			}
			var transDetail []migrations.TransactionDetail
			errTransDetail := database.DB.Model(&migrations.TransactionDetail{}).Where("transaction_id = ?", transactions[i].ID).Find(&transDetail).Error
			if errTransDetail != nil {
				return nil, errTransDetail
			}
			var transDetailInfo []migrations.TransactionDetailResponse
			var transRespon migrations.TransactionResponse
			for j := range transDetail {
				var temp migrations.TransactionDetailResponse
				temp.Quantity = int64(transDetail[j].Quantity)
				var product migrations.Product
				errTemp := database.DB.Model(&migrations.Product{}).Where("id = ?", transDetail[j].ProductID).First(&product).Error
				if errTemp != nil {
					return nil, errTemp
				}
				temp.Price = product.Price
				temp.Product = product.Name
				transDetailInfo = append(transDetailInfo, temp)
			}

			transRespon.ID = transactions[i].ID
			transRespon.CreatedAt = transactions[i].CreatedAt
			transRespon.UpdatedAt = transactions[i].UpdatedAt
			transRespon.DeletedAt = transactions[i].DeletedAt
			transRespon.UserID = transactions[i].UserID
			transRespon.User = transactions[i].User
			transRespon.TransactionDate = transactions[i].TransactionDate
			transRespon.Total = transactions[i].Total
			transRespon.Shipping = transactions[i].Shipping
			transRespon.Status = transactions[i].Status
			transRespon.Item = transDetailInfo
			transactionsResponse = append(transactionsResponse, transRespon)
		}
	}

	return transactionsResponse, nil
}
