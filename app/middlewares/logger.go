package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger Middleware
func Logger() echo.MiddlewareFunc {
	out, err := os.Create("public/logs.txt")
	if err != nil {
		out = os.Stdout
	}

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "Ip=${remote_ip}, Method=${method}, Url=\"${uri}\", Status=${status}, Latency:${latency_human} \n",
		Output: out,
	})
}
