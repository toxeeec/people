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

	pagination := people.NewPagination(params.Before, params.After, params.Limit)
	following, err := h.us.Followers(userID, &userID, pagination)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, following)
}

func (h *handler) GetUsersHandleFollowers(c echo.Context, handle string, params people.GetUsersHandleFollowersParams) error {
	userID, _ := people.FromContext(c.Request().Context(), people.UserIDKey)
	u, err := h.us.GetAuthUser(handle)
	if err != nil {
		return echo.ErrNotFound
	}

	println(userID)
	pagination := people.NewPagination(params.Before, params.After, params.Limit)
	following, err := h.us.Followers(*u.ID, &userID, pagination)
	if err != nil {
		println(err.Error())
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, following)
}
