package migrations

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	CreatedAt   time.Time      `json:"createdAt" form:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt" form:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt" form:"deletedAt"`
	Name        string         `json:"name" form:"name" gorm:"size:255;not null"`
	Description string         `json:"description" form:"description"`
}
