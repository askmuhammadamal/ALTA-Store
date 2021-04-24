package main

import (
	"alta-store/config"
	"alta-store/lib/database"
	"github.com/labstack/echo/v4"
)

func main() {
	database.Connection()
	e := echo.New()
	e.Logger.Fatal(e.Start(config.Env("APP_PORT")))
}