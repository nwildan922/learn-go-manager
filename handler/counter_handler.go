package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/nwildan922/learn-go-manager/service"
)

type CounterHandler struct {
	service *service.CounterService
}

func NewCounterHandler(s *service.CounterService) *CounterHandler {
	return &CounterHandler{
		service: s,
	}
}

func (c *CounterHandler) GenerateCounter(ctx echo.Context) error {
	res := c.service.Count()
	return ctx.JSON(res.Code, res)
}
