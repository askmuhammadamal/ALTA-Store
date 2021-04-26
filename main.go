package main

import (
	"alta-store/config"
	"alta-store/lib/database"
	"alta-store/routes"
)

func main() {
	database.Connection()
	e := routes.New()
	e.Logger.Fatal(e.Start(config.Env("APP_PORT")))
}
