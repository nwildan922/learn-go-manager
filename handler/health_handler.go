package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nwildan922/learn-go-manager/service"
)

func HealthCheck(c echo.Context) error {
	res := service.HealthCheck("running")
	return c.JSON(http.StatusOK, res)
}
