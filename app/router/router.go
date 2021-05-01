package router

import (
	"goBoilterplate/app/controllers"
	"goBoilterplate/app/middlewares"
	_ "goBoilterplate/docs" // For Swagger

	"log"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Init Router
func Init(app *echo.Echo) {
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger())
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())

	app.GET("/", controllers.Index)
	app.GET("/ready", controllers.Ready)
	app.GET("/docs/*", echoSwagger.WrapHandler)

	api := app.Group("/api", middlewares.Jwt())
	{
		api.POST("/login", controllers.Login)
		api.GET("/logout", controllers.Logout)

		projects := api.Group("/projects")
		{
			projects.GET("", controllers.ProjectList)
			projects.GET("/", controllers.ProjectList)
			projects.GET("/:id", controllers.ProjectDetail)
			projects.POST("/:id", controllers.ProjectUpvote)
			projects.DELETE("/:id", controllers.ProjectDownvote)
		}
	}

	log.Printf("Server started...")
}
