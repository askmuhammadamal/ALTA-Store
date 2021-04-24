package migrations

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name 		string	`json:"name" form:"name" gorm:"size:255;not null"`
	Description string	`json:"description" form:"description"`
}
