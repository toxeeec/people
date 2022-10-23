package main

import (
	"log"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/handler"
)

func main() {
	swagger, err := people.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}
	db, err := people.PostgresConnect()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.GET("openapi.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, swagger)
	})

	g := e.Group("", middleware.OapiRequestValidator(swagger))
	h := handler.New(db)
	people.RegisterHandlers(g, h)
	e.Logger.Fatal(e.Start(":8000"))
}
