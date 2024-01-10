package server

import (
	"fmt"

	"github.com/DagmarC/clean-arch-example/config"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *gorm.DB) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {
	// Init routers here

	s.app.Use(middleware.Logger())
	serverURL := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverURL))
}
