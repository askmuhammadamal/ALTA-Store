package migrations

import (
	"time"

	"gorm.io/gorm"
)

type TransactionDetail struct {
	ID            uint           `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	CreatedAt     time.Time      `json:"createdAt" form:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt" form:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt" form:"deletedAt"`
	TransactionID int            `json:"transactionId" form:"transactionId" gorm:"type:int(10);index;not null"`
	Transaction   Transaction    `json:"transaction" form:"transaction"`
	ProductID     int            `json:"productId" form:"productId" gorm:"type:int(10);index;not null"`
	Product       Product        `json:"product" form:"product"`
	Quantity      int            `json:"quantity" form:"quantity" gorm:"type:int(10);not null"`
}
