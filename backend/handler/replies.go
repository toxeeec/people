package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	people "github.com/toxeeec/people/backend"
)

func (h *handler) PostPostsPostIDReplies(c echo.Context, postID people.PostIDParam) error {
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

	if exists := h.ps.Exists(uint(postID)); !exists {
		return echo.ErrNotFound
	}

	res, err := h.ps.CreateReply(uint(postID), userID, p)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) GetPostsPostIDReplies(c echo.Context, postID people.PostIDParam, params people.GetPostsPostIDRepliesParams) error {
	pagination := people.NewPagination((*uint)(params.Before), (*uint)(params.After), (*uint)(params.Limit))
	posts, err := h.ps.Replies(uint(postID), pagination)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, posts)
}
