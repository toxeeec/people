package http

import (
	"log"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository/postgres"
	"github.com/toxeeec/people/backend/service/auth"
	"github.com/toxeeec/people/backend/service/image"
	"github.com/toxeeec/people/backend/service/post"
	"github.com/toxeeec/people/backend/service/user"
)

type handler struct {
	as auth.Service
	us user.Service
	ps post.Service
	is image.Service
}

func NewServer(db *sqlx.DB, v *validator.Validate) *echo.Echo {
	pr := postgres.NewPostRepository(db)
	ur := postgres.NewUserRepository(db)
	tr := postgres.NewTokenRepository(db)
	fr := postgres.NewFollowRepository(db)
	lr := postgres.NewLikeRepository(db)
	ir := postgres.NewImageRepository(db)

	var h handler
	h.us = user.NewService(v, ur, fr, lr)
	h.is = image.NewService(ir)
	h.as = auth.NewService(ur, tr, h.us)
	h.ps = post.NewService(v, pr, ur, fr, lr, h.us, h.is)

	e := echo.New()
	e.Use(echomiddleware.CORS())

	swagger, err := people.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	e.GET("openapi.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, swagger)
	})
	e.Static("/images", "images")

	e.Use(middleware.OapiRequestValidatorWithOptions(swagger,
		&middleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: h.newAuthenticator(),
			},
			Skipper: func(c echo.Context) bool {
				return strings.Contains(c.Path(), "/images*") || strings.Contains(c.Path(), "openapi.json")
			},
		}))

	people.RegisterHandlers(e, people.NewStrictHandler(&h, nil))

	return e
}
