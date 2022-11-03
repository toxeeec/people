package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/post"
)

func (h *handler) PutPostsPostIDLikes(c echo.Context, postID uint) error {
	userID, ok := people.FromContext(c.Request().Context(), people.UserIDKey)
	if !ok {
		return echo.ErrInternalServerError
	}

	l, err := h.ps.Like(postID, userID)
	if errors.Is(err, post.ErrAlreadyLiked) {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	if errors.Is(err, post.ErrInvalidPostID) {
		return echo.ErrNotFound
	}

	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, l)
}
