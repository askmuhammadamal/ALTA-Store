package migrations

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	CreatedAt   time.Time      `json:"createdAt" form:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt" form:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt" form:"deletedAt"`
	Name        string         `json:"name" form:"name" gorm:"size:255;not null"`
	Description string         `json:"description" form:"description" gorm:"not null"`
	Stock       int            `json:"stock" form:"stock" gorm:"type:int(10);not null"`
	Price       float64        `json:"price" form:"price" gorm:"type:double;not null"`
	CategoryID  uint           `json:"categoryId" form:"categoryId" gorm:"type:int(10);index;not null"`
	Category    Category       `json:"category" form:"category"`
}
