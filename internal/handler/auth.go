package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"privacy-check/internal/models"
)

// Register godoc
// @Summary      User Registration
// @Tags         Auth
// @Description  Registers a new user with their details.
// @ID           register
// @Accept       json
// @Produce      json
// @Param        input body models.RegisterDTO true "Registration details"
// @Success      201 {object} map[string]string "User registered successfully"
// @Failure      400 {string} string "Bad request"
// @Failure      500 {string} string "Internal server error"
// @Router       /auth/register [post]
func (h *handler) Register(c echo.Context) error {
	var input models.RegisterDTO

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err := h.service.Create(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User registered successfully",
	})
}

// Login godoc
// @Summary      User Login
// @Tags         Auth
// @Description  Logs in a user with email and password, returning a JWT token.
// @ID           login
// @Accept       json
// @Produce      json
// @Param        input body models.LoginDTO true "User credentials"
// @Success      200 {object} models.TokenResponse "JWT token for authentication"
// @Failure      400 {string} string "Bad request"
// @Failure      401 {string} string "Unauthorized"
// @Failure      500 {string} string "Internal server error"
// @Router       /auth/login [post]
func (h *handler) Login(c echo.Context) error {
	var input models.LoginDTO

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := h.service.GenerateToken(input.Email, input.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.TokenResponse{
		Token: token,
	})
}
