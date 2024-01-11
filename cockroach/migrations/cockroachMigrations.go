package main

import (
	"github.com/DagmarC/clean-arch-example/cockroach/entities"
	"github.com/DagmarC/clean-arch-example/config"
	"github.com/DagmarC/clean-arch-example/database"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	migrate(db)
}

func migrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Cockroach{})
	db.GetDb().Statement.CreateInBatches([]entities.Cockroach{
		{Amount: 1},
		{Amount: 2},
		{Amount: 2},
		{Amount: 5},
		{Amount: 3},
	}, 10)
}
