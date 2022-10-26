package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	people "github.com/toxeeec/people/backend"
)

func (h *handler) GetMeFollowersHandle(c echo.Context, handle people.HandleParam) error {
	userID, ok := people.FromContext(c.Request().Context(), people.UserIDKey)
	if !ok {
		return echo.ErrInternalServerError
	}

	following, err := h.us.IsFollowed(userID, handle)
	if err != nil {
		return echo.ErrInternalServerError
	}

	if !following {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *handler) GetMeFollowers(c echo.Context, params people.GetMeFollowersParams) error {
	userID, ok := people.FromContext(c.Request().Context(), people.UserIDKey)
	if !ok {
		return echo.ErrInternalServerError
	}

	pagination := people.NewPagination((*uint)(params.Page), (*uint)(params.Limit))
	following, err := h.us.Followers(userID, pagination)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, following)
}
