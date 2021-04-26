package migrations

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string    `json:"email" form:"email" gorm:"size:100;index;unique_index;not null"`
	Password    string    `json:"password" form:"password" gorm:"not null"`
	FullName    string    `json:"full_name" form:"full_name" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" form:"phone_number" gorm:"size:20;not null"`
	Gender      string    `json:"gender" form:"gender" gorm:"type:enum('male', 'female');not null"`
	DateOfBirth time.Time `json:"date_of_birth" form:"date_of_birth" gorm:"not null"`
	District    string    `json:"district" form:"district" gorm:"size:200;not null"`
	SubDistrict string    `json:"sub_district" form:"sub_district" gorm:"size:200;not null"`
	Address     string    `json:"address" form:"address" gorm:"not null"`
	Role        string    `json:"role" form:"role" gorm:"type:enum('user', 'admin');default:user;not null"`
}
