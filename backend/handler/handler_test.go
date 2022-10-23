package handler_test

import (
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/handler"
	"github.com/toxeeec/people/backend/service/auth"
	"github.com/toxeeec/people/backend/service/user"
)

type HandlerSuite struct {
	suite.Suite
	db *sqlx.DB
	e  *echo.Echo
	us people.UserService
	as people.AuthService
}

func (suite *HandlerSuite) SetupSuite() {
	db, err := people.PostgresConnect()
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.db = db
	suite.us = user.NewService(db)
	suite.as = auth.NewService(db, suite.us)
	suite.e = echo.New()
	swagger, err := people.GetSwagger()
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.e.Use(middleware.OapiRequestValidator(swagger))
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
