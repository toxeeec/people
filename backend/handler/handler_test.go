package handler_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
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
	db          *sqlx.DB
	e           *echo.Echo
	us          people.UserService
	as          people.AuthService
	ps          people.PostService
	user1       people.AuthUser
	user2       people.AuthUser
	user3       people.AuthUser
	unknownUser people.AuthUser
	id1         uint
	id2         uint
	id3         uint
	at1         string
	post1       people.Post
	post2       people.Post
	postBody1   people.PostBody
	postBody2   people.PostBody
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
	gofakeit.Struct(&suite.user1)
	gofakeit.Struct(&suite.user2)
	gofakeit.Struct(&suite.user3)
	gofakeit.Struct(&suite.unknownUser)
	gofakeit.Struct(&suite.postBody1)
	gofakeit.Struct(&suite.postBody2)
	suite.id1, _ = suite.us.Create(suite.user1)
	suite.id2, _ = suite.us.Create(suite.user2)
	suite.id3, _ = suite.us.Create(suite.user3)
	suite.at1, _ = token.NewAccessToken(suite.id1)
	suite.post1, _ = suite.ps.Create(suite.id1, suite.postBody1)
	suite.post2, _ = suite.ps.Create(suite.id1, suite.postBody2)
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}
