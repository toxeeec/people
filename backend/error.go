package people

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrTakenHandle        = echo.NewHTTPError(http.StatusBadRequest, "Handle is already taken")
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Invalid handle or password")
)
