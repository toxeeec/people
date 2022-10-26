package handler_test

import (
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/handler"
	"github.com/toxeeec/people/backend/service/auth"
	"github.com/toxeeec/people/backend/service/post"
	"github.com/toxeeec/people/backend/service/user"
	"github.com/toxeeec/people/backend/token"
)

type HandlerSuite struct {
	suite.Suite
	db *sqlx.DB
	e  *echo.Echo
	us people.UserService
	as people.AuthService
	ps people.PostService
}

func (suite *HandlerSuite) SetupSuite() {
	db, err := people.PostgresConnect()
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.db = db
	suite.us = user.NewService(db)
	suite.as = auth.NewService(db, suite.us)
	suite.ps = post.NewService(db)
	suite.e = echo.New()
	swagger, err := people.GetSwagger()
	if err != nil {
		suite.T().Fatal(err)
	}
	validator := middleware.OapiRequestValidatorWithOptions(swagger,
		&middleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: token.NewAuthenticator(),
			},
		})
	suite.e.Use(validator)
	h := handler.New(db)
	people.RegisterHandlers(suite.e, h)
}

func (suite *HandlerSuite) TearDownSuite() {
	suite.db.Close()
}

func (suite *HandlerSuite) SetupTest() {
	suite.db.MustExec("TRUNCATE user_profile CASCADE")
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}
