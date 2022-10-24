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

	err := h.us.Follow(userID, handle)
	if errors.Is(err, user.ErrAlreadyFollowed) {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	if errors.Is(err, user.ErrSameUser) || errors.Is(err, user.ErrInvalidHandle) {
		return echo.ErrNotFound
	}

	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusNoContent)
}
