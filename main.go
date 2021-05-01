package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"goBoilterplate/app/models"
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

	app.Logger.Info("Db connected")


	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Project{})

	//db.Model(&models.User{}).CreateTable()
	//db.Model(&models.Project{}).CreateTable()
	//db.Create(&models.User{Name: "admin", Email: "adem@flutterturkiye.org", Password: "1234", TeamName: "admin"})
	//db.Create(&models.Project{TeamName: "test", Name: "admin", Description: "lorem ipsum"})

	defer db.Close()

	//console.Schedule()

	router.Init(app)

	app.Server.Addr = ":8080"
	graceful.ListenAndServe(app.Server, 5*time.Second)
}
