package database

import (
	"fmt"

	"github.com/askmuhammadamal/alta-store/config"
	"github.com/askmuhammadamal/alta-store/lib/database/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Env("DB_USERNAME"), config.Env("DB_PASSWORD"), config.Env("DB_HOST"), config.Env("DB_PORT"), config.Env("DB_DATABASE"))

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&migrations.User{}, &migrations.Product{}, &migrations.Category{}, &migrations.Transaction{}, &migrations.TransactionDetail{})
}
