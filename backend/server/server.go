package server

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
	peoplehttp "github.com/toxeeec/people/backend/http"
	"github.com/toxeeec/people/backend/repository/postgres"
	"github.com/toxeeec/people/backend/service/auth"
	"github.com/toxeeec/people/backend/service/image"
	"github.com/toxeeec/people/backend/service/message"
	"github.com/toxeeec/people/backend/service/notification"
	"github.com/toxeeec/people/backend/service/post"
	"github.com/toxeeec/people/backend/service/user"
	"github.com/toxeeec/people/backend/ws"
)

func New(db *sqlx.DB, v *validator.Validate) *echo.Echo {
	pr := postgres.NewPostRepository(db)
	ur := postgres.NewUserRepository(db)
	tr := postgres.NewTokenRepository(db)
	fr := postgres.NewFollowRepository(db)
	lr := postgres.NewLikeRepository(db)
	ir := postgres.NewImageRepository(db)
	mr := postgres.NewMessageRepository(db)

	is := image.NewService(ir)
	us := user.NewService(v, ur, fr, lr, is)
	as := auth.NewService(ur, tr, us)
	ps := post.NewService(v, pr, ur, fr, lr, us, is)
	notif := make(chan people.Notification, 256)
	ns := notification.NewService(notif, ur)
	ms := message.NewService(mr, ur, ns, us)

	e := echo.New()
	e.Use(echomiddleware.CORS())
	e.Static("/images", "images")

	hub := ws.NewHub(notif, ms)
	go hub.Run()
	e.GET("/ws", ws.Serve(hub), peoplehttp.AuthMiddleware)

	swagger, err := people.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}
	e.GET("openapi.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, swagger)
	})
	h := peoplehttp.NewHandler(v, as, us, ps, is, ms)
	e.Use(middleware.OapiRequestValidatorWithOptions(swagger,
		&middleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: h.NewAuthenticator(),
			},
			Skipper: func(c echo.Context) bool {
				return !strings.HasPrefix(c.Path(), "/api/")
			},
		}))
	people.RegisterHandlersWithBaseURL(e, people.NewStrictHandler(h, nil), "/api")
	return e
}
