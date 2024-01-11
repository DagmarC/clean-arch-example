package usecases

import "github.com/DagmarC/clean-arch-example/cockroach/models"

type Cockroach interface {
	Process(in *models.AddCockroachData) error
}