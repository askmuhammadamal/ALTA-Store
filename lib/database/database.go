package database

import (
	"alta-store/config"
	"alta-store/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := config.Env("DB_USERNAME") + ":" + config.Env("DB_PASSWORD") + "@tcp(" + config.Env("DB_HOST") + ":" + config.Env("DB_PORT") + ")/" + config.Env("DB_DATABASE") + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	//defer db.Close()
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
