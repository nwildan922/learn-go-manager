package routes

import (
	"github.com/labstack/echo/v4"
	c "github.com/nwildan922/learn-go-manager/controller"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", c.HealthCheck)
}
