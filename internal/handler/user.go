package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Search godoc
// @Summary      Search user leak data
// @Description  Search user leak data by email
// @Tags         User
// @ID           search-user-leak-data
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200 {object} models.LeakData "Successful operation"
// @Failure      401 {string} string "Unauthorized"
// @Failure      500 {object} string "Internal server error"
// @Router       /search-my-leak-data [get]
func (h *handler) Search(c echo.Context) error {
	userId, err := h.getUserId(c)
	if err != nil {
		return err
	}

	leakData, err := h.service.SearchLeakDataById(userId)
	{
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, leakData)
}
