package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"privacy-check/internal/service"
)

type handler struct {
	service service.Service
}

func NewHandler(router *echo.Group, service service.Service) {

	group := router.Group("/user")
	{
		handler := &handler{service: service}

		group.POST("", handler.Create)
		//group.GET("/page", handler.Page)
		//group.GET("/:id", handler.GetById)
		//group.PATCH("/:id", handler.Update)
		//group.DELETE("/:id", handler.Delete)

	}
}

func (h *handler) Create(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "not implemented")
}
