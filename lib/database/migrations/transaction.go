package migrations

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID              uint           `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	CreatedAt       time.Time      `json:"createdAt" form:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt" form:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deletedAt" form:"deletedAt"`
	UserID          uint           `json:"userId" form:"userId" gorm:"type:int(10);index;not null"`
	User            User           `json:"user" form:"user"`
	TransactionDate time.Time      `json:"transactionDate" form:"transactionDate" gorm:"not null"`
	Total           float64        `json:"total" form:"total" gorm:"type:double;not null"`
	Shipping        float64        `json:"shipping" form:"shipping" gorm:"type:double;not null"`
	Status          string         `json:"status" form:"status" gorm:"type:enum('cart', 'checkout', 'paid', 'confirmed', 'sent', 'received');default:cart;not null"`
}

type TransactionRequest struct {
	ProductID uint  `json:"product" form:"product"`
	Quantity  int64 `json:"quantity" form:"quantity"`
}

type TransactionResponse struct {
	ID              uint                        `json:"id" form:"id"`
	CreatedAt       time.Time                   `json:"createdAt" form:"createdAt"`
	UpdatedAt       time.Time                   `json:"updatedAt" form:"updatedAt"`
	DeletedAt       gorm.DeletedAt              `json:"deletedAt" form:"deletedAt"`
	UserID          uint                        `json:"userId" form:"userId"`
	User            User                        `json:"user" form:"user"`
	TransactionDate time.Time                   `json:"transactionDate" form:"transactionDate"`
	Total           float64                     `json:"total" form:"total"`
	Shipping        float64                     `json:"shipping" form:"shipping"`
	Status          string                      `json:"status" form:"status"`
	Item            []TransactionDetailResponse `json:"item" form:"item"`
}

type TransactionDetailResponse struct {
	Product  string  `json:"product" form:"product"`
	Quantity int64   `json:"quantity" form:"quantity"`
	Price    float64 `json:"price" form:"price"`
}
