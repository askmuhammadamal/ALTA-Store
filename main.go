package main

import (
	"github.com/askmuhammadamal/alta-store/config"
	"github.com/askmuhammadamal/alta-store/lib/database"
	"github.com/askmuhammadamal/alta-store/routes"
)

func main() {
	database.Connection()
	e := routes.New()
	e.Logger.Fatal(e.Start(config.Env("APP_PORT")))
}
