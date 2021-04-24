package migrations

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	IdUser 			string		`json:"id_user" form:"id_user" gorm:"type:int(10);index;not null"`
	TransactionDate time.Time	`json:"transaction_date" form:"transaction_date" gorm:"not null"`
	Total 			int64		`json:"total" form:"total" gorm:"type:double;not null"`
	Shipping 		int32		`json:"shipping" form:"shipping" gorm:"type:double;not null"`
	Status 			string		`json:"status" form:"status" gorm:"type:enum('cart', 'checkout', 'paid', 'confirmed', 'sent', 'received');default:cart;not null"`
}
