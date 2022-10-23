package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	people "github.com/toxeeec/people/backend"
)

func (h *handler) PostRegister(c echo.Context) error {
	var u people.AuthUser
	if err := c.Bind(&u); err != nil {
		return echo.ErrInternalServerError
	}

	if err := u.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if exists := h.us.Exists(u.Handle); exists {
		return people.ErrHandleTaken
	}

	id, err := h.us.Create(u)
	if err != nil {
		return echo.ErrInternalServerError
	}

	tokens, err := h.as.NewTokens(id)
	if err != nil {
		go h.us.Delete(u.Handle)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, tokens)
}

func (h *handler) PostLogin(c echo.Context) error {
	var u people.AuthUser
	if err := c.Bind(&u); err != nil {
		return echo.ErrInternalServerError
	}

	id, err := h.as.VerifyCredentials(u)
	if err != nil {
		return people.ErrInvalidCredentials
	}

	tokens, err := h.as.NewTokens(id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, tokens)
}
