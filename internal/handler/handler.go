package handler

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "privacy-check/docs"
	"privacy-check/internal/service"
)

type handler struct {
	service service.Service
}

func NewHandler(router *echo.Group, service service.Service) {
	handler := &handler{service: service}

	router.GET("/swagger/*", echoSwagger.WrapHandler)
	{
		router.POST("/auth/register", handler.Register)
		router.POST("/auth/login", handler.Login)
		router.GET("/search-my-leak-data", handler.Search, handler.middleware)
	}
}
