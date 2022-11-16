package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	people "github.com/toxeeec/people/backend"
)

func (h *handler) GetUsersHandle(c echo.Context, handle people.HandleParam) error {
	userID, _ := people.FromContext(c.Request().Context(), people.UserIDKey)

	u, err := h.us.Get(handle, &userID)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, u)
}
