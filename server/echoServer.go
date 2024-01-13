package server

import (
	"fmt"

	"github.com/DagmarC/clean-arch-example/cockroach/handlers"
	"github.com/DagmarC/clean-arch-example/cockroach/repositories"
	"github.com/DagmarC/clean-arch-example/cockroach/usecases"
	"github.com/DagmarC/clean-arch-example/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func (e *echoServer) Start() {
	// Init routers here
	e.app.Use(middleware.Logger())

	e.initCockroachhttpHandler()

	serverURL := fmt.Sprintf(":%d", e.cfg.App.Port)
	e.app.Logger.Fatal(e.app.Start(serverURL))
}

func (e *echoServer) initCockroachhttpHandler() {
	// Init all layers from the bottom to top
	postgresRepository := repositories.NewCockroachPostgresRepository(e.db)
	messaging := repositories.NewCockroachFCMMessaging()

	cockroachUsecase := usecases.NewCockroachUsecaseImpl(
		postgresRepository,
		messaging,
	)

	cockroachHttpHandler := handlers.NewCockroachHttpHandler(cockroachUsecase)

	//Routers
	// routers := e.app.Group("/v4/cockroach")
	e.app.GET("/hello", handlers.HelloWorld)
	e.app.POST("", cockroachHttpHandler.DetectCockroach)
}
// HINT:
// curl -X POST -H "Content-Type: application/json" -d '{"amount": 21}' http://localhost:8080\

// If you prefer to keep using application/x-www-form-urlencoded, [curl -X POST -d "amount=21" http://localhost:8080]
// Update your Go code to use c.FormValue or similar methods to retrieve form values instead of binding JSON data.