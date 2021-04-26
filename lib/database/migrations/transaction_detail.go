package migrations

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	TransactionID int         `json:"transactionId" form:"transactionId" gorm:"type:int(10);index;not null"`
	Transaction   Transaction `json:"transaction" form:"transaction"`
	ProductID     int         `json:"productId" form:"productId" gorm:"type:int(10);index;not null"`
	Product       Product     `json:"product" form:"product"`
	Quantity      int         `json:"quantity" form:"quantity" gorm:"type:int(10);not null"`
}
