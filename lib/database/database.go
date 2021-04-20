package database

import (
	"alta-store/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() {
	dsn := config.Env("DB_USERNAME") + ":" + config.Env("DB_PASSWORD") + "@tcp(" + config.Env("DB_HOST") + ":" + config.Env("DB_PORT") + ")/" + config.Env("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	defer db.Close()
}