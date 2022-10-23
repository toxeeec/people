package people

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrHandleTaken = echo.NewHTTPError(http.StatusBadRequest, "Handle is already taken")
)
