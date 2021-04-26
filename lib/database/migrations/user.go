package migrations

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	CreatedAt   time.Time      `json:"createdAt" form:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt" form:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt" form:"deletedAt"`
	Email       string         `json:"email" form:"email" gorm:"size:100;index;unique_index;not null"`
	Password    string         `json:"password" form:"password" gorm:"not null"`
	FullName    string         `json:"fullName" form:"fullName" gorm:"not null"`
	PhoneNumber string         `json:"phoneNumber" form:"phoneNumber" gorm:"size:20;not null"`
	Gender      string         `json:"gender" form:"gender" gorm:"type:enum('male', 'female');not null"`
	DateOfBirth time.Time      `json:"dateOfBirth" form:"dateOfBirth" gorm:"not null"`
	District    string         `json:"district" form:"district" gorm:"size:200;not null"`
	SubDistrict string         `json:"subDistrict" form:"subDistrict" gorm:"size:200;not null"`
	Address     string         `json:"address" form:"address" gorm:"not null"`
	Role        string         `json:"role" form:"role" gorm:"type:enum('user', 'admin');default:user;not null"`
}

type UserRespon struct {
	ID          uint           `json:"id" form:"id"`
	CreatedAt   time.Time      `json:"createdAt" form:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt" form:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" form:"deletedAt"`
	Email       string         `json:"email" form:"email"`
	Password    string         `json:"password" form:"password"`
	FullName    string         `json:"fullName" form:"fullName"`
	PhoneNumber string         `json:"phoneNumber" form:"phoneNumber"`
	Gender      string         `json:"gender" form:"gender"`
	DateOfBirth time.Time      `json:"dateOfBirth" form:"dateOfBirth"`
	District    string         `json:"district" form:"district"`
	SubDistrict string         `json:"subDistrict" form:"subDistrict"`
	Address     string         `json:"address" form:"address"`
	Role        string         `json:"role" form:"role"`
	Token       string         `json:"token" form:"token"`
}
