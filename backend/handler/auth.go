package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/token"
)

func (h *handler) PostRegister(c echo.Context) error {
	var u people.AuthUser
	if err := c.Bind(&u); err != nil {
		return echo.ErrBadRequest
	}

	if err := u.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if exists := h.us.Exists(u.Handle); exists {
		return people.ErrTakenHandle
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
		return echo.ErrBadRequest
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

func (h *handler) PostRefresh(c echo.Context) error {
	var t people.Tokens
	if err := c.Bind(&t); err != nil {
		return echo.ErrBadRequest
	}

	rt, err := token.ParseRefreshToken(t.RefreshToken)
	if err != nil {
		return echo.ErrForbidden
	}

	valid := h.as.CheckRefreshToken(rt)
	if !valid {
		return echo.ErrForbidden
	}

	at, err := token.NewAccessToken(rt.UserID)
	if err != nil {
		return echo.ErrInternalServerError
	}

	newRT, err := h.as.UpdateRefreshToken(rt.UserID, rt.ID)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, people.Tokens{AccessToken: at, RefreshToken: newRT.Value})
}
