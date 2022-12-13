package http

import (
	"log"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository/postgres"
	"github.com/toxeeec/people/backend/service/auth"
	"github.com/toxeeec/people/backend/service/post"
	"github.com/toxeeec/people/backend/service/user"
)

type handler struct {
	as auth.Service
	us user.Service
	ps post.Service
}

func NewServer(db *sqlx.DB, v *validator.Validate) *echo.Echo {
	pr := postgres.NewPostRepository(db)
	ur := postgres.NewUserRepository(db)
	tr := postgres.NewTokenRepository(db)
	fr := postgres.NewFollowRepository(db)
	lr := postgres.NewLikeRepository(db)

	var h handler
	h.as = auth.NewService(v, ur, tr)
	h.us = user.NewService(ur, fr, lr)
	h.ps = post.NewService(v, pr, ur, fr, lr, h.us)
	e := echo.New()

	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
	}))

	swagger, err := people.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	e.GET("openapi.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, swagger)
	})
	e.Use(middleware.OapiRequestValidatorWithOptions(swagger,
		&middleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: h.newAuthenticator(),
			},
		}))

	people.RegisterHandlers(e, people.NewStrictHandler(&h, nil))

	return e
}
