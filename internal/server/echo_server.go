package server

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sawalreverr/recything/config"
	"github.com/sawalreverr/recything/internal/database"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
	gr   *echo.Group
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	app := echo.New()
	app.Validator = &CustomValidator{validator: validator.New()}

	group := app.Group("/api/v1")

	return &echoServer{
		app:  app,
		db:   db,
		conf: conf,
		gr:   group,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(middleware.CORS())

	// Public Handler
	s.publicHttpHandler()

	serverPORT := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverPORT))
}
