package handler

import (
	"disease-api/internal/service"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service.Service
}

type Handler interface {
	Register(e *echo.Echo)
}

func New(services *service.Service) Handler {
	return &handler{service: services}
}

func (h *handler) Register(e *echo.Echo) {
	e.Use()
	c := e.Group("/api/group")
	{
		c.POST("/do", h.HandleSomething)
	}
}
