package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toxeeec/people/backend/spec"
)

type handler struct{}

func (h *handler) Get(c echo.Context) error {
	return c.JSONBlob(200, []byte(`{"hello": "world"}`))
}

func main() {
	e := echo.New()
	e.GET("openapi.json", func(c echo.Context) error {
		spec, err := spec.GetSwagger()
		if err != nil {
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, spec)
	})
	spec.RegisterHandlers(e, &handler{})
	e.Logger.Fatal(e.Start(":8000"))
}
