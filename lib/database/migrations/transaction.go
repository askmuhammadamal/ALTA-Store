package migrations

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID          uint      `json:"userId" form:"userId" gorm:"type:int(10);index;not null"`
	User            User      `json:"user" form:"user"`
	TransactionDate time.Time `json:"transactionDate" form:"transactionDate" gorm:"not null"`
	Total           int64     `json:"total" form:"total" gorm:"type:double;not null"`
	Shipping        int32     `json:"shipping" form:"shipping" gorm:"type:double;not null"`
	Status          string    `json:"status" form:"status" gorm:"type:enum('cart', 'checkout', 'paid', 'confirmed', 'sent', 'received');default:cart;not null"`
}
