package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	people "github.com/toxeeec/people/backend"
)

func (h *handler) GetMeFeed(c echo.Context, params people.GetMeFeedParams) error {
	userID, ok := people.FromContext(c.Request().Context(), people.UserIDKey)
	if !ok {
		return echo.ErrInternalServerError
	}

	pagination := people.NewPagination((*uint)(params.Before), (*uint)(params.After), (*uint)(params.Limit))
	posts, err := h.ps.Feed(userID, pagination)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, posts)
}
