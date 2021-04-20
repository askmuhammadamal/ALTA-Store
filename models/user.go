package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	FullName    string `json:"full_name" form:"full_name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	//Gender      string    `json:"gender" form:"gender"`
	//DateOfBirth time.Time `json:"date_of_birth" form:"date_of_birth"`
	District    string `json:"district" form:"district"`
	SubDistrict string `json:"sub_district" form:"sub_district"`
	Address     string `json:"address" form:"address"`
	//Role        string    `json:"role" form:"role"`
}
