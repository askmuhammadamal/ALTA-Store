package migrations

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	IdTransaction int `json:"id_transaction" form:"id_transaction" gorm:"type:int(10);index;not null"`
	IdProduct     int `json:"id_product" form:"id_product" gorm:"type:int(10);index;not null"`
	Quantity      int `json:"quantity" form:"quantity" gorm:"type:int(10);not null"`
}
