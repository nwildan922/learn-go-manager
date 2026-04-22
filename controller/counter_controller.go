package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/nwildan922/learn-go-manager/service"
)

type CounterController struct {
	service *service.CounterService
}

func NewCounterController(s *service.CounterService) *CounterController {
	return &CounterController{
		service: s,
	}
}

func (c *CounterController) GenerateCounter(ctx echo.Context) error {
	res := c.service.Count()
	return ctx.JSON(res.Code, res)
}
