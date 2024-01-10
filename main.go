package main

import (
	"github.com/DagmarC/clean-arch-example/config"
	"github.com/DagmarC/clean-arch-example/database"
	"github.com/DagmarC/clean-arch-example/server"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	server.NewEchoServer(&cfg, db.GetDb()).Start()
}

