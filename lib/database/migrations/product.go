package migrations

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string     `json:"name" form:"name" gorm:"size:255;not null"`
	Description string     `json:"description" form:"description" gorm:"not null"`
	Stock       int        `json:"stock" from:"stock" gorm:"size:10;not null"`
	Price       int64      `json:"price" form:"price" gorm:"type:double;not null"`
	Category    []Category `json:"category" form:"category" gorm:"type:int(10);index;not null"`
}
