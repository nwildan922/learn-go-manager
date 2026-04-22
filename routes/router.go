package routes

import (
	"github.com/labstack/echo/v4"
	c "github.com/nwildan922/learn-go-manager/controller"
)

type Router struct {
	Counter *c.CounterController
}

func RegisterRoutes(e *echo.Echo, rc *Router) {
	e.GET("/health", c.HealthCheck)
	e.GET("/counter", rc.Counter.GenerateCounter)
}
