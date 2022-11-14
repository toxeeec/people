package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/user"
)

func (h *handler) PutMeFollowingHandle(c echo.Context, handle people.HandleParam) error {
	userID, ok := people.FromContext(c.Request().Context(), people.UserIDKey)
	if !ok {
		return echo.ErrInternalServerError
	}

	follows, err := h.us.Follow(userID, handle)
	if errors.Is(err, user.ErrAlreadyFollowed) {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	if errors.Is(err, user.ErrSameUser) || errors.Is(err, user.ErrInvalidHandle) {
		return echo.ErrNotFound
	}

	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, follows)
}

func (h *handler) DeleteMeFollowingHandle(c echo.Context, handle people.HandleParam) error {
	userID, ok := people.FromContext(c.Request().Context(), people.UserIDKey)
	if !ok {
		return echo.ErrInternalServerError
	}

	follows, err := h.us.Unfollow(userID, handle)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, follows)
}

func (h *handler) GetMeFollowingHandle(c echo.Context, handle people.HandleParam) error {
	userID, ok := people.FromContext(c.Request().Context(), people.UserIDKey)
	if !ok {
		return echo.ErrInternalServerError
	}

	following, err := h.us.IsFollowing(userID, handle)
	if err != nil {
		return echo.ErrInternalServerError
	}

	if !following {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *handler) GetMeFollowing(c echo.Context, params people.GetMeFollowingParams) error {
	userID, ok := people.FromContext(c.Request().Context(), people.UserIDKey)
	if !ok {
		return echo.ErrInternalServerError
	}

	pagination := people.NewPagination(params.Before, params.After, params.Limit)
	following, err := h.us.Following(userID, pagination)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, following)
}

func (h *handler) GetUsersHandleFollowing(c echo.Context, handle string, params people.GetUsersHandleFollowingParams) error {
	u, err := h.us.GetAuth(handle)
	if err != nil {
		return echo.ErrNotFound
	}

	pagination := people.NewPagination(params.Before, params.After, params.Limit)
	following, err := h.us.Following(*u.ID, pagination)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, following)
}

func (h *handler) GetUsersHandleFollowers(c echo.Context, handle string, params people.GetUsersHandleFollowersParams) error {
	u, err := h.us.GetAuth(handle)
	if err != nil {
		return echo.ErrNotFound
	}

	pagination := people.NewPagination(params.Before, params.After, params.Limit)
	following, err := h.us.Following(*u.ID, pagination)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, following)
}
