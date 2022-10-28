package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	people "github.com/toxeeec/people/backend"
)

func (h *handler) PostPosts(c echo.Context) error {
	userID, ok := people.FromContext(c.Request().Context(), people.UserIDKey)
	if !ok {
		return echo.ErrInternalServerError
	}

	var p people.PostBody
	if err := c.Bind(&p); err != nil {
		return echo.ErrBadRequest
	}

	p.TrimContent()
	if err := p.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.ps.Create(userID, p)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) GetPostsPostID(c echo.Context, postID people.PostIDParam) error {
	p, err := h.ps.Get(uint(postID))
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, p)
}

func (h *handler) DeletePostsPostID(c echo.Context, postID people.PostIDParam) error {
	userID, ok := people.FromContext(c.Request().Context(), people.UserIDKey)
	if !ok {
		return echo.ErrInternalServerError
	}

	err := h.ps.Delete(uint(postID), userID)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *handler) GetUsersHandlePosts(c echo.Context, handle string, params people.GetUsersHandlePostsParams) error {
	pagination := people.NewSeekPagination((*uint)(params.Before), (*uint)(params.After), (*uint)(params.Limit))
	posts, err := h.ps.FromUser(handle, pagination)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, posts)
}
