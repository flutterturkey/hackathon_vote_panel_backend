package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"goBoilterplate/app/console"
	"goBoilterplate/app/router"
	"goBoilterplate/config"
	"gopkg.in/tylerb/graceful.v1"
	"time"
)

// @title Golang Echo API
// @version 1.0
// @description API documentation by Swagger

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

var (
	Version = "dev"
)

func main() {
	app := echo.New()
	app.Logger.SetLevel(log.INFO)
	app.Logger.Info("Http server started with version ", Version)

	db, err := config.Database()
	if err != nil {
		app.Logger.Fatal(err)
	}
	defer db.Close()

	config.Redis()
	console.Schedule()
	router.Init(app)

	app.Server.Addr = ":8080"
	graceful.ListenAndServe(app.Server, 5*time.Second)
}
