package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtxId           = "userId"
)

func (h *handler) middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get(authorizationHeader)
		if header == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Empty auth header")
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid auth header")
		}

		userId, err := h.service.ParseToken(headerParts[1])
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Invalid token: %v", err))
		}

		c.Set(userCtxId, userId)

		return next(c)
	}
}

func (h *handler) getUserId(c echo.Context) (int, error) {
	id := c.Get(userCtxId)
	if id == nil {
		return 0, echo.NewHTTPError(http.StatusInternalServerError, "User id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, echo.NewHTTPError(http.StatusInternalServerError, "User id is invalid type")
	}

	return idInt, nil
}
