package routes

import (
	"github.com/labstack/echo/v4"
	handler "github.com/nwildan922/learn-go-manager/handler"
)

type Router struct {
	Counter *handler.CounterHandler
}

func RegisterRoutes(e *echo.Echo, rc *Router) {
	e.GET("/health", handler.HealthCheck)
	e.GET("/counter", rc.Counter.GenerateCounter)
}
